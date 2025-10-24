package pwd

import (
	"fmt"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	is_valid := CheckPwd("$2a$04$7ote8n3Zjd9OdyFk2HdutO1sXq1Tk.csb8Ret1v/db7yaqeS2zJ5G", "123456")
	fmt.Printf("CheckPwd result: %v\n", is_valid)
}

func TestGenHashPwd(t *testing.T) {
	pwd := HashPwd("123456")
	fmt.Printf("GenHashPwd1 result: %v\n", pwd)

}
