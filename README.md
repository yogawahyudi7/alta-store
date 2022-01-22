# ToDoList App
- e-commerce app is an application used to sell products online


Run project with: 
```
go run main.go
```

## Stack-tech :dart:
- [x] RESTful API Using Go, Echo, Gorm, MySQL
- [x] AWS for service api

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register`
* Login : `POST /users/login/`

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose Token is provided with the request:

- Get user profile data by User ID : `GET /users/profile/:id`
- Update user data by User ID : `PUT /users/:id`
- Delete user data by User ID : `DELETE /users/:id`

### Transaction related

Each endpoint manipulates or displays information related to the Transaction whose Token is provided with the request:

- Get transaction data by ID : `GET /transactions/:id`
- Get transaction data by UserID : `GET /transactions/:id`
- Create transaction by user ID : `POST /transactions/:id`
- Update transaction by user ID : `PUT /transactions/:id`
- Delete transaction data by user ID : `DELETE /transactions/:id`

### Detail Cart related

Each endpoint manipulates or displays information related to the Detail Cart whose Token is provided with the request:

- Get detail cart data by cart ID : `GET /detail_cart/:id`
- Get detail cart data by cart ID & user ID : `GET /detail_cart/:id`
- Create detail cart: `POST /detail_cart/:id`
- Update detail cart by cart ID & user ID : `PUT /detail_cart/:id`
- Delete detail cart data by cart ID & user ID : `DELETE /detail_cart/:id`

## Endpoints that require Check Role isAdmin
The endpoint below requires checking that the currently logged in user role is admin

### Product related

Each endpoint manipulates or displays information related to the Product whose Token and the role is admin that provided with the request:

- Get products data : `GET pproduct`
- Get products data by ID : `GET /produts/:id`
- Get product stock update history data by ID : `GET /produts/stocks/:id`
- Create products : `POST /products`
- Create product stock update history : `POST /product/stocks/:id`
- Update products : `PUT /products/:id`
- Delete products by ID : `DELETE /projects/:id`

### Category related

Each endpoint manipulates or displays information related to the Product Category whose Token is provided with the request:

- Get categories data: `GET /categorys/:id`
- Get category data by ID : `GET /categorys/:id`
- Create category : `POST /categorys/:id`
- Update category : `PUT /categorys/:id`
- Delete category data : `DELETE /categorys/:id`