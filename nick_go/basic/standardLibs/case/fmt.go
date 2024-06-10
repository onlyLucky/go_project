package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	// 打印到标准输出
	fmt.Println("今天天气很好") // 今天天气很好
	// 格式化，并打印到标准输出
	fmt.Printf("%s天气很好", "今天、明天") // 今天、明天天气很好
	// 格式化
	str := fmt.Sprintf("%s天气很好", "今天、明天")
	// 输出到 io.write
	fmt.Fprint(os.Stderr, str) // 今天、明天天气很好
	fmt.Println()
}

// 占位符
func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}
	// 通用占位符
	fmt.Printf("默认格式的值：%v \n", a)    // 默认格式的值：{10}
	fmt.Printf("包含字段名的值：%+v \n", a)  //包含字段名的值：{value:10}
	fmt.Printf("go语法表示的值：%#v \n", a) // go语法表示的值：_case.simple{value:10}
	fmt.Printf("go语法表示的类型：%T \n", a) // go语法表示的类型：_case.simple
	fmt.Printf("输入字面上的百分号：%%10 \n")  // 输入字面上的百分号：%10

	// 整数占位符
	v1 := 10
	v2 := 20170                           // "今"字码点值
	fmt.Printf("二进制：%b \n", v1)           // 二进制：1010
	fmt.Printf("Unicode 码点转字符：%c \n", v2) // Unicode 码点转字符：今
	fmt.Printf("十进制：%d \n", v1)           // 十进制：10
	fmt.Printf("八进制：%o \n", v1)           // 八进制：12
	fmt.Printf("0o 为前缀的八进制：%O \n", v1)    // 0o 为前缀的八进制：0o12
	fmt.Printf("用单引号将字符的值包起来：%q \n", v2)  // 用单引号将字符的值包起来：'今'
	fmt.Printf("十六进制：%x \n", v1)          // 十六进制：a
	fmt.Printf("十六进制大写：%X \n", v1)        // 六进制大写：A
	fmt.Printf("Unicode 格式：%U \n", v2)    // Unicode 格式：U+4ECA

	// 宽度设置
	fmt.Printf("%v 的二进制： %b; go语法表示二进制为： %#b; 指定二进制宽度为8， 不足8位补0：%08b \n", v1, v1, v1, v1) // 10 的二进制： 1010; go语法表示二进制为： 0b1010; 指定二进制宽度为8， 不足8位补0：00001010
	fmt.Printf("%v 的十六进制： %x; 使用go语法表示为，指定十六进制宽度为8， 不足8位补0：%#08x \n", v1, v1, v1)         // 10 的十六进制： a; 使用go语法表示为，指定十六进制宽度为8， 不足8位补0：0x0000000a
	fmt.Printf("%v 的字符为： %c; 指定宽度为5位，不足5位补空格：%5c \n", v2, v2, v2)                         // 20170 的字符为： 今; 指定宽度为5位，不足5位补空格：    今

	// 浮点数占位符
	var f1 = 123.789
	var f2 = 12345678910.7899
	fmt.Printf("指数为2的幂的无小数科学计数法： %b \n", f1)                             // 指数为2的幂的无小数科学计数法： 8710876473008849p-46
	fmt.Printf("科学计数法： %e \n", f1)                                       // 科学计数法： 1.237890e+02
	fmt.Printf("科学计数法,大写： %E \n", f1)                                    // 科学计数法,大写： 1.237890E+02
	fmt.Printf("有小数而无指数,及常规的浮点数格式，默认宽度和精度： %f \n", f1)                   // 有小数而无指数,及常规的浮点数格式，默认宽度和精度： 123.789000
	fmt.Printf("宽度为9，精确默认： %9f \n", f1)                                  // 宽度为9，精确默认： 123.789000
	fmt.Printf("默认宽度，精确度保留两位小数：%.2f \n", f1)                             // 默认宽度，精确度保留两位小数：123.79
	fmt.Printf("宽度为9，精确度保留两位小数： %9.2f \n", f1)                           // 宽度为9，精确度保留两位小数：    123.79
	fmt.Printf("宽度为9，精确度保留两位小数,缺少位补0： %09.2f \n", f1)                    // 000123.79
	fmt.Printf("宽度为9，精确度保留0位小数： %9.f \n", f1)                            // 宽度为9，精确度保留0位小数：       124
	fmt.Printf("根据情况自动选%%e 或 %%f 来输出，以产生更紧凑的输出（末位无0）： %g %g \n", f1, f2) // 根据情况自动选%e 或 %f 来输出，以产生更紧凑的输出（末位无0）： 123.789 1.23456789107899e+10
	fmt.Printf("根据情况自动选%%E 或 %%f 来输出，以产生更紧凑的输出（末位无0）： %G %G \n", f1, f2) // 根据情况自动选%E 或 %f 来输出，以产生更紧凑的输出（末位无0）： 123.789 1.23456789107899E+10
	fmt.Printf("以十六进制方式表示： %x \n", f1)                                   // 以十六进制方式表示： 0x1.ef27ef9db22d1p+06
	fmt.Printf("以十六进制方式表示，大写： %X \n", f1)                                // 以十六进制方式表示，大写： 0X1.EF27EF9DB22D1P+06

	/* 字符串占位符 */
	var str = "今天是个好日子"
	fmt.Printf("%s \n", str) // 今天是个好日子
	// 双引号包裹
	fmt.Printf("%q \n", str) // "今天是个好日子"
	// 16进制表示
	fmt.Printf("%x \n", str) // e4bb8ae5a4a9e698afe4b8aae5a5bde697a5e5ad90
	// 以空格作为两数之间的分隔符，并用大写16禁止
	fmt.Printf("%X \n", str) // E4BB8AE5A4A9E698AFE4B8AAE5A5BDE697A5E5AD90

	/* 指针占位符 */
	var str1 = "今天是个好日子"
	bytes := []byte(str1)
	// 切片第0个元素的地址
	fmt.Printf("%p \n", bytes) // 0xc0000133b0
	mp := make(map[string]int, 0)
	fmt.Printf("%p \n", mp) // 0xc00007a7e0
	var p *map[string]int = new(map[string]int)
	fmt.Printf("%p \n", p) // 0xc000060070
}
