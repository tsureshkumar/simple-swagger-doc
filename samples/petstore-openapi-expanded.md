# Swagger Petstore
A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification

# APIs
| Method | Path | Summary |
|--|--|--|
| GET | <a href='#findPets'>/pets</a> |  |
| POST | <a href='#addPet'>/pets</a> |  |
| GET | <a href='#find pet by id'>/pets/{id}</a> |  |
# API Details

## GET /pets
<a name='findPets'></a>

Returns all pets from the system that the user has access to
Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet. Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.

Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra, condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium, pulvinar elit eu, euismod sapien.


### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| tags | query |  | tags to filter by |
| limit | query |  | maximum number of results to return |

### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |
| 200 | application/json |  | pet response |

## POST /pets
<a name='addPet'></a>

Creates a new pet in the store.  Duplicates are allowed
### Request Body
Pet to add to the store

| Content Type | Data |
|--|--|
| application/json | <a href='#/components/schemas/NewPet'>NewPet<a> |

### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| 200 | application/json | <a href='#/components/schemas/Pet'>Pet<a> | pet response |
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |

## GET /pets/{id}
<a name='find pet by id'></a>

Returns a user based on a single ID, if the user does not have access to the pet

### Request Parameters

| Name | Location | Type | Description |
|--|---|---|--|
| id | path |  | ID of pet to fetch |

### Responses

| Code | Content Type | Data | Description |
|--|---|---|--|
| default | application/json | <a href='#/components/schemas/Error'>Error<a> | unexpected error |
| 200 | application/json | <a href='#/components/schemas/Pet'>Pet<a> | pet response |
# Components
# Schema Objects
## Pet
<a name='/components/schemas/Pet'>type: object</a>

```
     NewPet  
       name string 
       tag string 
     Pet  
       id integer 
```
## NewPet
<a name='/components/schemas/NewPet'>type: object</a>

```
   NewPet object 
     name string 
     tag string 
```
## Error
<a name='/components/schemas/Error'>type: object</a>

```
   Error object 
     code integer 
     message string 
```
