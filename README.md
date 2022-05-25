#  Notification Service

##  About
This Project is a test by Xendit, a very simple notification-service. The system uses PostgreSQL as the database. The system can handle API to create a Url Webhook, Test to receive notification, and receive notification, etc. I also use Unit Tests in this system. To make sure user has an easier debugging experience, I implements logging in the system.

##  How to Use  
###  Database
User can set their own database string connection address by editing the file `.env`.

###  APIs
To run this project user need to execute file main.go with command `go run main.go` through the terminal. After executing the command, in order to hit the APIs, user can use API Platform software like Postman, etc.