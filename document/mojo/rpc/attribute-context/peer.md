| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `ip` | `string` |  | 否 |  | The IP address of the peer. |
| `port` | `integer` | `int64` | 否 |  | The network port of the peer. |
| `labels` | `Map<string, string>` |  | 否 |  | The labels associated with the peer. |
| `principal` | `string` |  | 否 |  | The identity of this peer. Similar to `Request.auth.principal`, butrelative to the peer instead of the request. For example, theidenity associated with a load balancer that forwared the request. |
| `regionCode` | `string` |  | 否 |  | The CLDR country/region code associated with the above IP address.If the IP address is private, the `region_code` should reflect thephysical location where this peer is running. |
