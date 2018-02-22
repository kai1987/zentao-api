package zentaoapi

import (
	"fmt"
	"testing"
)

var (
	product    int    = 4
	project    int    = 5
	module     int    = 334
	assignedTo string = "liangkai"
	bugType    string = "codeerror"
)

func TestCreateBug(t *testing.T) {
	Login()
	steps := `
	aaaa
	bbbb
	eeee
		eeeeffff
	`
	params := BuildParamsFull(product, project, module, 3, 3, 0, 0, assignedTo, bugType, "这是一个测试apiaa", steps, "linux", "chrome", "red", "2018-02-23", "")
	ret, err := New(product, 0, 0, params)
	if err != nil {
		fmt.Printf("  err:%v", err)
		return
	}
	fmt.Printf("ret = %s\n", ret)

}
