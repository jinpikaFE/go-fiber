package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Test struct {
	Model

	// query tag是query参数别名，json xml，form适合post
	Name string `validate:"required,min=3,max=32" query:"name" json:"name" xml:"name" form:"name"`
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetTestTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Test{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetTests(pageNum int, pageSize int, maps interface{}) ([]*Test, error) {
	var tests []*Test
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tests).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tests, nil
}

func AddTest(test *Test) error {
	if err := db.Create(&test).Error; err != nil {
		return err
	}

	return nil
}

func EditTest(id int, data interface{}) error {
	if err :=db.Model(&Test{}).Where("id = ?", id).Updates(data).Error;err!= nil {
		return err
	}

	return nil
}

func DeleteTest(id int) error {
	if err := db.Where("id = ?", id).Delete(Test{}).Error; err!= nil {
		return err
	}

	return nil
}

// 根据id判断test 对象是否存在
func ExistTestByID(id int) bool {
	var test Test
	db.Select("id").Where("id = ?", id).First(&test)

	return test.ID > 0
}

// gorm所支持的回调方法：

// 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
// 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
// 删除：BeforeDelete、AfterDelete
// 查询：AfterFind

func (test *Test) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (test *Test) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
