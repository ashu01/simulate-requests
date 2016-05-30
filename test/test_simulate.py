import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'pmp', host = 'localhost', user = 'ashutosh', password = '123')
cursor = postgres.cursor()

#   generate simulation module for users{relation} here user_id is unique and 
#   it will be used for the purpose of mapping from user_description{relation} and
#   this relation will have the information of user details like which device is 
#   used recently and corresponding to that it will send the push notifications
#   In user_description relation we will create a primary key user_id and on the 
#   basis of this very primary key we will map to the users and after mapping we
#   will send the push notification to the concerned device with the appropriate 
#   server i.e. {use apns for apple, gcm for android etc}.
 
 
 ######################################################################################
 ######################################################################################
 #                               CHECK FOR TABLE USERS                                #
 ######################################################################################
 ###################################################################################### 
 
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'users')
""") 
# it returns true if there is at least row in database and it will be a tuple
presence = cursor.fetchone()[0]
print(presence)



 ######################################################################################
 ######################################################################################
 #                               DELETE TABLE USERS IF EXISTS                         #
 ######################################################################################
 ######################################################################################
 
if presence:
    print('users table exists, deleting here')
    cursor.execute("""
        drop table users
    """)



 ######################################################################################
 ######################################################################################
 #                               CREATE TABLE USERS                                   #
 ######################################################################################
 ######################################################################################
cursor.execute("""
    create table users( 
        userid bigserial,
        username text,
        name text,
        email text
    )
""")
a = ["abs", "xyz", "vrt", "alpha", "beta"]
b = {"abs":"as01hu@gmail.com","xyz":"xyz@outlook.com","vrt": "alpha@yahoomail.com", "alpha":"xwe@yahoo.in", "beta":"ade@gmail.com"}

for i in range(5):
    username = random.choice(a)
    name = random.choice(["Ashu", "pp"])
    email = b[username]
    a.remove(username)
    cursor.execute("""
        INSERT INTO users (username, name, email) values(%s, %s, %s)
    """,(username, name, email))

######################################################################################
######################################################################################
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
######################################################################################
######################################################################################

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


 ######################################################################################
 ######################################################################################
 #                               CHECK FOR TABLE USERSDESCRIPTION                     #
 ######################################################################################
 ######################################################################################
 
cursor.execute("""
   select exists(select 0 from pg_class where relname = 'usersdescription')
""") 
# it returns true if there is at least row in database and it will be a tuple
presence = cursor.fetchone()[0]
print(presence)


 ######################################################################################
 ######################################################################################
 #                       DELETE USERSDESCRIPTION IF EXISTS                            #
 ######################################################################################
 ######################################################################################
 
 
if presence:
    print('usersdescription table exists, deleting here')
    cursor.execute("""
        drop table usersdescription
    """)

 ######################################################################################
 ######################################################################################
 #                       CREATE TABLE USERSDESCRIPTION                                #
 ######################################################################################
 ######################################################################################

cursor.execute("""
    create table usersdescription( 
            userid bigint,
            deviceid text,
            platform text
            -- CONSTRAINT userdescription_pk PRIMARY KEY (userid)
    )
""")

######################################################################################
######################################################################################
#                               DUMP RANDOM DATA                                     #
######################################################################################
######################################################################################

for i in range(15):
    userid = i+1
    deviceid = random.choice(["gdjg#5r67fgvehjdg@7781uigd", "iljilwefhl%^%^ghfgwekuqeg2", "iouyoh6587123468ykjqhdkqdg"])
    platform = random.choice(["ios", "apple", "windows"]) 
    
    cursor.execute("""
        INSERT INTO usersdescription (userid, deviceid, platform) values(%s, %s, %s)
    """,(userid, deviceid, platform))
    
##########          check what is generated          ###############################
    
#####################################################################################
#####################################################################################    
# cursor.execute("""   
#    select * from usersdescription 
# """)
#####################################################################################
#####################################################################################


######################################################################################
######################################################################################
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
######################################################################################
######################################################################################

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)
    
a = ["abs", "xyz", "vrt", "alpha", "beta"]
b = {"abs":"as01hu@gmail.com","xyz":"xyz@outlook.com","vrt": "alpha@yahoomail.com", "alpha":"xwe@yahoo.in", "beta":"ade@gmail.com"}

for i in range(5):
    username = random.choice(a)
    name = random.choice(["Ashu", "pp"])
    email = b[username]
    a.remove(username)
    message_send = {
        "username" : username,
        "name" : name,
        "email" : email
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
