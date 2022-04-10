package longrunning

import "github.com/mojo-lang/core/go/pkg/mojo/core"

func (x *Operation) GetValue() *core.Any {
    return x.GetResponse()
}

func (x *Operation) SetError(err *core.Error) *Operation {
    if x != nil {
        x.Error = err
    }
    return x
}

func (x *Operation) SetValue(any *core.Any) *Operation {
    if x != nil {
        x.Response = any
    }
    return x
}
