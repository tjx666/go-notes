package main

import "fmt"

// 声明变量约定俗成使用小驼峰
// 全局变量没有被使用也不会报错
var globalVariable int

// go 允许声明全局变量，函数， 接口等，但是不允许写语句，语句应该被封装到函数中
// globalVariable = 666; // syntax error: non-declaration statement outside function body

var standardDefinition = 999

// !: 短变量声明不允许在全局环境下声明，老老实实使用 var 声明
// shortDefinition := 666 // syntax error: non-declaration statement outside function body

// ------------------------ 标识符 -----------------------------
// 和 js 一样，Go 的标识符只能以 下划线，字母，数字组成，并且只能以下划线和字母开头
// js 由于私有属性的存在其实还可以以 # 开头
var _underScore int
var letter int
var call10086 int

// ------------------------批量声明  -----------------------------
var (
	v1 int
	v2 int
	v3 int
)

/*
var v4 int, v5 int
=> syntax error: unexpected comma after top level declaration
*/

/*
var v4 int v5 int
=> syntax error: unexpected v5 after top level declaration
*/

/*
var v4 int
v5 int
=> syntax error: non-declaration statement outside function body
*/

// ------------------------ 变量的初始化 -----------------------------
func testInitValue() {
	// 变量在声明的时候会根据变量的类型分配对应大小的内存，并初始化为对应类型的初始值
	// int 初始化为 0，string 初始化为空字符串，bool 类型初始化为 false，切片、函数、指针变量的默认为nil。
	var it int
	var str string
	fmt.Println(it, str) // => 0

	// 可以声明变量的同时初始化
	var num int = 666
	// !: 函数中的变量必须被使用，不像全局变量，声明不使用会报：num declared but not used
	fmt.Println(num) // => 666

	// 多个变量的声明和初始化
	var v1, v2 = 1, 2
	fmt.Println(v1, v2) // => 1 2

	// 声明了几个变量就必须给几个初始值
	// var v3, v4 = 3 // 报错：assignment mismatch: 2 variables but 1 values

	// 类型推导
	// 比较新的编程语言都有些共同特征，例如 Kotlin，Go，TypeScript，类型声明放后面，支持类型推导
	// var integer = 1
	// var ssr = "ShadowSocksR"
	// ssr = integer // 报错：cannot use integer (type int) as type string in assignment

	// 短变量声明
	// 可以使用 := 声明并初始化变量，其实就是省的你写 var 了
	// !: 只能在函数中使用
	words := "c.c is my wife!" // 等同于 var words = "c.c is my wife!"
	fmt.Println(words)
	words = "Rem is my wife!"
	// 说明了短变量声明的是可变的变量
	fmt.Println(words) // => Rem is my wife!

	// 匿名变量：使用下划线作为变量名
	// 匿名变量不会分配内存
	var _, ele2 = 5, 6
	fmt.Println(ele2) // => 6
	// _ 不能被使用
	// fmt.Println(_) // 报错：cannot use _ as value
}

// ------------------------ 常量 -----------------------------
func testConst() {
	const PI = 3.14159265457
	radius := 2.0
	fmt.Printf("Area is %f\n", PI*radius*radius) // => Area is 12.566371

	// 常量必须初始化
	//  const E int // const declaration cannot have type without expression

	// 批量声明常量
	const (
		HOST = "127.0.0.1"
		PORT = 80
	)
	fmt.Printf("Address is http://%s:%d\n", HOST, PORT) // => Address is http://127.0.0.1:80

	// 批量声明常量时，如果前面初始化了，后续不初始化的变量会最后一个初始化的变量值保持一致
	const (
		c1 = 0
		c2 = 1
		c3     // 1
		c4     // 1
		c5 = 5 // 5
		c6     // 5
	)

	fmt.Println(c1, c2, c3, c4, c5, c6) // => 0 1 1 1 5 5
}

// ------------------------ iota -----------------------------
func testItoa() {
	// 发音：/aɪˈoʊtə/
	// 这个词就是全拼，可以理解为 for 循环中的 i 下标，表示迭代索引的意思
	//参考：https://stackoverflow.com/questions/31650192/whats-the-full-name-for-iota-in-golang
	// 在 go 语言这个关键字用于常量计数，只能用于常量声明中
	const constant = iota
	fmt.Println(constant) // => 0

	// iota 每次到一个新的常量声明表达式会置 0
	const (
		c1 = iota
		c2 // 1
		c3 // 2
		c4 // 3
	)
	fmt.Println(c1, c2, c3, c4)

	// 可以使用 - 跳过一些值
	// iota 后面每跟一行声明，iota 的值就 +1
	const (
		c5 = iota
		_
		c6
	)
	fmt.Println(c5, c6) // => 0 2

	// 插队声明
	const (
		n1 = iota // 0
		n2 = 2    // 100
		n3 = iota // 2
		n4        // 3
	)
	const n5 = iota
	fmt.Println(n1, n2, n3, n4, n5) // => 0 2 2 3 0

	// 一个应用：定义数量级
	// 常量不使用不会报错
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	// iota 定义在一行的时候，此时 iota 是相同值
	// e 和 f 的值说明了 iota 值 +1 是根据声明的行数来计算的
	const (
		a, b = iota + 1, iota + 2 // 1 2
		c, d                      // 2 3
		e, f = iota, iota         // 2 2
	)
	fmt.Println(a, b, c, d, e, f)
}

func main() {
	testInitValue()
	testConst()
	testItoa()
}
