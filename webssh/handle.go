/*
Create: 2022/8/28
Project: Monica
Github: https://github.com/landers1037
Copyright Renj
*/

// Package webssh
package webssh

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/JJApplication/fushin/log/private"
	fushinServer "github.com/JJApplication/fushin/server/http"
	"github.com/gorilla/websocket"
)

type WebSSHConfig struct {
	RemoteAddr string
	User       string
	Password   string
	AuthModel  AuthModel
	PkPath     string
}

type WebSSH struct {
	*WebSSHConfig
}

func NewWebSSH(conf *WebSSHConfig) *WebSSH {
	return &WebSSH{
		WebSSHConfig: conf,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (w WebSSH) ServeConn(c *fushinServer.Context) {
	// 校验authcode码
	auth := c.Query("auth")
	if auth == "" {
		c.AbortWithStatus(407)
	}
	if auth != os.Getenv("Token") {
		c.AbortWithStatus(407)
	}
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	defer wsConn.Close()
	var config *SSHClientConfig
	switch w.AuthModel {
	case PASSWORD:
		config = SSHClientConfigPassword(
			w.RemoteAddr,
			w.User,
			w.Password,
		)
	case PUBLICKEY:
		config = SSHClientConfigPulicKey(
			w.RemoteAddr,
			w.User,
			w.PkPath,
		)
	}

	client, err := NewSSHClient(config)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer client.Close()

	turn, err := NewTurn(wsConn, client)

	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer turn.Close()

	var logBuff = bufPool.Get().(*bytes.Buffer)
	logBuff.Reset()
	defer bufPool.Put(logBuff)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := turn.LoopRead(logBuff, ctx); err != nil {
			client.Close()
			private.Log.Error(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		if err := turn.SessionWait(); err != nil {
			client.Close()
			private.Log.Error(err.Error())
		}
		cancel()
	}()
	wg.Wait()
}
