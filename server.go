/*
Create: 2022/8/28
Project: Monica
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"github.com/JJApplication/fushin/log/private"
	"github.com/JJApplication/fushin/server/http"
	"github.com/JJApplication/monica/webssh"
)

// new a websocket server

func loadConfig() {
	private.Log.InfoF("load config:\nAddress: %s Port: %d Token: %s RemoteAddr: %s PublicKey: %s",
		Address,
		Port,
		Token,
		RemoteAddr,
		PublicKey)
}

func initServer() *http.Server {
	serve := &http.Server{
		EnableLog:    true,
		Logger:       nil,
		Debug:        false,
		RegSignal:    nil,
		Address:      http.Address{Host: Address, Port: Port},
		Headers:      nil,
		Copyright:    Monica,
		MaxBodySize:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
		PProf:        false,
	}

	serve.Init()
	addHandler(serve)
	return serve
}

func addHandler(s *http.Server) {
	config := &webssh.WebSSHConfig{
		RemoteAddr: RemoteAddr,
		User:       SSHUser,
		Password:   SSHPassWD,
		AuthModel:  webssh.PUBLICKEY,
		PkPath:     PublicKey,
	}
	handle := webssh.NewWebSSH(config)
	s.Route(http.GET, "/ws", handle.ServeConn)
	s.Route(http.GET, "/", func(c *http.Context) {
		c.ResponseStr(200, "Hello Monica")
	})
}
