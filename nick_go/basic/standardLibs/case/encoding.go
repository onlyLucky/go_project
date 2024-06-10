package _case

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func EncodingCase() {
	type user struct {
		ID   int64
		Name string
		Age  uint8
	}
	u := user{ID: 1, Name: "nick", Age: 18}
	// json序列化
	bytes, err := json.Marshal(u)
	fmt.Println("json序列化: ", bytes, err) // json序列化:  [123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 110 105 99 107 34 44 34 65 103 101 34 58 49 56 125] <nil>
	u1 := user{}
	// json反序列化
	err = json.Unmarshal(bytes, &u1)
	fmt.Println("json反序列化: ", u1, err) // json反序列化:  {1 nick 18} <nil>

	// base64编解码
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("base64编解码: ", str) // base64编解码:  eyJJRCI6MSwiTmFtZSI6Im5pY2siLCJBZ2UiOjE4fQ==
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println("base64编解码: ", bytes1, err) // base64编解码:  [123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 110 105 99 107 34 44 34 65 103 101 34 58 49 56 125] <nil>

	// 16进制编解码
	str1 := hex.EncodeToString(bytes1)
	fmt.Println("16进制编解码: ", str1) // 16进制编解码:  7b224944223a312c224e616d65223a226e69636b222c22416765223a31387d
	bytes2, err := hex.DecodeString(str1)
	fmt.Println("16进制编解码: ", bytes2, err) // 16进制编解码:  [123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 110 105 99 107 34 44 34 65 103 101 34 58 49 56 125] <nil>
}
