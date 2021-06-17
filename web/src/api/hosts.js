import service from '@/utils/request'

// @Tags Hosts
// @Summary 创建Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "创建Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosts/createHosts [post]
export const createHosts = (data) => {
     return service({
         url: "/hosts/createHosts",
         method: 'post',
         data
     })
 }


// @Tags Hosts
// @Summary 删除Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "删除Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosts/deleteHosts [delete]
 export const deleteHosts = (data) => {
     return service({
         url: "/hosts/deleteHosts",
         method: 'delete',
         data
     })
 }

// @Tags Hosts
// @Summary 删除Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosts/deleteHosts [delete]
 export const deleteHostsByIds = (data) => {
     return service({
         url: "/hosts/deleteHostsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Hosts
// @Summary 更新Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "更新Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosts/updateHosts [put]
 export const updateHosts = (data) => {
     return service({
         url: "/hosts/updateHosts",
         method: 'put',
         data
     })
 }


// @Tags Hosts
// @Summary 用id查询Hosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hosts true "用id查询Hosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosts/findHosts [get]
 export const findHosts = (params) => {
     return service({
         url: "/hosts/findHosts",
         method: 'get',
         params
     })
 }


// @Tags Hosts
// @Summary 分页获取Hosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Hosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosts/getHostsList [get]
 export const getHostsList = (params) => {
     return service({
         url: "/hosts/getHostsList",
         method: 'get',
         params
     })
 }