/*
Create: 2022/8/28
Project: Monica
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"github.com/JJApplication/fushin/utils/env"
)

var (
	Monica     = "Monica"
	Address    string
	Port       int
	Token      string
	RemoteAddr string
	SSHUser    string
	SSHPassWD  string
)

func init() {
	loader := env.EnvLoader{}
	Address = loader.Get("Address").Raw()
	Port = loader.Get("Port").MustInt(6789)
	Token = loader.Get("Token").Raw()
	RemoteAddr = loader.Get("RemoteAddr").Raw()
	SSHUser = loader.Get("SSHUser").Raw()
	SSHPassWD = loader.Get("SSHPassWD").Raw()
}
