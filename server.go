package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// Database connectivity variables
var db *pgx.ConnPool
var db_err error

//Initialise connection to the database
func init() {
	db, db_err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Database: "pmp",
			User:     "ashutosh",
			Password: "123",
			Port:     5432,
		},
		MaxConnections: 10,
	})

	if db_err != nil {
		fmt.Println("Can't connect to database")
	}
}

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

	r.POST("/createuserid", func(c *gin.Context) {
		var request UserIDCreate
		c.BindJSON(&request)

		fmt.Println("\n\nEmail Received : %#v\n\n", request)

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function
		var response UserIDResp

		// insert into users table
		tx.QueryRow(`
        INSERT INTO users (username, name, email) values ($1, $2, $3) returning userid
    `, request.UserName, request.Name, request.Email).Scan(&response.Userid)
		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}

		c.JSON(200, response)

	})

	r.GET("/api/verifyemail", func(c *gin.Context) {
		// receive userid and map it with the table users and get email
		var userid UserIDResp
		c.BindJSON(&userid)

		if userid.Userid <= 0 {
			response_map := make(map[string]string)
			response_map["error"] = "Invalid Userid"
			c.JSON(404, response_map)
			return
		}

		var email string
		db.QueryRow(`
			SELECT email 
			FROM users
			WHERE userid = $1
		`, userid.Userid).Scan(&email)

		if email == "" {
			response_email := make(map[string]string)
			response_email["error"] = "Userid Not found"
			c.JSON(403, response_email)
			return
		}

		email_map := make(map[string]string)
		email_map["email"] = email
		c.JSON(200, email_map)
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

// Request for creating userid for received email

type UserIDCreate struct {
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"username,omitempty"`
}

type UserIDResp struct {
	Userid int64 `json:"userid,omitempty"`
}
