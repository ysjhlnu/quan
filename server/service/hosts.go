package service

import (
	"quan/global"
	"quan/model"
	"quan/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateHosts
//@description: 创建Hosts记录
//@param: hosts model.Hosts
//@return: err error

func CreateHosts(hosts model.Hosts) (err error) {
	err = global.GVA_DB.Create(&hosts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHosts
//@description: 删除Hosts记录
//@param: hosts model.Hosts
//@return: err error

func DeleteHosts(hosts model.Hosts) (err error) {
	err = global.GVA_DB.Delete(&hosts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHostsByIds
//@description: 批量删除Hosts记录
//@param: ids request.IdsReq
//@return: err error

func DeleteHostsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Hosts{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateHosts
//@description: 更新Hosts记录
//@param: hosts *model.Hosts
//@return: err error

func UpdateHosts(hosts model.Hosts) (err error) {
	err = global.GVA_DB.Save(&hosts).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHosts
//@description: 根据id获取Hosts记录
//@param: id uint
//@return: err error, hosts model.Hosts

func GetHosts(id uint) (err error, hosts model.Hosts) {
	err = global.GVA_DB.Where("id = ?", id).First(&hosts).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHostsInfoList
//@description: 分页获取Hosts记录
//@param: info request.HostsSearch
//@return: err error, list interface{}, total int64

func GetHostsInfoList(info request.HostsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Hosts{})
    var hostss []model.Hosts
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Instanceid != "" {
        db = db.Where("`instanceid` LIKE ?","%"+ info.Instanceid+"%")
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Privateip != "" {
        db = db.Where("`privateip` LIKE ?","%"+ info.Privateip+"%")
    }
    if info.Region != "" {
        db = db.Where("`region` LIKE ?","%"+ info.Region+"%")
    }
    if info.Env != "" {
        db = db.Where("`env` LIKE ?","%"+ info.Env+"%")
    }
    if info.Status != "" {
        db = db.Where("`status` LIKE ?","%"+ info.Status+"%")
    }
    if info.Type != "" {
        db = db.Where("`type` = ?",info.Type)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hostss).Error
	return err, hostss, total
}