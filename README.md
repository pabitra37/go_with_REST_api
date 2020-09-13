# SIMPLE GOLANG MICROSERVICE USING REST API #
## Decription: Here I have created a simple structure of an Icecream shop using go lang microservices with RESTful services. We can perform various CRUD operations on the structure using PUT, POST, DELETE and GET services.

## Quick start
 Build the docker image using:
 ```
 docker build -t ice-cream-shop .
 ```
 Run it using
 ```
 docker run -p 8000:8000 -it ice-cream-shop
 ```
<b>`NOTE: I used postman to send https requests to the server` <b>
   
<b>` Please use VS code to fully utilize this README. md file` <b>

## Endpoints 

Get all Icecream flavours
```
GET url/flavours
```
Get a single Icecream flavour
```
GET url/flavour/{flavourID}
```
Delete an Icecream flavour
```
DELETE url/flavour/{flavourID}
```
Create an Icecream Flavour
```
POST url/flavours
    Request sample
    {
        "flavourid": 4545454,
        "name":"choclate",
        "price": 35.75,
        "serving": "cup"
    }
```
Update an Icecream Flavour
```
PUT url/flavour/{flavourID}
    Request sample
    {
        "flavourid": 4545454,
        "name":"choclate",
        "price": 35.75,
        "serving": "cup"
    }
```