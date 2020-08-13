package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month(), int(time.Now().Month()))
	fmt.Println(time.Now().Day(), int(time.Now().Day()))
	fmt.Println(time.Now().Weekday(), int(time.Now().Weekday()))
}
