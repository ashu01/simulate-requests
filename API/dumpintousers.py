import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'pmp', host = 'localhost', user = 'ashutosh', password = '123')
cursor = postgres.cursor()

a = ["abs", "xyz", "vrt", "alpha", "beta"]
b = {"abs":"as01hu@gmail.com","xyz":"xyz@outlook.com","vrt": "alpha@yahoomail.com", "alpha":"xwe@yahoo.in", "beta":"ade@gmail.com"}

for i in range(5):
    username = random.choice(a)
    name = random.choice(["Ashu", "Pk"])
    email = b[username]
    a.remove(username)
    message_send = {
        "username" : username,
        "name" : name,
        "email" : email
    }
    resp = requests.post("http://localhost:7000/dumpdataintousers", data = json.dumps(message_send))
    body = resp.text
    print(body)
    

