| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `service` | `string` |  | 否 |  | The API service name. It is a logical identifier for a networked API,such as "pubsub.googleapis.com". The naming syntax depends on theAPI management system being used for handling the request. |
| `operation` | `string` |  | 否 |  | The API operation name. For gRPC requests, it is the fully qualified APImethod name, such as "google.pubsub.v1.Publisher.Publish". For OpenAPIrequests, it is the `operationId`, such as "getPet". |
| `protocol` | `string` |  | 否 |  | The API protocol used for sending the request, such as "http", "https","grpc", or "internal". |
| `version` | `string` |  | 否 |  | The API version associated with the API operation above, such as "v1" or"v1alpha1". |
