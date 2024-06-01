package _case

import "fmt"

/* 算术运算符 */
func ArithmeticCase() {
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("a+b 的值为 %d\n", c) //a+b 的值为 31
	c = a - b
	fmt.Printf("a-b 的值为 %d\n", c) // a-b 的值为 11
	c = a * b
	fmt.Printf("a*b 的值为 %d\n", c) // a*b 的值为 210
	c = a / b
	fmt.Printf("a/b 的值为 %d\n", c) //a/b 的值为 2
	c = a % b
	fmt.Printf("a%%b 的值为 %d\n", c) //a%b 的值为 1
	a++
	fmt.Printf("a++ 的值为 %d\n", a) //a++ 的值为 22
	a--
	fmt.Printf("a-- 的值为 %d\n", a) //a-- 的值为 21
}

/* 关系运算符 */
func RelationCase() {
	var a = 21
	var b = 10
	fmt.Println("a == b", a == b) // a == b false
	fmt.Println("a > b", a > b)   // a > b true
	fmt.Println("a < b", a < b)   // a < b false
	fmt.Println("a >= b", a >= b) // a >= b true
	fmt.Println("a <= b", a <= b) // a <= b false
	fmt.Println("a != b", a != b) // a != b true
}

/* 逻辑运算符 */
func LogicCase() {
	var a = true
	var b = false
	fmt.Println("a && b", a && b)       // a && b false
	fmt.Println("a||b", a || b)         // a||b true
	fmt.Println("!(a && b)", !(a && b)) // !(a && b) true
}

/* 位运算 */
func BitCase() {
	var a uint8 = 60
	var b uint8 = 13
	var c uint8 = 0
	// 二进制打印
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00000000
	fmt.Println()

	c = a & b
	fmt.Println("按位与运算 a & b: 同1为1 其他就为0")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00001100
	fmt.Println()

	c = a | b
	fmt.Println("按位或运算 a | b: 有任意一个为1就为1 其他就为0")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00111101
	fmt.Println()

	c = a ^ b
	fmt.Println("按位异或运算 a ^ b: 两个不相同为1 两个相同为0")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00110001
	fmt.Println()

	c = a << 2
	fmt.Println("左移运算符 a << 2: 两个不相同为1 两个相同为0")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", c) //11110000
	fmt.Println()

	c = a >> 2
	fmt.Println("右移运算符 a >> 2: 两个不相同为1 两个相同为0")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", c) //00001111
	fmt.Println()

	c = ^a
	fmt.Println("^a:")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", c) //11000011
	fmt.Println()

	c = a &^ b
	fmt.Println("a &^ b:")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00110001
	fmt.Println()

	c = a & ^b
	fmt.Println("a & ^b:")
	fmt.Printf("%08b\n", a) //00111100
	fmt.Printf("%08b\n", b) //00001101
	fmt.Printf("%08b\n", c) //00110001
	fmt.Println()
}

/* 赋值运算 */
func AssignmentCase() {
	var a = 21
	var c int
	c = a
	fmt.Println("c = a，c值为：", c) // 21
	c += a
	fmt.Println("c += a，c值为：", c) //42
	c -= a
	fmt.Println("c -= a，c值为：", c) // 21
	c *= a
	fmt.Println("c *= a，c值为：", c) //441
	c /= a
	fmt.Println("c /= a，c值为：", c) //21
	c %= a
	fmt.Println("c %= a，c值为：", c) // 0

	var b uint8 = 60
	fmt.Printf("b 值为 %d，二进制表示： %08b\n", b, b) // 60  00111100
	// 2   00000010
	b <<= 2
	fmt.Printf("b <<= 2， b 值为 %d，二进制表示： %08b\n", b, b) // 240  11110000
	b >>= 2
	fmt.Printf("b >>= 2， b 值为 %d，二进制表示： %08b\n", b, b) // 60  00111100
	b &= 2
	fmt.Printf("b &= 2， b 值为 %d，二进制表示： %08b\n", b, b) // 0   00000000
	b |= 2
	fmt.Printf("b |= 2， b 值为 %d，二进制表示： %08b\n", b, b) // 2  00000010
	b ^= 2
	fmt.Printf("b ^= 2， b 值为 %d，二进制表示： %08b\n", b, b) // 0  00000000
}
