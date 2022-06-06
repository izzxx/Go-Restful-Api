# Products API

A simple RESTful API service that I made, you can see the various products that are already available on this service. You can use it to learn how to fetch APIs from third parties.
to be able to access it you must be registered as a user on this api service.

Read the instructions below to register, then you can use it.

## Authentication
The token will always be used on every request to the url, save the token in the header. <br />`Authorization: Bearer <token>`

#### Register
* POST ```https://firstproductapi.herokuapp.com/api/v1/users/register```

```json
{
    "username": "useraname",
    "email": "email@gmail.com",
    "password": "password"
}
```

#### Login
* POST ```https://firstproductapi.herokuapp.com/api/v1/users/login```

```json
{
    "email": "email@gmail.com",
    "password": "password"
}
```

#### Response
```json
{
  "status_code": 200,
  "message": "success",
  "data": {
    "id": "09d20d85-742f-4158-89c4-11f9c196cefd",
    "email": "email@gmail.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoidXNlcm5hbWUiLCJFbWFpbCI6ImVtYWlsQGdtYWlsLmNvbSIsIklzQWRtaW4iOmZhbHNlLCJleHAiOjE2NTQzNTI3NDV9.AAcYHpVmmc052mEcBzViQumdZua2pNMvKgGZoRgW-8I"
  }
}
```

## Products


### User
The access you get when you are registered as a user.

###### Get All Products
* GET ```https://firstproductapi.herokuapp.com/api/v1/products```

###### Get Product By Id
* GET ```https://firstproductapi.herokuapp.com/api/v1/products/{id}```

```json
{
  "status_code": 200,
  "message": "success get product",
  "data": {
    "id": "741585071-8",
    "name": "Juice - Apple, 341 Ml",
    "price": 502234.403,
    "quantity": 394,
    "created_at": "2021-12-08T23:23:16Z"
  }
}
```

### Admin
Only admins can access this page to make data changes

| Method | Url |
| --- | --- |
| POST | `https://firstproductapi.herokuapp.com/api/v1/products/admin/create` |
| PUT | `https://firstproductapi.herokuapp.com/api/v1/products/admin/update/{id}`|
| DELETE | `https://firstproductapi.herokuapp.com/api/v1/products/admin/delete/{id}`|

## Closing 

I'm sure there are still many shortcomings in my project.
If you have criticism or suggestions, I am very open to accept them.
Thank you🤙.
