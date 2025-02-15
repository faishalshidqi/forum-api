# Forum API
This is a Forum API

## Version: 1.0

### Security

| bearerauth | *Bearer Auth* |
|------------|---------------|
| In         | header        |
| Name       | Authorization |

---
### /authentications

#### DELETE
##### Summary

Sign Out

##### Description

Signing User Out. Requires refresh token

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| refreshToken | body | refresh token possessed by the user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.SuccessResponse](#domainssuccessresponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

#### POST
##### Summary

Login with Username & Password

##### Description

authenticate user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| username | body | username address of the user | Yes | string |
| password | body | password of the user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.LoginResponse](#domainsloginresponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domainserrorresponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

#### PUT
##### Summary

Refresh Authentication

##### Description

Generating new access token using a refresh token. Only valid refresh token will generate new

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| refreshToken | body | refresh token possessed by the user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.RefreshResponse](#domainsrefreshresponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domainserrorresponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

---
### /threads

#### POST
##### Summary

Create Thread

##### Description

Creating a new thread. Only valid users can create a thread

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer Token | Yes | string |
| title | body | title of the thread | Yes | string |
| body | body | body of the thread | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.AddThreadResponse](#domainsaddthreadresponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domainserrorresponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

---
### /threads/{thread_id}/comments

#### POST
##### Summary

Soft Delete Comment

##### Description

Soft Delete a  comment. Only valid users can delete their own comment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer Token | Yes | string |
| thread_id | path | Thread ID | Yes | string |
| comment_id | path | Comment ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.AddCommentResponse](#domainsaddcommentresponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domainserrorresponse) |
| 403 | Forbidden | [domains.ErrorResponse](#domainserrorresponse) |
| 404 | Not Found | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

---
### /users

#### POST
##### Summary

Register A User

##### Description

New user must have a unique email address

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| email | body | email address of the new user, must be unique | Yes | string |
| password | body | password of the new user | Yes | string |
| name | body | name of the new user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.SignupResponse](#domainssignupresponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domainserrorresponse) |
| 409 | Conflict | [domains.ErrorResponse](#domainserrorresponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domainserrorresponse) |

---
### Models

#### domains.AddCommentResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.AddCommentResponseData](#domainsaddcommentresponsedata) |  | No |
| message | string |  | No |
| status | string |  | No |

#### domains.AddCommentResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | string |  | No |
| id | string |  | No |
| owner | string |  | No |

#### domains.AddThreadResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.AddThreadResponseData](#domainsaddthreadresponsedata) |  | No |
| message | string |  | No |
| status | string |  | No |

#### domains.AddThreadResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| owner | string |  | No |
| title | string |  | No |

#### domains.ErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| status | string |  | No |

#### domains.LoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.LoginResponseData](#domainsloginresponsedata) |  | No |
| status | string |  | No |

#### domains.LoginResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | string |  | No |
| refreshToken | string |  | No |

#### domains.RefreshResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.RefreshResponseData](#domainsrefreshresponsedata) |  | No |
| status | string |  | No |

#### domains.RefreshResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | string |  | No |

#### domains.SignupResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.SignupResponseData](#domainssignupresponsedata) |  | No |
| message | string |  | No |
| status | string |  | No |

#### domains.SignupResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| fullname | string |  | No |
| id | string |  | No |
| username | string |  | No |

#### domains.SuccessResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| status | string |  | No |
