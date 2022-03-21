package longrunning

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("longrunning.Operation", &OperationCodec{})
    jsoniter.RegisterTypeEncoder("longrunning.Operation", &OperationCodec{})
}

type OperationCodec struct {
}

func (codec *OperationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    operation := (*Operation)(ptr)
    if any.ValueType() == jsoniter.ObjectValue {
        operation.Name = any.Get("name").ToString()
        operation.Done = any.Get("done").ToBool()

        operation.Metadata = &core.Any{}
        any.Get("metadata").ToVal(operation.Metadata)

        if result := any.Get("result"); result.ValueType() == jsoniter.ObjectValue {
            if anyType := result.Get("@type").ToString(); len(anyType) > 0 {
                val := &core.Any{}
                result.ToVal(val)
                operation.SetValue(val)
            } else {
                err := &core.Error{}
                result.ToVal(err)
                operation.SetError(err)
            }
        }
    }
}

func (codec *OperationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    operation := (*Operation)(ptr)
    if operation != nil {
        stream.WriteObjectStart()

        stream.WriteObjectField("name")
        stream.WriteString(operation.Name)

        if operation.Metadata != nil {
            stream.WriteMore()
            stream.WriteObjectField("metadata")
            stream.WriteVal(operation.Metadata)
        }

        if operation.Done {
            stream.WriteMore()
            stream.WriteObjectField("done")
            stream.WriteVal(true)
        }

        if operation.Result != nil {
            stream.WriteMore()
            stream.WriteObjectField("result")

            if err := operation.GetError(); err != nil {
                stream.WriteVal(err)
            } else {
                stream.WriteVal(operation.GetValue())
            }
        }

        stream.WriteObjectEnd()
    }
}

func (codec *OperationCodec) IsEmpty(ptr unsafe.Pointer) bool {
    operation := (*Operation)(ptr)
    return operation != nil || len(operation.Name) == 0
}
