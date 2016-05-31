import requests
import json
import random

a = ["abs", "xyz", "vrt", "alpha", "beta"]
b = {"abs":"as01hu@gmail.com","xyz":"xyz@outlook.com","vrt": "alpha@yahoomail.com", "alpha":"xwe@yahoo.in", "beta":"ade@gmail.com"}

for i in range(5):
    username = random.choice(a)
    name = random.choice(["Ashu", "pp"])
    email = b[username]
    #print(username, name, email)
    message_send = {
        "username" : username,
        "name" : name,
        "email" : email
    }
    resp = requests.post("http://localhost:7000/createuserid", data = json.dumps(message_send))
    body = resp.text
    print(body)
    body = json.loads(body) 
    userid = body["userid"]
    uid_send ={
        "userid" : userid
    }
    re = requests.get("http://localhost:7000/api/verifyemail", data = json.dumps(uid_send))
    validate = re.text
    print(validate)
    a.remove(username)