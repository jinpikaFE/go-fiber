package models

import (
	// "time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	Model

	// query tag是query参数别名，json xml，form适合post
	Username  *string `validate:"" gorm:"unique;comment:'用户名'" query:"username" json:"username" xml:"username" form:"username"`
	Password  string  `validate:"" gorm:"comment:'密码'" query:"password" json:"password" xml:"password" form:"password"`
	Mobile    *string `validate:"" gorm:"unique;comment:'手机号'" query:"mobile" json:"mobile" xml:"mobile" form:"mobile"`
	AvatarUrl *string `validate:"" gorm:"comment:'头像'" query:"avatarUrl" json:"avatarUrl" xml:"avatarUrl" form:"avatarUrl"`
	NickName  *string `validate:"" gorm:"comment:'昵称'" query:"nickName" json:"nickName" xml:"nickName" form:"nickName"`
	Gender    *string `validate:"" sql:"type:enum('0', '1', '2')" gorm:"comment:'性别,0未知 1男 2女';default:'0'" query:"gender" json:"gender" xml:"gender" form:"gender"`
	Unionid   *string `validate:"" gorm:"comment:'小程序unionid'" query:"unionid" json:"unionid" xml:"unionid" form:"unionid"`
	Openid    *string `validate:"" gorm:"comment:'小程序openid'" query:"openid" json:"openid" xml:"openid" form:"openid"`

	Code   *string `validate:"" gorm:"comment:'对应code'" query:"code" json:"code" xml:"code" form:"code"`
	Index  *string `validate:"" gorm:"comment:'对应下标'" query:"index" json:"index" xml:"index" form:"index"`
	Region *string `validate:"" gorm:"comment:'省,市,区'" query:"region" json:"region" xml:"region" form:"region"`
}

type Users struct {
	User

	Password  string `json:"-"`
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetUserTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*Users, error) {
	var users []*Users
	err := db.Preload(clause.Associations).Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func GetUser(maps interface{}) (*User, error) {
	var user User
	err := db.Preload(clause.Associations).Where(maps).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func AddUser(user User) error {
	// 开始事务
	tx := db.Begin()
	if err := db.Create(&user).Error; err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		return err
	}
	// 否则，提交事务
	tx.Commit()
	return nil
}

func EditUser(id int, data User) error {
	// 开始事务
	tx := db.Begin()
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		return err
	}
	// 否则，提交事务
	tx.Commit()
	return nil
}

func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// 根据id判断test 对象是否存在
func ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)

	return user.ID > 0
}
