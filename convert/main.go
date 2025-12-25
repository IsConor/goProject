package main

import "fmt"

func arrayDemo() {
	fmt.Println("\n=== 数组示例 ===")

	// 声明数组
	var arr [5]int
	fmt.Printf("声明但不初始化：%v\n", arr)

	// 初始化数组
	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("初始化: %v\n", arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	// 自动推断长度
	arr2 := [...]int{1, 2, 3}
	fmt.Printf("自动长度: %v, 长度: %d\n", arr2, len(arr2))
}

func sliceDemo() {
	fmt.Println("\n=== 切片示例 ===")

	// 声明切片，跟声明一个数组的区别，就是 声明数组要有初始化长度，而切片不需要
	var slice []int
	fmt.Printf("空切: %v, nil:%t\n", slice, slice == nil)

	// 使用make创建切片
	slice = make([]int, 3, 5) // 长度3，容量5 什么叫容量5，就是初始化的slice，也就是这个动态数组的当前最大容量是5个，如果超过会自动扩容
	fmt.Printf("make创建: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 直接初始化
	slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("初始化: %v\n", slice)

	// 追加元素
	slice = append(slice, 6)
	fmt.Printf("追加后: %v, 容量: %d\n", slice, len(slice))

	// 切片截取
	subSlice := slice[1:3]
	fmt.Printf("切片[1:3]: %v\n", subSlice)

	// 切片共享底层数组， 切片后subSlice, 和slice共享同一个底层数组，所以修改subSlice会影响slice
	subSlice[0] = 999
	// 如果在切片末尾追加，那么切片会自动扩容，并且影响原来数组
	subSlice = append(subSlice, 888)
	fmt.Printf("修改了subSlice后: %v\n", slice)
}

func mapDemo() {
	fmt.Println("\n=== 映射示例 ===")

	// 声明映射
	var m map[string]int
	fmt.Printf("声明但是不初始化: %v, nil: %t\n", m, m == nil)

	// 使用make初始化
	m = make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	fmt.Printf("添加后: %v\n", m)

	// 直接初始化
	m2 := map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 8,
	}
	fmt.Printf("直接初始化: %v\n", m2)

	// 读取值
	value := m2["apple"]
	fmt.Println("apple的值: %d\n", value)
	value2 := m2["orange"]
	fmt.Printf("orange的值: %d\n", value2)

	m2["orange"] = 100
	fmt.Printf("orange的值: %d\n", m2["orange"])

	// 检查key是否存在
	value, ok := m2["grape"]
	fmt.Printf("grade存在: %t, 值: %d\n", ok, value)

	valueA, okA := m2["apple"]
	fmt.Printf("apple存在: %t, 值: %d\n", okA, valueA)
}

func pointerDemo() {
	fmt.Println("\n=== 指针示例 ===")

	x := 10
	fmt.Printf("x的值为: %d\n", x)

	// 获取地址
	p := &x
	fmt.Printf("x的地址: %p\n", p)
	fmt.Printf("指针的值: %d\n", *p)

	m := p
	fmt.Printf("m: %p\n", m)
	fmt.Printf("m指向的值: %d\n", *m)

	// 通过指针修改值
	*p = 20
	fmt.Printf("修改后x的值为: %d\n", x)

	// 指针作为函数参数
	increment(&x)
	fmt.Printf("函数修改后的x的值为: %d\n", x)

	// nil指针
	var ptr *int
	fmt.Printf("nil指针: %v\n", ptr)

	fmt.Println("\n=== 值传递 vs 指针传递 ===")

	// 值传递示例
	a := 10
	fmt.Printf("调用前a=%d\n", a)
	valuePass(a)
	fmt.Printf("调用后 a = %d (值未改变)\n", a)

	// 指针传递示例
	b := 10
	fmt.Printf("\n调用前 b = %d\n", b)
	pointerPass(&b)
	fmt.Printf("调用后 b = %d(值已改变)\n", b)

	// 详细说明
	fmt.Println("\n关键区别")
	fmt.Println("1. 值传递：函数接收的是值的副本，修改副本不影响原来的值")
	fmt.Println("2. 指针传递：函数接收的是地址，通过地址可以直接修改原来的值")
}

func valuePass(num int) {
	fmt.Printf(" 函数内接受到的值: %d\n", num)
	num = 100 // 修改副本，不影响赋值
	fmt.Printf(" 函数内修改后: %d\n", num)
}

func pointerPass(num *int) {
	fmt.Printf(" 函数内接收到的地址: %p\n", num)
	fmt.Printf(" 函数内接收到的值: %d\n", *num)
	*num = 100 // 通过指针修改原值
	fmt.Printf(" 函数内修改后: %d\n", *num)
}

func increment(ptr *int) {
	*ptr++
}

func main() {
	//arrayDemo()
	//sliceDemo()
	//mapDemo()
	pointerDemo()
}
