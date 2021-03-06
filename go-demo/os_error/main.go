/*
 * File: \test\os_error\main.go                                                *
 * Project: go-demo                                                            *
 * Created At: Monday, 2022/05/23 , 15:27:36                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:34:54                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.IsExist(nil))
	s, err1 := os.Stat("storage/uploads")
	fmt.Println(s)
	fmt.Println(os.IsNotExist(err1))
	fmt.Println("-------------------------------")
	f, err2 := os.Stat("/workspaces/Golang_Starting/go_start/go-demo/os_error/main.go")
	fmt.Println(f)
	fmt.Println(os.IsNotExist(err2))
}
