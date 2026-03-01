package main
import  (
	"fmt"
)
func sums(nums ... int) int{
	total := 0
	for _, v := range nums{
		total += v
	}
	return  total
}

func main(){
	results := sums(1, 3, 4)
	fmt.Println(results)
}
