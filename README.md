# SpotlasTest
Recruitment test

## Prerequisite
Name the database for <b>test_db</b> for the go programs to work. If you don't wan't to do that change <b>DB_NAME</b> in server.go. Switch off the passwords for postgresql.

## Task 1
* Log into postgres user by sudo -u postgres -i
* Enter

    ```psql```
 
    You will see the postgres prompt
 
* Create a database by 

    ```create database <database name>```

    Choose ```test_db``` if it has to work for the go programs.


* Connect to the database by entering

    ```\c <database name>```

*  Execute the script for populating the table ```"MY_TABLE"```. This file was given with the test material.

    ```\i <path to spots.sql>```

* Execute the script which is part of the candidate response
 
    ```\i <absolute path to taskOne.sql>```

    This will update the table ```"MY_TABLE"``` modifying the field ```website``` as required in point number one in Task 1. It will execute other queries
    as required in the task
  
  

## Task 2
Execute ```taskTwoEndPoint.go``` to initialize an endpoint at ```localhost:5000/``` . Enter the command below in the directory.


```go run taskTwoEndPoint.go```

The database name is hard coded in the file ```taskTwoEndPoint.go``` as

```DB_NAME := "test_db"```

If you wan't to change the db name go to the ```taskTwoEndPoint.go``` file and change the value of ```DB_NAME```.

Password is not used at all. You may wan't to disable the password for the postgresql.
    

There is a client program called which can be used to test the end point

```go run client.go```
