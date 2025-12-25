package main

import "fmt"

func ifDemo() {
	fmt.Println("=== if/else示例 ===")

	age := 20
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	// 多条件判断
	scores := []int{95, 85, 70, 50}
	for index, score := range scores {
		fmt.Printf("索引%d: %d\n", index, score)
		if score >= 90 {
			fmt.Printf("分数%d: 优秀\n", score)
		} else if score >= 80 {
			fmt.Printf("分数%d: 良好\n", score)
		} else if score >= 60 {
			fmt.Printf("分数%d: 及格\n", score)
		} else {
			fmt.Printf("分数%d: 不及格\n", score)
		}
	}
}

func main() {
	ifDemo()
}
