package main

import (
	"fmt"
	// "log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

var db *pgx.ConnPool
var db_err error

func init() {
	// connecting database of psql by go
	db, db_err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Database: "pmp",
			User:     "ashutosh",
			Password: "123",
		},
		MaxConnections: 5, // number of connections from client{application} to server{Database}
	})

	// database can't be conn
	if db_err != nil {
		fmt.Println("Can't connect to database")
		// log.Fatalf("Error : %#v", db_err)
		os.Exit(1)
	}

}

func main() {
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		var request Request
		c.BindJSON(&request)
		fmt.Printf("\n\nRequest received : %#v\n\n", request)

		// database transaction begins

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var user_id int32 //variable for storing user_id

		// transaction query for inserting user_name and name and it will return user_id

		tx.QueryRow(`
            INSERT INTO alpha (user_name, name) values ($1, $2) returning user_id
        `, request.Username, request.Name).Scan(&user_id)

		commit_err := tx.Commit() // commit the transaction

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "Duplicate insertion") // commit error should be mentioned according to the corresponding errors
			return
		}

		var response Response
		response.UserID = user_id
		response.DeviceID = "12Asd@#"
		response.Timestamp = time.Now().Unix()
		c.JSON(200, response)
	})

	fmt.Print("Going live on :7070\n")
	r.Run(":7070")
}

// Request holds
type Request struct {
	Username string `json:"user_name,omitempty"` // username
	Name     string `json:"name,omitempty"`
}

//Response
type Response struct {
	UserID    int32  `json:"user_id,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}
