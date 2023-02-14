package utilityFunc

import "fmt"

func StartApp(){
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	fmt.Println(string(colorYellow), "Started Application listening on 8080", string(colorReset))
}