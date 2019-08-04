# go-gin-mongodb
Example of RESTful API with Go, Gin and mongodb

API:


Welcome page:
GET
https://xgowebapp.herokuapp.com/


Todos:

List TODO items:
GET
https://xgowebapp.herokuapp.com/todos

Create todo item:
POST
https://xgowebapp.herokuapp.com/todos

`
{
    "ID": "5d470a6a4cd71a2870b33765",
    "Description": "Call mom",
    "Labels": [
        {
            "ID": "5d470a6a4cd71a2870b33766",
            "Name": "home"
        },
        {
            "ID": "5d470a6a4cd71a2870b33767",
            "Name": "work"
        }
    ]
}
`
