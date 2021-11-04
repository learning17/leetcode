package main
// https://learnku.com/docs/the-way-to-go/switch-structure/3594
import (
	"fmt"
)

func main() {
	switch1()
}

/*
switch 条件表达式 {
case 常量表达式1:
    语句 1
case 常量表达式2:
    语句 2
    .
    .
    .
case 常量表达式n:
    语句 n
default:
    语句 n+1
}
（1）计算条件表达式的值value
（2）如果value满足某条case语句，则执行该语句，执行完跳出switch语句
（3）如果value不满足所有的case语句：
（3.1）如果有default，则执行该语句，执行完跳出switch语句
（3.2）如果没有default，则直接跳出switch语句
如果switch关键字后面没有条件表达式，则必须在case语句中进行条件判断，即类似于 if else if 语句
*/

func getValue() int {
	return 100
}

func switch1() {
	var num int = 100
	switch num {
	case 100:
		fmt.Println("== 100")
	default:
		fmt.Println("not 100")
	}

	switch {
	case num == 100:
		fmt.Println("== 100")
	default:
		fmt.Println("not 100")
	}

	switch value := getValue(); {
	case value == 100:
		fmt.Println("== 100")
	default:
		fmt.Println("not 100")
	}

	var a interface{}
	a = 10
	switch v := a.(type) {
	case int:
		fmt.Println("int:",v)
	default:
		fmt.Println("not int")
	}

}
