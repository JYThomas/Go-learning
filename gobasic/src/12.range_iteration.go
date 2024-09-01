package main
import "fmt"

func main(){
	// 定义一个map 用于循环迭代
	intmap := make(map[string]int)
	intmap["A"] = 90
	intmap["B"] = 80
	intmap["C"] = 70

	for k, v := range intmap{
		if(k == "A"){
			fmt.Printf("%s, %d\n", k, v)
		}
	}
}