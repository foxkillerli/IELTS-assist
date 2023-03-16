package controllers

import (
	"github.com/foxkillerli/IELTS-assist/middleware/openai"
	"github.com/foxkillerli/IELTS-assist/models/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostArticleEdit
// 1. Get Article from gin's context
// 2. Get aiming IELTS band from gin's context
func ArticleEdit(c *gin.Context) {
	var postArticle request.PostArticleEdit
	if err := c.BindJSON(&postArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch := make(chan string, 1)
	result := openai.ArticleEditThroughChat(postArticle.Article, postArticle.Band)
	ch <- result
	select {
	case result = <-ch:
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "success",
			"data":   result,
		})
	}
}

func ArticleEditionSuggestion(c *gin.Context) {
	var postArticle request.PostArticleEdit
	if err := c.BindJSON(&postArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch := make(chan string, 1)
	result := openai.ArticleEditSuggestion(postArticle.Article, postArticle.Band)
	ch <- result
	select {
	case result = <-ch:
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "success",
			"data":   result,
		})
	}
}
