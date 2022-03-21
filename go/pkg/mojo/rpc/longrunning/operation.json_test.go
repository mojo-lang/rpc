package longrunning

import (
    "testing"

    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    _ "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/stretchr/testify/assert"
)

const operationMetadata = `{
    "name": "uuid",
    "metadata": {
        "@type": "mojo.core.Checksum",
        "algorithm": "sha256",
        "value": "value"
    }
}`

const operationError = `{
    "name": "uuid",
    "done": true,
    "metadata": {
        "@type": "mojo.core.Checksum",
        "algorithm": "sha256",
        "value": "value"
    },
    "result": {
        "code": "404"
        "message": "something wrong"
    }
}`

const operationValue = `{
    "name": "uuid",
    "done": true,
    "metadata": {
        "@type": "mojo.core.Checksum",
        "algorithm": "sha256",
        "value": "value"
    },
    "result": {
        "@type": "mojo.core.Checksum",
        "algorithm": "sha256",
        "value": "value"
    }
}`

func TestOperationCodec_Decode(t *testing.T) {
    operation := &Operation{}
    jsoniter.ConfigFastest.Unmarshal([]byte(operationMetadata), operation)

    assert.Equal(t, "uuid", operation.Name)
    assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Decode_Error(t *testing.T) {
    operation := &Operation{}
    jsoniter.ConfigFastest.Unmarshal([]byte(operationError), operation)

    assert.Equal(t, "uuid", operation.Name)
    assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Decode_Value(t *testing.T) {
    operation := &Operation{}
    jsoniter.ConfigFastest.Unmarshal([]byte(operationValue), operation)

    assert.Equal(t, "uuid", operation.Name)
    assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Encode(t *testing.T) {
    operation := &Operation{
        Name:     "uuid",
        Metadata: core.NewAny(&core.Checksum{Algorithm: core.Checksum_ALGORITHM_SHA256, Value: "value"}),
        Done:     false,
    }

    str, err := jsoniter.ConfigFastest.MarshalToString(operation)
    assert.NoError(t, err)

    var f interface{}
    jsoniter.ConfigFastest.UnmarshalFromString(operationMetadata, &f)
    out, _ := jsoniter.ConfigFastest.MarshalToString(f)
    assert.Equal(t, out, str)
}
