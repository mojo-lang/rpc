| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `origin` | `mojo.rpc.AttributeContext.Peer` |  | 否 |  | The origin of a network activity. In a multi hop network activity,the origin represents the sender of the first hop. For the first hop,the `source` and the `origin` must have the same content. |
| `source` | `mojo.rpc.AttributeContext.Peer` |  | 否 |  | The source of a network activity, such as starting a TCP connection.In a multi hop network activity, the source represents the sender of thelast hop. |
| `destination` | `mojo.rpc.AttributeContext.Peer` |  | 否 |  | The destination of a network activity, such as accepting a TCP connection.In a multi hop network activity, the destination represents the receiver ofthe last hop. |
| `request` | `mojo.rpc.AttributeContext.Request` |  | 否 |  | Represents a network request, such as an HTTP request. |
| `response` | `mojo.rpc.AttributeContext.Response` |  | 否 |  | Represents a network response, such as an HTTP response. |
| `resource` | `mojo.rpc.AttributeContext.Resource` |  | 否 |  | Represents a target resource that is involved with a network activity.If multiple resources are involved with an activity, this must be theprimary one. |
| `api` | `mojo.rpc.AttributeContext.Api` |  | 否 |  | Represents an API operation that is involved to a network activity. |
| `extensions` | `Array<mojo.core.Any>` |  | 否 |  | Supports extensions for advanced use cases, such as logs and metrics. |
