package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type VerificationQuestion struct {
	Question1 *string `gorm:"column:question" json:"question1"` //验证问题
	Answer1   *string `gorm:"column:answer" json:"answer1"`     //验证问题的答案
	Question2 *string `gorm:"column:question" json:"question2"` //验证问题
	Answer2   *string `gorm:"column:answer" json:"answer2"`     //验证问题的答案
	Question3 *string `gorm:"column:question" json:"question3"` //验证问题
	Answer3   *string `gorm:"column:answer" json:"answer3"`     //验证问题的答案
}

// Value 实现 driver.Valuer 接口, driver.Valuer 用于将自定义类型转换为数据库支持的类型（写入数据库）。
func (v VerificationQuestion) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Scan 实现 sql.Scanner 接口， 当从数据库读取数据时，GORM或标准库的database/sql会调用该类型的Scan()方法，
// 将数据库返回的数据（可能是[]byte或string等）转换回自定义类型。
//
// 例如，如果数据库存储的是JSON字符串，你可以在Scan()方法中将其反序列化到自定义结构体。
func (v *VerificationQuestion) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, v)
}

// GormDataType 指定 GORM 数据类型
func (VerificationQuestion) GormDataType() string {
	return "json"
}
