package errcode

var (
	ErrorGetUserListFail = NewError(20010001, "获取用户列表失败")
	ErrorCreateUserFail  = NewError(20010002, "创建用户失败")
	ErrorUpdateUserFail  = NewError(20010003, "更新用户失败")
	ErrorDeleteUserFail  = NewError(20010004, "删除用户失败")
	ErrorCountUserFail   = NewError(20010005, "统计用户失败")
	ErrorExistUserFail   = NewError(20010006, "已注册用户失败")

	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
