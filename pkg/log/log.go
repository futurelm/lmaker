package log

import (
	"fmt"
)

func RowMsg(msg ...interface{}) {
	for _, obj := range msg {
		//	bytes,_:= json.Marshal(obj)
		//	fmt.Println(string(bytes))
		fmt.Println(obj)
	}
}
