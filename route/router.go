package route

import (
	"github.com/foxkillerli/IELTS-assist/controllers"
	jwt "github.com/foxkillerli/IELTS-assist/middleware/myjwt"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/api/v1/user/register", controllers.UserRegister)
	r.POST("/api/v1/user/login", controllers.UserLogin)
	v1NeedAuth := r.Group("/api/v1")
	v1NeedAuth.Use(jwt.NeedJwtAuth())
	{
		v1NeedAuth.POST("article/edit", controllers.ArticleEdit)
		v1NeedAuth.POST("oral/chat", controllers.OralChat)
	}

	return r
}

func init() {
	log.Printf("route init")
}
