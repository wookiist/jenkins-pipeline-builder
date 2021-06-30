package main

import (
	"flag"
	"fmt"
)

func main() {
	taskName := flag.String("name", "example-project", "name of the task")
	fmt.Println(*taskName)
}
