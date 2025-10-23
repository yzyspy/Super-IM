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

// Value 实现 driver.Valuer 接口
func (v VerificationQuestion) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Scan 实现 sql.Scanner 接口
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
