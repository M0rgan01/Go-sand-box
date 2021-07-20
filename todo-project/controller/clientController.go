package controller

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/model"
	"log"
)

// ServeWs serveWs handles websocket requests from the peer.
func ServeWs(hub *model.Hub, c *gin.Context) {
	conn, err := model.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &model.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
