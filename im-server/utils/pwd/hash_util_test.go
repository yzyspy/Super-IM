package pwd

import (
	"fmt"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	hashedPwd := HashPwd("123456") //数据库保存hash之后的值

	is_valid := CheckPwd(hashedPwd, "123456")
	fmt.Printf("CheckPwd result: %v\n", is_valid)
}

func TestGenHashPwd(t *testing.T) {
	pwd := HashPwd("1111")
	fmt.Printf("GenHashPwd1 result: %v\n", pwd)

	pwd2 := HashPwd("1111")
	fmt.Printf("GenHashPwd2 result: %v\n", pwd2)
}
