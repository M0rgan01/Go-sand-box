package controllers

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/logger"
	"github.com/morgan/Go-sand-box/todo-project/models"
)

func SetupWebsocketRoutes(engine *gin.RouterGroup) {
	hub := models.NewHub()
	go hub.Run()
	r := engine.Group("/ws")
	{
		r.GET("/chat", func(context *gin.Context) {
			ServeWs(hub, context)
		})
	}
}

// ServeWs serveWs handles websocket requests from the peer.
func ServeWs(hub *models.Hub, c *gin.Context) {
	conn, err := models.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	client := &models.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
