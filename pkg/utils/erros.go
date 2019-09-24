package utils

import (
	"fmt"
)

func LogErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
func HandleErr(err error, f *[]error) {
	if err != nil {
		*f = append(*f, err)
	}
}
