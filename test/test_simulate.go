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
		// Dump the incoming request data into a struct variable
		c.BindJSON(&request)
		fmt.Println("\n\n\n Received : %#v\n\n\n", request)
		// database transaction begins

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var userid int64 //variable for storing user_id
		var deviceid string
		var platform string
		// transaction query for inserting user_name and name and it will return user_id

		// tx.QueryRow(`
		//     INSERT INTO alpha (user_name, name) values ($1, $2) returning user_id
		// `, request.Username, request.Name).Scan(&user_id)

		// Fetch the user_id for the corresponding input request
		fmt.Println("Username and email Randomly generated from test_simulate.py : ", request.Username, request.Email)
		// db.QueryRow(`
		// 		SELECT userid
		// 		FROM users
		// 		WHERE username = $1 and email = $2`, request.Username, request.Email).Scan(&userid)
		// fmt.Print("User id after query : ", userid)

		// // Select the device corresponding to the above user_id
		// err := db.QueryRow(`
		// 		SELECT deviceid, platform
		// 		FROM usersdescription
		// 		WHERE usersdescription.userid = $1`, userid).Scan(&deviceid, &platform)

		// natural join of above two commands
		err := db.QueryRow(`
                SELECT usersdescription.deviceid, usersdescription.platform
                FROM users, usersdescription
                WHERE usersdescription.userid = users.userid AND username = $1 AND email = $2
        `, request.Username, request.Email).Scan(&deviceid, &platform)

		fmt.Println("\n\nDeviceid and platform : ", deviceid, platform)
		fmt.Println(err)

		// commit_err := tx.Commit() // commit the transaction

		// if commit_err != nil {
		// 	tx.Rollback()
		// 	c.JSON(500, "Duplicate insertion") // commit error should be mentioned according to the corresponding errors
		// 	return
		// }

		var response Response
		response.UserID = userid
		response.DeviceID = deviceid
		response.Timestamp = time.Now().Unix()
		response.Platform = platform
		fmt.Println("\n\n\n Response :  %#v\n\n\n", response)
		c.JSON(200, response)
	})

	// r.POST("/gogs", func(c *gin.Context) {
	// 	m := make(map[string]interface{})
	// 	c.BindJSON(&m)
	// 	fmt.Printf("\n\n\nReceived new push : %#v\n\n\n\n", m)
	// 	c.JSON(200, "ok")
	// })

	fmt.Print("Going live on :7070\n")
	r.Run(":7070")
}

// Request holds
type Request struct {
	Username string `json:"username,omitempty"` // username
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
}

//Response
type Response struct {
	UserID    int64  `json:"userid,omitempty"`
	DeviceID  string `json:"deviceid,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Platform  string `json:"platform,omitempty"`
}
