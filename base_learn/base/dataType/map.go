package dataType

import "fmt"

var TypeMap = 1;

func init() {
	var initNum = 456
	fmt.Println("map init",initNum)
}

func MakeMap(){
	fmt.Println("MakeMap")
}

func GetTypeMap(){
	fmt.Println("GetTypeMap: ",TypeMap)
}