package router

import (
"github.com/gin-gonic/gin"
"go-admin/middleware"
"go-admin/apis/demoelementui"
jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerDemoElementuiRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

r := v1.Group("/demoelementui").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
r.GET("/:id", demoelementui.GetDemoElementui)
r.POST("", demoelementui.InsertDemoElementui)
r.PUT("", demoelementui.UpdateDemoElementui)
r.DELETE("/:id", demoelementui.DeleteDemoElementui)
}

l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
l.GET("/demoelementuiList",demoelementui.GetDemoElementuiList)
}

}
