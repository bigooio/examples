# Swagger Example API
# General API documentation

**Warning** this api is not production ready. Use at your own risk.

In order to re-generate the documentation you need to run

`swag init --md .`

## Version: 1.0

### Terms of service
<http://swagger.io/terms/>

**Contact information:**  
API Support  
<http://www.swagger.io/support>  
support@swagger.io  

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### /admin/user/

#### GET
##### Summary

List users from the store

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ok | [ [ [api.User](#apiuser) ] ] |

#### POST
##### Summary

Add a new user to the store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| message | body | User Data | Yes | [api.User](#apiuser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ok | string |
| 400 | We need ID!! | [api.APIError](#apiapierror) |
| 404 | Can not find ID | [api.APIError](#apiapierror) |

#### PUT
##### Summary

Add a new user to the store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| message | body | User Data | Yes | [api.User](#apiuser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ok | string |
| 400 | We need ID!! | [api.APIError](#apiapierror) |
| 404 | Can not find ID | [api.APIError](#apiapierror) |

### /admin/user/{id}

#### GET
##### Summary

Read user from the store

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User Id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [api.User](#apiuser) |
| 400 | We need ID!! | [api.APIError](#apiapierror) |
| 404 | Can not find ID | [api.APIError](#apiapierror) |

### Models

#### api.APIError

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| createdAt | string |  | No |
| errorCode | integer |  | No |
| errorMessage | string |  | No |

#### api.User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | No |
| id | integer |  | No |
| password | string |  | No |
