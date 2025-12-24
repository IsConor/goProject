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

func main() {
	arrayDemo()
}
