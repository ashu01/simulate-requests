import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'pmp', host = 'localhost', user = 'ashutosh', password = '123')
cursor = postgres.cursor()

number = [1, 2, 3, 4, 5] #userid
 # DEVICE ID 
a = ["asdf", "lkjhgt", "qwdfgbn", "poijhg", "sdfgh", "qwsdfpkj", "74jhg", "15kjhg", "35iuhg", "asdf95", "sdfg955", "pojhb3"]


b = { "asdf":"ios", "lkjhgt":"ios", "qwdfgbn":"ios", "poijhg":"android", "sdfgh":"android", "qwsdfpkj":"android", "74jhg":"window", "15kjhg":"window", "35iuhg":"window", "asdf95":"blackberry", "sdfg955":"blackberry", "pojhb3":"blackberry"}    
for i in range(6):
    print(i)
    userid = random.choice(number)
    deviceid = random.choice(a)
    platform = b[deviceid]
    a.remove(deviceid)
    
    message = {
        "userid":userid,
        "deviceid":deviceid,
        "platform":platform
    }
    print(message)
    resp = requests.post("http://localhost:7000/dumpintousersdescription", data = json.dumps(message))
    body = resp.text
    print(body)