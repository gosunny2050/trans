package interface_utils

import (
	"fmt"
	"testing"
)

// 测试结构体
type TestStruct1 struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Title string `json:"title"`
	Info   string `json:"Info"`
	Value string `json:"value"`
}
type TestStruct2 struct {
	ttt int
	Id       string `json:"id"`
	Name     string `json:"name"`
	Title string `json:"title"`
	Info   string `json:"Info"`
	Value string `json:"value"`
}
func TestInterfaceToStruct(t *testing.T) {
	// string 类型转struct
	str2 := `{"id":"0","name":"test","info":"","title":"test","value":"application"}`
	var tmp1 TestStruct1;
	var tmp2 TestStruct1;
	var tmp3 TestStruct1;
	err := InterfaceToStruct(str2, &tmp1)
	fmt.Println(err)
	fmt.Println(tmp1)
	// struct类型转struct，只能识别共同字段,不区分大小写
	stru := TestStruct2{
		ttt:1,
		Id:        "1",
		Name:      "zhangsan",
		Title:  "test",
		Info:  "vkeer",
		Value: "1",
	}
	err = InterfaceToStruct(stru, &tmp2)
	fmt.Println(err)
	fmt.Println(tmp2)
	//map类型转struct，只识别共同字段，不区分大小写
	 maps := map[string]interface{}{
	 	"title":"test",
	 	"sex":"man",
	 	"Name":"zhaosan",
	 	"info":"test2",
	 }
	err = InterfaceToStruct(maps, &tmp3)
	fmt.Println(err)
	fmt.Println(tmp3)
}
func TestInterfaceToBool(t *testing.T) {
	 bolvalue  := true
	 strvalue  := "true"
	fmt.Println(InterfaceToBool(true))
	fmt.Println(InterfaceToBool(&bolvalue))
	fmt.Println(InterfaceToBool(1))
	fmt.Println(InterfaceToBool(int8(1)))
	fmt.Println(InterfaceToBool(int16(1)))
	fmt.Println(InterfaceToBool(int32(1)))
	fmt.Println(InterfaceToBool(int64(1)))
	fmt.Println(InterfaceToBool("1"))
	fmt.Println(InterfaceToBool("true"))
	fmt.Println(InterfaceToBool(&strvalue))
	fmt.Println(InterfaceToBool(1.1))
}
func TestInterfaceToInt(t *testing.T) {
	fmt.Println(InterfaceToInt(1))
	fmt.Println(InterfaceToInt(1.1))
	fmt.Println(InterfaceToInt(1.5555))
	fmt.Println(InterfaceToInt("111"))
	fmt.Println(InterfaceToInt(true))
}
func TestInterfaceToString(t *testing.T) {
	fmt.Println(InterfaceToString("1432144"))
	fmt.Println(InterfaceToString(1))
	fmt.Println(InterfaceToString(int8(1)))
	fmt.Println(InterfaceToString(int16(1)))
	fmt.Println(InterfaceToString(int32(1)))
	fmt.Println(InterfaceToString(int64(1)))
	fmt.Println(InterfaceToString(1.1111111421))
	fmt.Println(InterfaceToString("1.1"))
	fmt.Println(InterfaceToString(true))
	fmt.Println(InterfaceToString(false))
}
func TestInterfaceToInt64(t *testing.T) {
	fmt.Println(InterfaceToInt64("1432144"))
	fmt.Println(InterfaceToInt64(1))
	fmt.Println(InterfaceToInt64("1.1"))
	fmt.Println(InterfaceToInt64(true))
	fmt.Println(InterfaceToInt64(false))
}
func TestInterfaceToFloat32(t *testing.T) {
	fmt.Println(InterfaceToFloat32(3.1415926))
	fmt.Println(InterfaceToFloat32("1432144.15"))
	fmt.Println(InterfaceToFloat32(1))
	fmt.Println(InterfaceToFloat32("1.1"))
	fmt.Println(InterfaceToFloat32(true))
	fmt.Println(InterfaceToFloat32(false))
}
func TestInterfaceToFloat64(t *testing.T) {
	fmt.Println(InterfaceToFloat64("1432144"))
	fmt.Println(InterfaceToFloat64(3.1415926))
	fmt.Println(InterfaceToFloat64(1))
	fmt.Println(InterfaceToFloat64("1.1"))
	fmt.Println(InterfaceToFloat64(true))
	fmt.Println(InterfaceToFloat64(false))
}