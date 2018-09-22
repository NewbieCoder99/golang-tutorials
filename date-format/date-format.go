/*
* @Author: Newbie Coder
* @Date:   2018-09-22 11:18:21
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 11:20:10
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	DateNow := time.Now().Format("02-01-2006")
	fmt.Println("Date now is ", DateNow)
}