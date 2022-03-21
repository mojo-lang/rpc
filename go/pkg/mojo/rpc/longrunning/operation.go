package longrunning

import "github.com/mojo-lang/core/go/pkg/mojo/core"

func (m *Operation) GetError() *core.Error {
    return m.GetMojoCoreError()
}

func (m *Operation) GetValue() *core.Any {
    return m.GetMojoCoreAny()
}

func (m *Operation) SetError(err *core.Error) *Operation {
    m.Result = &Operation_MojoCoreError{MojoCoreError: err}
    return m
}

func (m *Operation) SetValue(any *core.Any) *Operation {
    m.Result = &Operation_MojoCoreAny{MojoCoreAny: any}
    return m
}
