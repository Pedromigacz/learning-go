package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	strings.Split("Pedro", "")              // It does't print returned values
	fmt.Println(strings.Split("Pedro", "")) // Prints theresult of the function

}
