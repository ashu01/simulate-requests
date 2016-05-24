import requests

message_to_send = {
    "route" : "push",
    "recently_used" : {
        "platform" : "ios",
        "device_id" : "gdjg#5r67fgvehjdg@7781uigd"
    },
    "device_list" : [
        {
            "platform" : "ios",
            "device_id" : "iljilwefhl%^%^ghfgwekuqeg2"
        },
        {
            "platform" : "android",
            "device_id" : "iouyoh6587123468ykjqhdkqdg"
        },
        {
            "platform" : "windows",
            "device_id" : "hi876234uwhedgkjwqhgkfuegkequfg"
        }
    ],
    "content" : "This is the push notification content",
    "timeout_after" : 5,    # This is the timeout (in seconds) after which the push needs to be sent to other devices in the device_list
    "priority" : ["ios", "android"]   # Priority platforms
}

resp = requests.post("http://localhost:7000/send", json=message_to_send)
body = resp.text
print(body)