package longrunning

import "github.com/mojo-lang/core/go/pkg/mojo/core"

func (x *Operation) GetError() *core.Error {
    return x.GetMojoCoreError()
}

func (x *Operation) GetValue() *core.Any {
    return x.GetMojoCoreAny()
}

func (x *Operation) SetError(err *core.Error) *Operation {
    x.Result = &Operation_MojoCoreError{MojoCoreError: err}
    return x
}

func (x *Operation) SetValue(any *core.Any) *Operation {
    x.Result = &Operation_MojoCoreAny{MojoCoreAny: any}
    return x
}
