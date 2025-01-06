package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/convert"
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

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
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

func (r *Response) ToErrorResponse(err *errcode.Error) {
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

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		// return global.AppSetting.DefaultPageSize
		return 10
	}
	// if pageSize > global.AppSetting.MaxPageSize {
	// 	return global.AppSetting.MaxPageSize
	// }
	if pageSize > 100 {
		return 50
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
