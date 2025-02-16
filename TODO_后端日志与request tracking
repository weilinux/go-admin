package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
    "net/http"
    "time"
    "os"
)

// Todo model remains the same
type Todo struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" binding:"required" gorm:"not null"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var db *gorm.DB

// Custom logger middleware
func Logger() gin.HandlerFunc {
    // Set up the log file
    f, _ := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        // Custom log format
        return fmt.Sprintf("%s | %3d | %13v | %15s | %s | %s | %s\n",
            param.TimeStamp.Format("2006-01-02 15:04:05"),
            param.StatusCode,
            param.Latency,
            param.ClientIP,
            param.Method,
            param.Path,
            param.ErrorMessage,
        )
    })
}

// Request tracking middleware
func RequestTracker() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Generate request ID
        requestID := fmt.Sprintf("%d", time.Now().UnixNano())
        c.Set("RequestID", requestID)

        // Log request start
        log.Printf("[%s] Request started: %s %s",
            requestID,
            c.Request.Method,
            c.Request.URL.Path,
        )

        // Time tracking
        start := time.Now()

        // Process request
        c.Next()

        // Log request completion
        duration := time.Since(start)
        log.Printf("[%s] Request completed in %v with status %d",
            requestID,
            duration,
            c.Writer.Status(),
        )
    }
}

func initDB() {
    var err error
    dsn := "host=localhost user=postgres password=postgres dbname=todos port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Enable GORM logging
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = db.AutoMigrate(&Todo{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}

// Logging wrapper for database operations
func logDBOperation(operation string, fn func() error) error {
    start := time.Now()
    err := fn()
    duration := time.Since(start)

    if err != nil {
        log.Printf("DB Operation [%s] failed in %v: %v", operation, duration, err)
        return err
    }

    log.Printf("DB Operation [%s] completed in %v", operation, duration)
    return nil
}

// CREATE - Add a new todo with logging
func createTodo(c *gin.Context) {
    requestID, _ := c.Get("RequestID")
    log.Printf("[%v] Creating new todo", requestID)

    var newTodo Todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        log.Printf("[%v] Invalid todo data: %v", requestID, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := logDBOperation("CreateTodo", func() error {
        return db.Create(&newTodo).Error
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }

    log.Printf("[%v] Successfully created todo with ID: %d", requestID, newTodo.ID)
    c.JSON(http.StatusCreated, newTodo)
}

// READ - Get all todos with logging
func getTodos(c *gin.Context) {
    requestID, _ := c.Get("RequestID")
    var todos []Todo
    page := c.DefaultQuery("page", "1")
    pageSize := c.DefaultQuery("limit", "10")

    log.Printf("[%v] Fetching todos (page: %s, limit: %s)", requestID, page, pageSize)

    var total int64
    db.Model(&Todo{}).Count(&total)

    err := logDBOperation("GetTodos", func() error {
        return db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&todos).Error
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
        return
    }

    log.Printf("[%v] Successfully fetched %d todos", requestID, len(todos))
    c.JSON(http.StatusOK, gin.H{
        "todos": todos,
        "total": total,
        "page": page,
        "limit": pageSize,
    })
}

// Other CRUD operations with logging...
// (getTodo, updateTodo, deleteTodo remain similar but with added logging)

func main() {
    // Set Gin mode
    gin.SetMode(gin.ReleaseMode)

    // Initialize database
    initDB()

    // Create a new gin router with default middleware
    router := gin.New() // Don't use gin.Default() as we'll add our own logging

    // Add recovery middleware
    router.Use(gin.Recovery())

    // Add custom logging middleware
    router.Use(Logger())

    // Add request tracking middleware
    router.Use(RequestTracker())

    // Routes
    router.GET("/todos", getTodos)
    router.POST("/todos", createTodo)
    router.GET("/todos/:id", getTodo)
    router.PUT("/todos/:id", updateTodo)
    router.DELETE("/todos/:id", deleteTodo)

    // Start server with logging
    log.Println("Starting server on :8000")
    if err := router.Run(":8000"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}