package router

import (
	"quan/api/v1"
	"quan/middleware"
	"github.com/gin-gonic/gin"
)

func InitHostsRouter(Router *gin.RouterGroup) {
	HostsRouter := Router.Group("hosts").Use(middleware.OperationRecord())
	{
		HostsRouter.POST("createHosts", v1.CreateHosts)   // 新建Hosts
		HostsRouter.DELETE("deleteHosts", v1.DeleteHosts) // 删除Hosts
		HostsRouter.DELETE("deleteHostsByIds", v1.DeleteHostsByIds) // 批量删除Hosts
		HostsRouter.PUT("updateHosts", v1.UpdateHosts)    // 更新Hosts
		HostsRouter.GET("findHosts", v1.FindHosts)        // 根据ID获取Hosts
		HostsRouter.GET("getHostsList", v1.GetHostsList)  // 获取Hosts列表
	}
}
