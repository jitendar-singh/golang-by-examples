package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.March, 14, 02, 00, 00, 00, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
