package dao

import (
	"backend/core/db"
	model "backend/models"
)

func Create(u *model.UserInfo) error {
	result := db.Db.Model(&model.UserInfo{}).Create(u)
	return result.Error
}

// 通过id查询
func GetUserInfoByID(u *model.UserInfo) error {
	result := db.Db.Model(&model.UserInfo{}).First(u)
	return result.Error
}

func IsExist(u *model.UserInfo) error {
	result := db.Db.Model(&model.UserInfo{}).Where("name = ? and password = ?", u.Name, u.Password).First(u)
	return result.Error
}

type UpdateFields struct {
	Password string
	Intro    string
	// 添加其他需要更新的字段
}

func UpdatePassword(u *model.UserInfo) error {
	updateFields := UpdateFields{
		Password: u.Password,
		Intro:    u.Intro,
		// 添加其他需要更新的字段及其新值
	}
	result := db.Db.Model(&model.UserInfo{}).Where("name = ?", u.Name).Updates(updateFields)
	return result.Error
}

// 发布活动
func StoreArticle(u *model.ReleaseInfo) error {
	result := db.Db.Create(u)
	return result.Error
}

// 查询活动
func GetActInfo(u *model.ReleaseInfo) ([]model.ReleaseInfo, error) {
	var acts []model.ReleaseInfo
	if err := db.Db.Find(&acts).Error; err != nil {
		return nil, err
	}
	return acts, nil
}

// 审核通过后修改状态
func UpdateState(u *model.ReleaseInfo) error {
	result := db.Db.Model(&model.ReleaseInfo{}).Where("act_id = ?", u.ActID).Update("state", u.State)
	return result.Error
}

// 通过活动id查询活动
func GetActInfoByID(ActID int, u *model.ReleaseInfo) (model.ReleaseInfo, error) {
	var act model.ReleaseInfo
	if err := db.Db.First(&act, ActID).Error; err != nil {
		return model.ReleaseInfo{}, err
	}
	return act, nil
}

func GetMessageByID(u *model.MessageInfo) ([]model.MessageInfo, error) {
	var msgs []model.MessageInfo
	if err := db.Db.Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}
