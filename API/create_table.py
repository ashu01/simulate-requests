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
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


