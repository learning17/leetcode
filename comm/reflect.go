package main
// https://juejin.cn/post/6844903559335526407


import (
	"fmt"
	"reflect"
)

func main() {
	f2()
}

type User struct {
	Id int
	Name string
	Age int
}

func f1() {
	var num float64 = 1.2345
	fmt.Println("type:", reflect.TypeOf(num))
	fmt.Println("value:", reflect.ValueOf(num))
	
	value := reflect.ValueOf(num)
	pointer := reflect.ValueOf(&num)
	converPointer := pointer.Interface().(*float64)
	converValue := value.Interface().(float64)
	fmt.Println(converPointer)
	fmt.Println(converValue)
}

func f2() {
	user := User{1, "wang", 12}
	getType := reflect.TypeOf(user)
	fmt.Println("get Type is:", getType.Name())

	getValue := reflect.ValueOf(user)
	fmt.Println("get value is:", getValue)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

}
