package ctype

type VerificationQuestion struct {
	Question1 *string `gorm:"column:question" json:"question1"` //验证问题
	Answer1   *string `gorm:"column:answer" json:"answer1"`     //验证问题的答案
	Question2 *string `gorm:"column:question" json:"question2"` //验证问题
	Answer2   *string `gorm:"column:answer" json:"answer2"`     //验证问题的答案
	Question3 *string `gorm:"column:question" json:"question3"` //验证问题
	Answer3   *string `gorm:"column:answer" json:"answer3"`     //验证问题的答案
}
