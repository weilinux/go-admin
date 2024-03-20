package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"net/http"
)

// 为response返回结果的
type SuccessResponse struct {
	Code int         `json:"code" example:"200"`    // 响应编码：0成功 401请登录 403无权限 500错误
	Msg  string      `json:"msg" example:"success"` // 消息提示
	Data interface{} `json:"data"`                  // 数据对象
}

type CaptchaResponse struct {
	Code  int         `json:"code"`  // 响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   // 消息
	Data  interface{} `json:"data"`  // 数据内容
	IdKey string      `json:"idkey"` // 验证码ID
}

// 结果返回统一处理
/**
 * response 统一格式
 * {
 *    code: 200,
 *    msg: '消息[String]',
 *    data: '返回数据[Any]'
 * }
 */

type GoAdminResponse struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *GoAdminResponse {
	return &GoAdminResponse{
		Ctx: ctx,
	}
}

func (r *GoAdminResponse) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *GoAdminResponse) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, &ResponseCommonStruct{
		Data:  list,
		Count: totalRows,
		Pager: Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *GoAdminResponse) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(http.StatusOK, response)
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

type ResponseCommonStruct struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Pager Pager       `json:"pager"`
	Count int         `json:"count"`
}
