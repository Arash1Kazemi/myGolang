package main

import "fmt"

func main() {
	// days := []string{"SHANBE", "YRCKSHANBE", "DOSHANBE", "SESHANBE"}
	// fmt.Println("start")

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(days[i])
	// }

	// for j := range days {
	// 	fmt.Println(days[j])
	// }
	// for i,d := range days    {
	// 	fmt.Printf("index is %v value is %v\n",i,d)
	// }
	var i int = 0
	for i < 10 {
		if i == 5 {
			fmt.Println("i")
			i++
			continue
		}
		fmt.Println("m")
		i++
	}

}
