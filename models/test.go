package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Test struct {
	Model

	// query tag是query参数别名，json xml，form适合post
	Name string `query:"name" json:"name" xml:"name" form:"name"`
}

func GetTests(pageNum int, pageSize int, maps interface{}) (tests []Test) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tests)

	return
}

func AddTest(test *Test) bool {
	db.Create(&test)

	return true
}

func EditTest(id int, data interface{}) bool {
	db.Model(&Test{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteTest(id int) bool {
	db.Where("id = ?", id).Delete(Test{})

	return true
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
