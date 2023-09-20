# InvisQuery REST APIs Documentation

This document provides documentation for the REST APIs of the GoLang project located in the `api` package. 
These APIs are designed to interact with various aspects of the application, including user management, question management, and answer management.
## Table of Contents

1. [User Management](#user-management)
    - [Create Anonymous User](#create-anonymous-user)
    - [Create User](#create-user)
    - [Login User](#login-user)
    - [Social Login](#social-login)
    - [Get User](#get-user)
    - [Delete User](#delete-user)
    - [Request Password Reset](#request-password-reset)
    - [Reset Password](#reset-password)
    - [Logout User](#logout-user)

2. [Question Management](#question-management)
    - [Create Question](#create-question)
    - [Update Question](#update-question)
    - [Delete Question](#delete-question)
    - [Get Question by ID](#get-question-by-id)
    - [List Questions](#list-questions)

3. [Answer Management](#answer-management)
    - [Create Answer](#create-answer)
    - [Update Answer](#update-answer)
    - [Delete Answer](#delete-answer)
    - [Get Answer by ID](#get-answer-by-id)
    - [List Answers](#list-answers)
# API Information
 - **Base url**:[https://askme-backend-ei6l.onrender.com](https://askme-backend-ei6l.onrender.com)
 - **API Route**: `BaseUrl`+`/api`
 - **API Version**: 'BaseURL'+`/api'+`/v1'
 ## API Documentation
 [https://documenter.getpostman.com/view/16438737/2s9Y5R2mmi](https://documenter.getpostman.com/view/16438737/2s9Y5R2mmi)

## Authentication
Using Jwt token based authentication
```
  Authorization:bearer "token"
```
if request required token than `Authentication Required : true` in API documentation.

## User Management

### Create Anonymous User
**Endpoint**: `POST /api/v1/create-ano-user`

**Description**: Create an anonymous user.


**Request Body:**:
```
    {
    "fcm_token": "string" (optional)
    }
```

### Create User

**Endpoint** : `POST /api/v1/create-user`

**Description**: Create a new user.

**Request Body**:

```

{
  "email": "user@example.com",
  "password": "password",
  "fcm_token": "string"(optional)
}
```

### Login User
**Endpoint** : `POST /api/v1/login-user`

**Description**: Login with email and password.

**Request Body**:


```

{
  "email": "user@example.com",
  "password": "password",
  "fcm_token": "string"(optional)
}
```

### Social Login
**Endpoint** : `POST /api/v1/social-login`

**Description**: Social Login.

**Request Body**:

```
{
  "email": "user@example.com",
  "private_profile_image": "string",
  "provider": "string",
  "fcm_token": "string"(optional)
}
```

### Get User

**Endpoint**: `GET /api/v1/get-user`
**Description**: Get user information.
**Authentication Required**: YES

### Delete User

**Endpoint**: `GET /api/v1/delete-user`
**Description**: Delete user information.
**Authentication Required**: YES

### Request Password Reset

**Endpoint**: `POST /api/v1/request-password-reset`
**Description**: Send a password reset request.
**Request Body**: 
```
{
  "email": "user@example.com"
}

```

### Reset Password
**Endpoint**: `PATCH /resetpassword`
**Description**:Verify and reset the password.
**Request Body**: 
```
{
  "password": "new_password"
}

```

### Logout User
**Endpoint**: `POST /api/v1/logout`
**Description**: Log out the user.


## Question Management

### Create Question

**Endpoint**: `POST /api/v1/create-question`

**Description**: Create a new Question

**Authentication Required**: YES

**Request Body** : 
```
{
  "question": "Question content"
}

```
### Update Question

**Endpoint**: `POST /api/v1/update-question`

**Description**: Update a question by ID.

**Authentication Required**: YES

**Request Body** : 
```
{
  "id": 1,
  "question": "Updated question content"
}

```
### Delete Question
**Endpoint**: `GET /api/v1/delete-question/:id`
**Description**: Delete a question by id.

### Get Question by ID

**Endpoint**: `GET /api/v1/question/:id`

**Description**: Get a question by ID.


### List Questions

**Endpoint**: `GET /api/v1/questions`
**Description**: List question

**Query Parameters**
  - `page_id` : Page Number (minimum 1) **Required**
  - `page_size`: Page Size **Required**
  * `search` : Query search


## Answer Management

### Create Answer

**Endpoint**: `POST /api/v1/create-answer`

**Description**: Create a new Answer

**Authentication Required**: YES

**Request Body** : 

```
{
  "content": "Answer content",
  "question_id": 1
}

```

### Update Answer

**Endpoint**: `POST /api/v1/update-answer`

**Description**: Update a answer by ID.

**Authentication Required**: YES

**Request Body** : 
```
{
  "id": 1,
  "content": "Updated answer content"
}

```
### Delete Answer
**Endpoint**: `GET /api/v1/delete-answer/:id`
**Description**: Delete a answer by id.

### Get Answer by ID

**Endpoint**: `GET /api/v1/answer/:id`

**Description**: Get a answer by ID.


### List Answers

**Endpoint**: `GET /api/v1/answers`
**Description**: List answers to the specific question.

**Query Parameters**
  - `page_id` : Page Number (minimum 1) **Required**
  - `page_size`: Page Size **Required**
  - `search` : Query search
  - `user_id` : User ID 
  -  `question_id` : Question ID  **Required**

















    