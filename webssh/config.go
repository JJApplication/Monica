/*
Create: 2022/8/29
Project: Monica
Github: https://github.com/landers1037
Copyright Renj
*/

// Package webssh
package webssh

type WebSSHConfig struct {
	RemoteAddr string
	User       string
	Password   string
	AuthModel  AuthModel
	PkPath     string
}
