package zentaoapi

import (
	"fmt"
	"testing"
)

func TestGetSession(t *testing.T) {
	Login()
}

var (
	product    int    = 4
	project    int    = 5
	module     int    = 334
	assignedTo string = "liangkai"
	bugType    string = "codeerror"
)

func TestCreateBug(t *testing.T) {
	params := BuildParams(product, project, module, assignedTo, bugType, "这是一个测试ok...")
	ret, err := New(product, 0, 0, params)
	if err != nil {
		fmt.Printf("  err:%v", err)
		return
	}
	fmt.Printf("ret = %s\n", ret)

}
