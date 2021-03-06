package router

import (
	"quan/api/v1"
	"quan/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	JwtRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		JwtRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
