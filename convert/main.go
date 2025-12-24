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

func main() {
	//arrayDemo()
	sliceDemo()
}
