# Authentication
Authentication

CURL
curl --location 'localhost:8084/v1/daijai/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstname":"Sutida",
    "lastname":"ratanajaruen",
    "email":"deerS@hotmail.com",
    "password":"123456",
    "role":"user"
}'

curl --location 'localhost:8084/v1/daijai/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"deerS@hotmail.com",
    "password":"123456"
}'

curl --location 'localhost:8084/user/v1/daijai/all' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUwMDk2OTUsInVzZXJJZCI6Nn0.S1WaqNbqEqIM-rx7s3bs2wjYxIAO5_5Owsu_Ix50L6A' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ5ODIyODYsInVzZXJJZCI6NX0.IxU7eYOCS8tQPYrGyTdlOhnRkw6lUQ5SinXUIvZI-64'

curl --location 'localhost:8084/user/v1/daijai/profile' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUwNjQ4MDUsInVzZXJJZCI6IjA3OTc5ODcyLTZjYTgtNDdkZS05OThhLTJlMjBkNzNjYzk0NSJ9.EovFNSAcicMTtn-Y09h0jFYEauqYOFvz_dEVRmSX18E' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ5ODc5NjIsInVzZXJJZCI6NX0.X6Gm8daNbcfTSaMTDR3rG-vCgun2Rbgu58CVBN5V-J4'