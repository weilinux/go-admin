package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"net/http"
)

// func upload(r *http.Request) (string, error) {
// 	f, h, err := r.FormFile(param)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()
//
// 	p := filepath.Join(os.TempDir(), h.Filename)
// 	fw, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer fw.Close()
//
// 	if _, err = io.Copy(fw, f); err != nil {
// 		return "", err
// 	}
// 	return p, nil
// }

// UpLoad 上传图片接口
func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"url":    url,
	})
}
