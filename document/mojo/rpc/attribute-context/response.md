| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `code` | `integer` | `int64` | 否 |  | The HTTP response status code, such as `200` and `404`. |
| `size` | `integer` | `int64` | 否 |  | The HTTP response size in bytes. If unknown, it must be -1. |
| `headers` | `Map<string, string>` |  | 否 |  | The HTTP response headers. If multiple headers share the same key, theymust be merged according to HTTP spec. All header keys must belowercased, because HTTP header keys are case-insensitive. |
| `time` | `string` | `DateTime` | 否 |  | The timestamp when the `destination` service sends the last byte ofthe response. |
| `backendLatency` | `string` | `Duration` | 否 |  | The length of time it takes the backend service to fully respond to arequest. Measured from when the destination service starts to send therequest to the backend until when the destination service receives thecomplete response from the backend. |
