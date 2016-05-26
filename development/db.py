import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'pmp', host = 'localhost', user = 'ashutosh', password = '123')
cursor = postgres.cursor()

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'alpha')
""") 
# it returns true if there is at least row in database and it will be a tuple
presence = cursor.fetchone()[0]
print(presence)

if presence:
    print('alpha table exists, deleting here')
    cursor.execute("""
        drop table alpha
    """)

cursor.execute("""
    create table alpha(
        user_name text, 
        user_id serial,
        name text
    )
""")

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)
    

for i in range(5):
    user_name = random.choice(["abs", "xyz", "vrt"])
    name = random.choice(["Ashu", "pp"])
    
    message_send = {
        "user_name" : user_name,
        "name" : name
    }
    resp = requests.post("http://localhost:7070/send", data = json.dumps(message_send))
    body = resp.text
    print(body)
    # cursor.execute("""   
    #     INSERT INTO alpha (user_name, name) values (%s, %s)
    # """, (user_name, name))
	
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)

# cursor.execute("""   
#    select * from alpha 
# """)

# a = cursor.fetchall()
# for row in a : 
#     print(row)
