package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
	"github.com/parnurzeal/gorequest"
)

var db *pgx.ConnPool
var db_err error
var response Response

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
		log.Fatal(db_err)
	}
}

func main() {
	r := gin.Default()
	r.POST("/dumpdataintousers", func(c *gin.Context) {
	var rqst Request
	c.BindJSON(&rqst)
	fmt.Println("\n\n\n Received : %#v\n\n\n", rqst)

	rows, err := db.Query(`
                SELECT usersdescription.deviceid, usersdescription.platform
                FROM users, usersdescription
                WHERE usersdescription.userid = users.userid AND username = $1 AND email = $2
    `, rqst.Username, rqst.Email)
		
	if err != nil {
           log.Fatal(err)
    }
    defer rows.Close()
	
	var deviceid, platform string
	var devices []map[string]string
	
	for rows.Next() {
            if err := rows.Scan(&deviceid, &platform); err != nil {
                    log.Fatal(err)
            }
			
			device := make(map[string]string)
			device["DeviceID"] = deviceid
			device["Platform"] = platform
			
			devices = append(devices, device)
	}
    
	if err := rows.Err(); err != nil {
            log.Fatal(err)
    }
				
		response.To = rqst.Username
		response.Timestamp = time.Now().Unix()
		response.Priority = "high"
	    response.Devices = devices
		
		fmt.Println("\n\n\n Response :  %#v\n\n\n", response)
		
		response_map := make(map[string]string)
		response_map["response"] = "Success"
	    c.JSON(200, response_map)
		
	request := gorequest.New()
    resp, body, errs := request.Post("http://localhost:8000/gcmpush").
	    Set("Notes", "GO request is coming").
		Send(response).
		Send( `{"request" : "made"}` ).
		End()
	fmt.Print(resp, body, errs)
		
	})
	
	
	// fmt.Print("The Gorequest resp : \n", resp, "\n body : ",  body, "\n errs : ", errs, "\n")
	fmt.Print("Going live on :7000\n")
	r.Run(":7000")
	
}

type Request struct {
	Username 			string 		  			`json:"username,omitempty"` 
	Name     			string 		  			`json:"name,omitempty"`
	Email   		    string 		  			`json:"email,omitempty"`
}


type Response struct {
	To                  string        					   `json:"to,omitempty"`
	Timestamp 		    int64  		  					   `json:"timestamp,omitempty"`
	Priority            string        					   `json:"priority,omitempty"`
	Devices     		[]map[string]string                `json:"devies,omitempty"`
}


