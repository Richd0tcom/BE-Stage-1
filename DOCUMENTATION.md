# API Documentation

## Overview
This api is capable of CRUD operations on a "user" resource. It makes use of Golang while interfacing with a postgres database.
All requests for these operations are made to the host `https://backend-stage-1.onrender.com`

### Create a User
To create a user, a `POST` request is made to the `/api/users/` endpoint, with the request body
```json
    { "name": "Richdotcom6"}

```
You should receive a JSON response with the `id`  and `name` fields of the created user.

e.g.
```json
{
    "id": "d5f632fe-c3d4-45e5-acc7-a1e6562e138f",
    "name": "Richdotcom6"
}

```

### Fetch a User
To fetch a user, a `GET` request is made to the `/api/users/:user_id` endpoint. You should receive a JSON response with the `id`  and `name` fields of the fetched user.

e.g.
Request: 
> GET https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f

Response:
```json
{
    "id": "d5f632fe-c3d4-45e5-acc7-a1e6562e138f",
    "name": "Richdotcom6"
}
```
Dynamic parameter handling: You can also fetch a user by their name by passing the `name` field in the query string like so:
> https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f?name=Richdotcom6


### Update a User
To update a users name, a `PUT` request is made to the `/api/users/:user_id` endpoint, with the request body
```json
    { "new_name": "Richdotcom12"}
``` 
You should receive a JSON response with the `id`  and new `name` fields of the updated user.

e.g.
Request: 
> PUT https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f

Response:
```json
{
    "id": "d5f632fe-c3d4-45e5-acc7-a1e6562e138f",
    "name": "Richdotcom12"
}
```
Dynamic parameter handling: You can also update a user by their name by passing the `name` field in the query string like so:
> PUT https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f?name=Richdotcom6

Note the request body is the same as above


### Delete a User
To update a users name, a `DELETE` request is made to the `/api/users/:user_id` endpoint, 
You should receive a JSON response like so:

```json 
{}
```

e.g.
Request: 
> DELETE  https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f

Response:
```json 
{}
```
Dynamic parameter handling: You can also delete a user by their name by passing the `name` field in the query string like so:
> DELETE https://backend-stage-1.onrender.com/api/users/d5f632fe-c3d4-45e5-acc7-a1e6562e138f?name=Richdotcom6


## UML Diagram