| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | The server-assigned name, which is only unique within the same service thatoriginally returns it. If you use the default HTTP mapping, the`name` should be a resource name ending with `operations/{unique_id}`. |
| `metadata` | `mojo.core.Any` |  | 否 |  | Service-specific metadata associated with the operation.  It typicallycontains progress information and common metadata such as create time.Some services might not provide such metadata.  Any method that returns along-running operation should document the metadata type, if any. |
| `done` | `boolean` |  | 否 |  | If the value is `false`, it means the operation is still in progress.If `true`, the operation is completed, and either `error` or `response` isavailable. |
| `result` | `Union<mojo.core.Error,mojo.core.Any>` |  | 否 |  | The operation result, which can be either an `error` or a valid `response`.If `done` == `false`, neither `error` nor `response` is set.If `done` == `true`, exactly one of `error` or `response` is set. |
