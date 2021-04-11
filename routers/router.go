package routers

import (
	"github.com/ElijahCYN/Go-gin-api/pkg/setting"
	v1 "github.com/ElijahCYN/Go-gin-api/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 獲取標籤列表
		apiv1.GET("/tags", v1.GetTags)
		// 新增標籤
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定標籤
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 刪除指定標籤
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 獲取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 獲取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新增文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 刪除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return  r
}