package request

// Apart from being an excellent ORM for Go developers, it's built to be developer-friendly
// and easy to comprehend. Gorm is built on top of the database/sql packages.
// overview and features of the ORM are:
// Developer Friendly
// Every feature comes with a tests
// Hooks / Callbacks (Before/After Create/Save/Update/Delete/Find)
// Eager loading with Preload, Joins
// Context, Prepared Statement Mode, DryRun Mode
// SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
// Transactions, Nested Transactions, Save Point, Rollback to Saved Point
// Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
// SQL Builder
// Logger

// type UserReq struct {
// 	*model.Model
// 	UserName string
// }

// // Login User login structure
// type Login struct {
// 	Username  string `json:"username" validate:"required"`  // 用户名
// 	Password  string `json:"password" validate:"required"`  // 密码
// 	Captcha   string `json:"captcha" validate:"required"`   // 验证码
// 	CaptchaId string `json:"captchaId" validate:"required"` // 验证码ID
// }
//
// type AddUser struct {
// 	Username string `json:"username" validate:"required"` // 用户名
// 	Password string `json:"password" validate:"required"` // 密码
// 	// Phone       string `json:"phone" validate:"required,regexp=^[1][0-9]{10}$"` // 手机号
// 	Phone       string `json:"phone"`                      // 手机号
// 	Email       string `json:"email" validate:"email"`     // 邮箱
// 	Active      bool   `json:"active"`                     // 是否活跃
// 	RoleModelID uint   `json:"roleId" validate:"required"` // 角色ID
// }
//
// type EditUser struct {
// 	Id          uint   `json:"id" validate:"required"`
// 	Username    string `json:"username" validate:"required"` // 用户名
// 	Phone       string `json:"phone"`                        // 手机号
// 	Email       string `json:"email"`                        // 邮箱
// 	Active      bool   `json:"active"`                       // 是否活跃
// 	RoleModelID uint   `json:"roleId" validate:"required"`   // 角色ID
// }
//
// type ModifyPass struct {
// 	Id          uint   `json:"id" validate:"required"`
// 	OldPassword string `json:"oldPassword" validate:"required"` // 旧密码
// 	NewPassword string `json:"newPassword" validate:"required"` // 新密码
// }
//
// type SwitchActive struct {
// 	Id     uint `json:"id" validate:"required"`
// 	Active bool `json:"active"`
// }
