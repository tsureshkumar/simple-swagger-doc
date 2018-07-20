# Swagger Petstore
This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.

# APIs
| Method | Path | Summary |
|--|--|--|
| POST | <a href='#addPet'>/pet</a> | Add a new pet to the store |
| GET | <a href='#findPetsByStatus'>/pet/findByStatus</a> | Finds Pets by status |
| GET | <a href='#findPetsByTags'>/pet/findByTags</a> | Finds Pets by tags |
| GET | <a href='#getPetById'>/pet/{petId}</a> | Find pet by ID |
| POST | <a href='#updatePetWithForm'>/pet/{petId}</a> | Updates a pet in the store with form data |
| POST | <a href='#uploadFile'>/pet/{petId}/uploadImage</a> | uploads an image |
| GET | <a href='#getInventory'>/store/inventory</a> | Returns pet inventories by status |
| POST | <a href='#placeOrder'>/store/order</a> | Place an order for a pet |
| GET | <a href='#getOrderById'>/store/order/{orderId}</a> | Find purchase order by ID |
| POST | <a href='#createUser'>/user</a> | Create user |
| POST | <a href='#createUsersWithArrayInput'>/user/createWithArray</a> | Creates list of users with given input array |
| POST | <a href='#createUsersWithListInput'>/user/createWithList</a> | Creates list of users with given input array |
| GET | <a href='#loginUser'>/user/login</a> | Logs user into the system |
| GET | <a href='#logoutUser'>/user/logout</a> | Logs out current logged in user session |
| GET | <a href='#getUserByName'>/user/{username}</a> | Get user by user name |
# API Details

## POST /pet
<a name='addPet'>Add a new pet to the store</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| body | body | <a href='#/definitions/Pet'>Pet<a> | Pet object that needs to be added to the store |

### Responses

| Code | Data | Description |
|--|---|--|
| 405 |  | Invalid input |

## GET /pet/findByStatus
<a name='findPetsByStatus'>Finds Pets by status</a>

Multiple status values can be provided with comma separated strings

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| status | query | array | Status values that need to be considered for filter |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 |  | successful operation |
| 400 |  | Invalid status value |

## GET /pet/findByTags
<a name='findPetsByTags'>Finds Pets by tags</a>

Muliple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| tags | query | array | Tags to filter by |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 |  | successful operation |
| 400 |  | Invalid tag value |

## GET /pet/{petId}
<a name='getPetById'>Find pet by ID</a>

Returns a single pet

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| petId | path | integer | ID of pet to return |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 | <a href='#/definitions/Pet'>Pet<a> | successful operation |
| 400 |  | Invalid ID supplied |
| 404 |  | Pet not found |

## POST /pet/{petId}
<a name='updatePetWithForm'>Updates a pet in the store with form data</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| petId | path | integer | ID of pet that needs to be updated |
| name | formData | string | Updated name of the pet |
| status | formData | string | Updated status of the pet |

### Responses

| Code | Data | Description |
|--|---|--|
| 405 |  | Invalid input |

## POST /pet/{petId}/uploadImage
<a name='uploadFile'>uploads an image</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| petId | path | integer | ID of pet to update |
| additionalMetadata | formData | string | Additional data to pass to server |
| file | formData | file | file to upload |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 | <a href='#/definitions/ApiResponse'>ApiResponse<a> | successful operation |

## GET /store/inventory
<a name='getInventory'>Returns pet inventories by status</a>

Returns a map of status codes to quantities

### Responses

| Code | Data | Description |
|--|---|--|
| 200 |  | successful operation |

## POST /store/order
<a name='placeOrder'>Place an order for a pet</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| body | body | <a href='#/definitions/Order'>Order<a> | order placed for purchasing the pet |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 | <a href='#/definitions/Order'>Order<a> | successful operation |
| 400 |  | Invalid Order |

## GET /store/order/{orderId}
<a name='getOrderById'>Find purchase order by ID</a>

For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| orderId | path | integer | ID of pet that needs to be fetched |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 | <a href='#/definitions/Order'>Order<a> | successful operation |
| 400 |  | Invalid ID supplied |
| 404 |  | Order not found |

## POST /user
<a name='createUser'>Create user</a>

This can only be done by the logged in user.

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| body | body | <a href='#/definitions/User'>User<a> | Created user object |

### Responses

| Code | Data | Description |
|--|---|--|
| default |  | successful operation |

## POST /user/createWithArray
<a name='createUsersWithArrayInput'>Creates list of users with given input array</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| body | body |  | List of user object |

### Responses

| Code | Data | Description |
|--|---|--|
| default |  | successful operation |

## POST /user/createWithList
<a name='createUsersWithListInput'>Creates list of users with given input array</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| body | body |  | List of user object |

### Responses

| Code | Data | Description |
|--|---|--|
| default |  | successful operation |

## GET /user/login
<a name='loginUser'>Logs user into the system</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| username | query | string | The user name for login |
| password | query | string | The password for login in clear text |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 |  | successful operation |
| 400 |  | Invalid username/password supplied |

## GET /user/logout
<a name='logoutUser'>Logs out current logged in user session</a>



### Responses

| Code | Data | Description |
|--|---|--|
| default |  | successful operation |

## GET /user/{username}
<a name='getUserByName'>Get user by user name</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| username | path | string | The name that needs to be fetched. Use user1 for testing.  |

### Responses

| Code | Data | Description |
|--|---|--|
| 200 | <a href='#/definitions/User'>User<a> | successful operation |
| 400 |  | Invalid username supplied |
| 404 |  | User not found |
# Definitions
## ApiResponse
<a name='/definitions/ApiResponse'>type: object</a>

```
   ApiResponse object 
     code integer 
     type string 
     message string 
```
## Order
<a name='/definitions/Order'>type: object</a>

```
   Order object 
     id integer 
     petId integer 
     quantity integer 
     shipDate string 
     status string 	// Order Status
       enum:
         placed
         approved
         delivered
     complete boolean 
```
## User
<a name='/definitions/User'>type: object</a>

```
   User object 
     userStatus integer 	// User Status
     id integer 
     username string 
     firstName string 
     lastName string 
     email string 
     password string 
     phone string 
```
## Category
<a name='/definitions/Category'>type: object</a>

```
   Category object 
     id integer 
     name string 
```
## Tag
<a name='/definitions/Tag'>type: object</a>

```
   Tag object 
     id integer 
     name string 
```
## Pet
<a name='/definitions/Pet'>type: object</a>

```
   Pet object 
     id integer 
     category  
     name string 
     photoUrls array 
     tags array 
     status string 	// pet status in the store
       enum:
         available
         pending
         sold
```
