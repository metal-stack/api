package id

import (
	"fmt"
	"reflect"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Get[M proto.Message](m M) (string, error) {
	var primaryKey string

	if reflect.ValueOf(m).Kind() == reflect.Pointer {
		if reflect.ValueOf(m).IsNil() {
			return primaryKey, fmt.Errorf("given message is a nil pointer")
		}
	}

	var (
		value protoreflect.Value
		field protoreflect.FieldDescriptor
	)
	pm := m.ProtoReflect()
	field = pm.Descriptor().Fields().ByTextName("id")
	if field == nil {
		pm.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			if fd.Kind() != protoreflect.StringKind {
				return false
			}
			opts := fd.Options()
			if opts != nil && proto.HasExtension(opts, apiv2.E_TreatAsId) {
				// FIXME ensure in api_scopes_test only one tread_as_id annotation per message
				value = pm.Get(fd)
				field = fd
				return false
			}
			return true
		})
	}
	if field == nil {
		return primaryKey, fmt.Errorf("unable to fetch primaryKey from %v", m)
	}
	value = pm.Get(field)
	primaryKey = value.String()
	return primaryKey, nil
}
