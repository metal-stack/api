package enum

import (
	"fmt"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func GetStringValue[E protoreflect.Enum](f E) (string, error) {
	value := f.Descriptor().Values().Get(int(f.Number()))

	// Retrieve custom options
	opts := value.Options()

	// Check if the option is present
	if opts != nil && proto.HasExtension(opts, apiv2.E_EnumStringValue) {
		stringValue := proto.GetExtension(opts, apiv2.E_EnumStringValue).(string)
		return stringValue, nil
	}

	return "", fmt.Errorf("unable to fetch stringvalue from %v", f)
}
