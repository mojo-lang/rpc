
# Manages long-running operations with an API service.
When an API method normally takes long time to complete, it can be designedto return [Operation][google.longrunning.Operation] to the client, and the client can use thisinterface to receive the real response asynchronously by polling theoperation resource, or pass the operation resource to another API (such asGoogle Cloud Pub/Sub API) to receive the response.  Any API service thatreturns long-running operations should implement the `Operations` interfaceso developers can have a consistent client experience.
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## Lists operations that match the specified filter in the request. If theserver doesn't support this method, it returns `UNIMPLEMENTED`.
NOTE: the `name` binding allows API services to override the bindingto use different resource name schemes, such as `users/*/operations`. Tooverride the binding, API services can add a binding such as`"/v1/{name=users/*}/operations"` to their service configuration.For backwards compatibility, the default name includes the operationscollection id, however overriding users must ensure the name bindingis the parent resource, without the operations collection id.

### 请求路径
```http
GET /operation/v1/{{service}}/operations
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `parent` | `string` |  | 否 |  | The name of the operation's parent resource. |
| `filter` | `string` |  | 否 |  | The standard list filter. |
| `page_size` | `integer` | `int32` | 否 |  |  |
| `page_token` | `string` |  | 否 |  |  |
| `skip` | `integer` | `int32` | 否 |  |  |


### 返回值

#### 返回对象
| 字段 | 类型 | 说明 |
|---|---|---|
|  | `Array<mojo.rpc.longrunning.Operation>` |


#### `mojo.rpc.longrunning.Operation`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | The server-assigned name, which is only unique within the same service thatoriginally returns it. If you use the default HTTP mapping, the`name` should be a resource name ending with `operations/{unique_id}`. |
| `metadata` | `mojo.core.Any` |  | 否 |  | Service-specific metadata associated with the operation.  It typicallycontains progress information and common metadata such as create time.Some services might not provide such metadata.  Any method that returns along-running operation should document the metadata type, if any. |
| `done` | `boolean` |  | 否 |  | If the value is `false`, it means the operation is still in progress.If `true`, the operation is completed, and either `error` or `response` isavailable. |
| `result` | `Union<mojo.core.Error,mojo.core.Any>` |  | 否 |  | The operation result, which can be either an `error` or a valid `response`.If `done` == `false`, neither `error` nor `response` is set.If `done` == `true`, exactly one of `error` or `response` is set. |


#### `mojo.core.Any`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `type` | `string` |  | 否 |  |
| `val` | `string` | `base64` | 否 |  | Must be valid serialized data of the above specified type. |


#### 
| 类型 | 说明 |
|---|---|
| `mojo.core.Error` |  |
| `mojo.core.Any` |  |


#### `mojo.core.Error`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `code` | `mojo.core.ErrorCode` |  | 否 |  | The error code |
| `message` | `string` |  | 否 |  | A developer-facing error message, which should be in English. Anyuser-facing error message should be localized and sent in the[google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. |
| `details` | `Array<mojo.core.Any>` |  | 否 |  | A list of messages that carry the error details.  There is a common set ofmessage types for APIs to use. |


#### `mojo.core.ErrorCode`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `val` | `integer` | `int32` | 是 |  |  |
| `name` | `string` |  | 否 |  | The name of the error code. This is a constant value that identifies theerror code. Error code name are unique within a particulardomain of errors. This should be at most 63 characters and match/[A-Z0-9_]+/. |
| `domain` | `string` |  | 否 |  | system, runtime, ... |
| `description` | `string` |  | 否 |  | a detail description for the code |
| `document` | `string` | `url` | 否 |  | the api document url for this error code |
| `httpStatusCode` | `integer` | `int32` | 否 |  | the http status code, which is error code mapping to |


## Gets the latest state of a long-running operation.  Clients can use thismethod to poll the operation result at intervals as recommended by the APIservice.

### 请求路径
```http
GET /operation/v1/{{service}}/operations/{name}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `name` | `string` |  | The name of the operation resource. |


### 返回值

#### 返回对象
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | The server-assigned name, which is only unique within the same service thatoriginally returns it. If you use the default HTTP mapping, the`name` should be a resource name ending with `operations/{unique_id}`. |
| `metadata` | `mojo.core.Any` |  | 否 |  | Service-specific metadata associated with the operation.  It typicallycontains progress information and common metadata such as create time.Some services might not provide such metadata.  Any method that returns along-running operation should document the metadata type, if any. |
| `done` | `boolean` |  | 否 |  | If the value is `false`, it means the operation is still in progress.If `true`, the operation is completed, and either `error` or `response` isavailable. |
| `result` | `Union<mojo.core.Error,mojo.core.Any>` |  | 否 |  | The operation result, which can be either an `error` or a valid `response`.If `done` == `false`, neither `error` nor `response` is set.If `done` == `true`, exactly one of `error` or `response` is set. |


#### `mojo.core.Any`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `type` | `string` |  | 否 |  |
| `val` | `string` | `base64` | 否 |  | Must be valid serialized data of the above specified type. |


#### 
| 类型 | 说明 |
|---|---|
| `mojo.core.Error` |  |
| `mojo.core.Any` |  |


#### `mojo.core.Error`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `code` | `mojo.core.ErrorCode` |  | 否 |  | The error code |
| `message` | `string` |  | 否 |  | A developer-facing error message, which should be in English. Anyuser-facing error message should be localized and sent in the[google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. |
| `details` | `Array<mojo.core.Any>` |  | 否 |  | A list of messages that carry the error details.  There is a common set ofmessage types for APIs to use. |


#### `mojo.core.ErrorCode`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `val` | `integer` | `int32` | 是 |  |  |
| `name` | `string` |  | 否 |  | The name of the error code. This is a constant value that identifies theerror code. Error code name are unique within a particulardomain of errors. This should be at most 63 characters and match/[A-Z0-9_]+/. |
| `domain` | `string` |  | 否 |  | system, runtime, ... |
| `description` | `string` |  | 否 |  | a detail description for the code |
| `document` | `string` | `url` | 否 |  | the api document url for this error code |
| `httpStatusCode` | `integer` | `int32` | 否 |  | the http status code, which is error code mapping to |


## Deletes a long-running operation. This method indicates that the client isno longer interested in the operation result. It does not cancel theoperation. If the server doesn't support this method, it returns`google.rpc.Code.UNIMPLEMENTED`.

### 请求路径
```http
DELETE /operation/v1/{{service}}/operations/{name}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `name` | `string` |  | The name of the operation resource to be deleted. |


### 返回值

#### 返回对象
对象为空

## Starts asynchronous cancellation on a long-running operation.  The servermakes a best effort to cancel the operation, but success is notguaranteed.  If the server doesn't support this method, it returns`google.rpc.Code.UNIMPLEMENTED`.  Clients can use[Operations.GetOperation][google.longrunning.Operations.GetOperation] orother methods to check whether the cancellation succeeded or whether theoperation completed despite cancellation. On successful cancellation,the operation is not deleted; instead, it becomes an operation withan [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,corresponding to `Code.CANCELLED`.

### 请求路径
```http
POST /operation/v1/{{service}}/operations/{name}:cancel
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `name` | `string` |  | The name of the operation resource to be cancelled. |


### 返回值

#### 返回对象
对象为空

## Waits until the specified long-running operation is done or reaches at mosta specified timeout, returning the latest state.  If the operation isalready done, the latest state is immediately returned.  If the timeoutspecified is greater than the default HTTP/RPC timeout, the HTTP/RPCtimeout is used.  If the server does not support this method, it returns`google.rpc.Code.UNIMPLEMENTED`.Note that this method is on a best-effort basis.  It may return the lateststate before the specified timeout (including immediately), meaning even animmediate response is no guarantee that the operation is done.

### 请求路径
```http
POST /operation/v1/{{service}}/operations/{name}:wait
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `name` | `string` |  | The name of the operation resource to wait on. |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `timeout` | `string` | `Duration` | 否 |  | The maximum duration to wait before timing out. If left blank, the wait
will be at most the time permitted by the underlying HTTP/RPC protocol.
If RPC context deadline is also specified, the shorter one will be used. |


### 返回值

#### 返回对象
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | The server-assigned name, which is only unique within the same service thatoriginally returns it. If you use the default HTTP mapping, the`name` should be a resource name ending with `operations/{unique_id}`. |
| `metadata` | `mojo.core.Any` |  | 否 |  | Service-specific metadata associated with the operation.  It typicallycontains progress information and common metadata such as create time.Some services might not provide such metadata.  Any method that returns along-running operation should document the metadata type, if any. |
| `done` | `boolean` |  | 否 |  | If the value is `false`, it means the operation is still in progress.If `true`, the operation is completed, and either `error` or `response` isavailable. |
| `result` | `Union<mojo.core.Error,mojo.core.Any>` |  | 否 |  | The operation result, which can be either an `error` or a valid `response`.If `done` == `false`, neither `error` nor `response` is set.If `done` == `true`, exactly one of `error` or `response` is set. |


#### `mojo.core.Any`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `type` | `string` |  | 否 |  |
| `val` | `string` | `base64` | 否 |  | Must be valid serialized data of the above specified type. |


#### 
| 类型 | 说明 |
|---|---|
| `mojo.core.Error` |  |
| `mojo.core.Any` |  |


#### `mojo.core.Error`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `code` | `mojo.core.ErrorCode` |  | 否 |  | The error code |
| `message` | `string` |  | 否 |  | A developer-facing error message, which should be in English. Anyuser-facing error message should be localized and sent in the[google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. |
| `details` | `Array<mojo.core.Any>` |  | 否 |  | A list of messages that carry the error details.  There is a common set ofmessage types for APIs to use. |


#### `mojo.core.ErrorCode`
| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `val` | `integer` | `int32` | 是 |  |  |
| `name` | `string` |  | 否 |  | The name of the error code. This is a constant value that identifies theerror code. Error code name are unique within a particulardomain of errors. This should be at most 63 characters and match/[A-Z0-9_]+/. |
| `domain` | `string` |  | 否 |  | system, runtime, ... |
| `description` | `string` |  | 否 |  | a detail description for the code |
| `document` | `string` | `url` | 否 |  | the api document url for this error code |
| `httpStatusCode` | `integer` | `int32` | 否 |  | the http status code, which is error code mapping to |
