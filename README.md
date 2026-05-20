# RESTful APIs
This repo is a personal repo to docuement my RESTful APIs learning journey in Golang.

### What is a RESTful API?
**REST** stands for **RE**presentational **S**tate **T**ransfer. It's an architectural style for designing networked applications.

### Core Principles of REST:
**Client-Server Architecture:** Clear separation between **client** (who requests data) and **server** (who provides data).\
**Stateless:** Each request from client to server **must contain** all the information needed to understand and complete the request.\
**Cacheable:** Responses **should be** explicitly cacheable or not cacheable.\
**Uniform Interface:** Consistent way of interacting with resources.

### Key HTTP Methods in REST:
| HTTP Method | CRUD Operation | Description |
| :--- | :--- | :--- |
| **GET** | Read | Retrieve data |
| **POST** | Create | Create a new resource |
| **PUT** | Update | Replace the entire resource |
| **PATCH** | Update | Partially update a resource |
| **DELETE** | Delete | Remove a resource |

### Resources: The Building Blocks

In REST, everything is a resource. A resource is any object that can be identified by a **URI** (Uniform Resource Identifier)

Example resources:
| Resource | Explanation |
| :--- | :--- |
| **/users** | collection of users |
| **/users/123** | specific user with ID 123 |
| **/posts/456/comments** | comments on post 456 |

### Proper HTTP Status Codes
It is **BEST** practice to use proper HTTP Status Codes when creating RESTful APIs

| Status Code | Meaning |
| :--- | :--- |
|**200**| OK - Successs|
|**201**| Created - Resource Created|
|**400**| Bad Request - Invalid request|
|**401**| Unauthorised - Authentication request|
|**403**| Forbidden - No Permission|
|**404**| Not found - Resource not found|
|**500**| Internal Server Error - Server error|

### Why RESTful APIs Are Popular
**Simple:** Uses standard HTTP methods.\
**Stateless:** No server-side session storage.\
**Cacheable:** Responses can be cached.\
**Language-agnostic:** Works with any programming language.\
**Scalable:** Easy to load balance and scale.