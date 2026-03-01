package main
import (
	"fmt"
)

func sum(nums  ...int) int{
	total := 0
	for _, v := range nums{
		total += v
	}
	return total
}

func main(){
	results := sum(1, 2, 3, 4)
	fmt.Println(results)
}


