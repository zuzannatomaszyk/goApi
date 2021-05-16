
Installation
=============
Requirements:
 ========
* ```docker-compose```
* ```git```

Instructions:
 ========
1. Clone repositiry using ```git clone https://github.com/zuzannatomaszyk/goApi.git```
2. Build application using ```docker compose build```
3. Start the containers using ```docker compose up -d``` 

Server should be running on http://localhost:8081 



About goApi
=============
goApi is a simple chat server. It runs on http://localhost:8081 and supports the following REST API: 

 1. GET /messages

     list 100 most recent messages, sorted by 'timestamp' posted to the chat server.

     example:

     ```
     curl -H "Content-Type: application/json" http://localhost:8081/messages

     {
       "messages": [
         {"timestamp": 1491345710.18, "user": "superman", "text": "hello"},
         {"timestamp": 1491345713.18, "user": "batman", "text": "hello"}
       ]
     }

     ```

 2. POST /message 

     a request to post the given message. 
     when the message is processed by the server a unix timestamp is recorded with each message.

     example:

     ```
     curl -X POST -H "Content-Type: application/json" --data '{"user":"superman", "text":"hello"}' http://localhost:8081/message

     {
       "ok": true
     }
     ```

 3. GET /users

     a request to return a set of users that have posted messages so far.

     example:

     ```
     curl -H "Content-Type: application/json" http://localhost:8081/users

     {
       "users": [
         "superman", "batman"
       ]
     }
     ```

