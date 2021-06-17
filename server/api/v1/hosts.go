package v1

import (
	"quan/global"
    "quan/model"
    "quan/model/request"
    "quan/model/response"
    "quan/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Hosts
// @Summary 创建Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "创建Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosts/createHosts [post]
func CreateHosts(c *gin.Context) {
	var hosts model.Hosts
	_ = c.ShouldBindJSON(&hosts)
	if err := service.CreateHosts(hosts); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Hosts
// @Summary 删除Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "删除Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosts/deleteHosts [delete]
func DeleteHosts(c *gin.Context) {
	var hosts model.Hosts
	_ = c.ShouldBindJSON(&hosts)
	if err := service.DeleteHosts(hosts); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hosts
// @Summary 批量删除Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosts/deleteHostsByIds [delete]
func DeleteHostsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteHostsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Hosts
// @Summary 更新Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "更新Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosts/updateHosts [put]
func UpdateHosts(c *gin.Context) {
	var hosts model.Hosts
	_ = c.ShouldBindJSON(&hosts)
	if err := service.UpdateHosts(hosts); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Hosts
// @Summary 用id查询Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "用id查询Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosts/findHosts [get]
func FindHosts(c *gin.Context) {
	var hosts model.Hosts
	_ = c.ShouldBindQuery(&hosts)
	if err, rehosts := service.GetHosts(hosts.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosts": rehosts}, c)
	}
}

// @Tags Hosts
// @Summary 分页获取Hosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HostsSearch true "分页获取Hosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosts/getHostsList [get]
func GetHostsList(c *gin.Context) {
	var pageInfo request.HostsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHostsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
