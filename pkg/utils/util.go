package utils

import (
	"fmt"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"os"
)

func AAA() {

	var a []error
	testArr(&a)
	fmt.Print(a)

}

func testArr(a *[]error) {
	*a = append(*a, fmt.Errorf("reg err"))
}

func GetPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.RowMsg(err)
	}
	return dir
}
