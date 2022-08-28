/*
Create: 2022/8/28
Project: Monica
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"fmt"

	"github.com/JJApplication/fushin/errors"
)

func main() {
	errors.Try(func() {
		loadConfig()
		s := initServer()
		s.Run()
	}).Catch(func(exception interface{}) {
		fmt.Printf("error exit, %v", exception)
	})
}
