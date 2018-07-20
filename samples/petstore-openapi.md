# Swagger Petstore


# APIs
| Method | Path | Summary |
|--|--|--|
| GET | <a href='#listPets'>/pets</a> | List all pets |
| POST | <a href='#createPets'>/pets</a> | Create a pet |
| GET | <a href='#showPetById'>/pets/{petId}</a> | Info for a specific pet |
# API Details

## GET /pets
<a name='listPets'>List all pets</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| limit | query |  | How many items to return at one time (max 100) |

### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| 200 | application/json | <a href='#/components/schemas/Pets'>Pets<a> | A paged array of pets |
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |

## POST /pets
<a name='createPets'>Create a pet</a>



### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| 201 |  |  | Null response |
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |

## GET /pets/{petId}
<a name='showPetById'>Info for a specific pet</a>



### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| petId | path |  | The id of the pet to retrieve |

### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| 200 | application/json | <a href='#/components/schemas/Pets'>Pets<a> | Expected response to a valid request |
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |
# Components
# Schema Objects
## Pets
<a name='/components/schemas/Pets'>type: array</a>

```
   Pets array 
```
## Error
<a name='/components/schemas/Error'>type: object</a>

```
   Error object 
     code integer 
     message string 
```
## Pet
<a name='/components/schemas/Pet'>type: object</a>

```
   Pet object 
     id integer 
     name string 
     tag string 
```
