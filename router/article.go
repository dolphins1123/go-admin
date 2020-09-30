package router

import (
"github.com/gin-gonic/gin"
"go-admin/apis/article"
)

// 无需认证的路由代码
func registerArticleRouter(v1 *gin.RouterGroup) {

v1.GET("/articleList",article.GetArticleList)

r := v1.Group("/article")
{
r.GET("/:id", article.GetArticle)
r.POST("", article.InsertArticle)
r.PUT("", article.UpdateArticle)
r.DELETE("/:id", article.DeleteArticle)
}
}
