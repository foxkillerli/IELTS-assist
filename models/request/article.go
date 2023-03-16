package request

type PostArticleEdit struct {
	Article string `json:"article" binding:"required"`
	Band    int    `json:"band" binding:"required"`
}
