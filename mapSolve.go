package main

import "fmt"

func main() {
	names := []string{"liugi", "nonso", "zuluu", "major", "NEWLINE"}

	for i, v := range  names {
		if i == 2 {
			fmt.Printf("continue at pos: %v \n", i)
			continue
		}
		//if i > 3 {
		//	fmt.Println(v)
		//	break
		//}
		fmt.Printf("Index:%v Value:%v \n", i, v)
	}

}