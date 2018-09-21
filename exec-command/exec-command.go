/*
* @Author: Newbie Coder
* @Date:   2018-09-21 23:52:34
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-21 23:52:38
*/
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}