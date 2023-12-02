package utils

import (
	"fmt"
)

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
