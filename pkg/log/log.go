package log

import "fmt"

func RowMsg(msg ...interface{}) {
	for _, obj := range msg {
		fmt.Println(obj)
	}
}
