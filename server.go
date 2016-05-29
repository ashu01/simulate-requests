package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		var request Request
		c.BindJSON(&request)
		fmt.Printf("\n\nReceived new Push:\n\n%#v\n\n", request)
		var response Response
		response.NotificationID = int64(100001221082108)
		response.Timestamp = time.Now().Unix()
		c.JSON(200, response)
	})

	r.POST("/gogs", func(c *gin.Context) {
		m := make(map[string]interface{})
		c.BindJSON(&m)
		fmt.Printf("\n\n\nReceived new push : \n\n\n %#v\n\n\n\n", m)
		c.JSON(200, "ok")
	})
	fmt.Printf("Going live on :7000")
	r.Run(":7000")
}

// Request holds the incoming request for a push notification.
type Request struct {
	Route        string   `json:"route,omitempty"`
	RecentlyUsed device   `json:"recently_used,omitempty"`
	DeviceList   []device `json:"device_list,omitempty"`
	Content      string   `json:"content,omitempty"`
	Timeout      int64    `json:"timeout_after,omitempty"`
	Priority     []string `json:"priority,omitempty"`
}

type device struct {
	Platform string `json:"platform,omitempty"`
	DeviceID string `json:"device_id,omitempty"`
}

// Response is the struct that holds the notification ID and the UNIX timestamp of when it was sent
type Response struct {
	NotificationID int64 `json:"notification_id,omitempty"`
	Timestamp      int64 `json:"timestamp,omitempty"`
}
