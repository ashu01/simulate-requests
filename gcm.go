package main

import(
        "fmt"
        "github.com/gin-gonic/gin"
        "github.com/google/go-gcm"
)

type HttpMessageLocal struct {
    To                  string                  `json:"to,omitempty"`
    Timestamp 		    int64  		  			`json:"timestamp,omitempty"`
	Priority            string        			`json:"priority,omitempty"`
	Devices				[]map[string]string 	`json:"devices,omitempty"`
}

func main(){
    r := gin.Default()
    r.POST("/gcmpush", func(c *gin.Context) {
        var httpMessageLocal HttpMessageLocal
        var httpMessage gcm.HttpMessage
        
        c.BindJSON(&httpMessageLocal)
        httpMessage.To = httpMessageLocal.To
        fmt.Println("200")
       fmt.Println(httpMessage.To)
        httpMessage.Priority = httpMessageLocal.Priority
        fmt.Println(httpMessage.Priority)
        var apiKey string
        
        var registration_ids []string
        
        for _,val := range httpMessageLocal.Devices{  
            registration_ids = append(registration_ids, val["DeviceID"])        
        }
        
        for _, val := range httpMessageLocal.Devices{
            if val["Platform"] == "Android"{
                    apiKey = "AIzaSyBfI5t4-GW5VovfzQ6BpvhTd2dkUB7L9R"
                    httpMessage.Notification.Title = "Android Notification"
                    httpMessage.Notification.Body = "Successfully sent a Push Notification to an Android Device."
                    httpMessage.RegistrationIds = registration_ids
                    fmt.Println("Http message is", httpMessage)
                    gcm.SendHttp(apiKey,httpMessage)
            } 
        }
         
    })
    fmt.Print("GCM connecting on Port :8000")
    r.Run(":8000")
}