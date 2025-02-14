package enum

import (
	"fmt"

	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"
	"google.golang.org/protobuf/proto"
)

func GetFormatStringValue(f apiv2.Format) (string, error) {
	value := f.Descriptor().Values().Get(int(*f.Enum()))

	// Retrieve custom options
	opts := value.Options()

	// Check if the option is present
	if opts != nil && proto.HasExtension(opts, apiv2.E_EnumStringValue) {
		stringValue := proto.GetExtension(opts, apiv2.E_EnumStringValue).(string)
		return stringValue, nil
	}

	return "", fmt.Errorf("unable to fetch stringvalue from %s", f.String())
}

func GetGPTTypeStringValue(f apiv2.GPTType) (string, error) {
	value := f.Descriptor().Values().Get(int(*f.Enum()))

	// Retrieve custom options
	opts := value.Options()

	// Check if the option is present
	if opts != nil && proto.HasExtension(opts, apiv2.E_EnumStringValue) {
		stringValue := proto.GetExtension(opts, apiv2.E_EnumStringValue).(string)
		return stringValue, nil
	}

	return "", fmt.Errorf("unable to fetch stringvalue from %s", f.String())
}

func GetLVMTypeStringValue(f apiv2.LVMType) (string, error) {
	value := f.Descriptor().Values().Get(int(*f.Enum()))

	// Retrieve custom options
	opts := value.Options()

	// Check if the option is present
	if opts != nil && proto.HasExtension(opts, apiv2.E_EnumStringValue) {
		stringValue := proto.GetExtension(opts, apiv2.E_EnumStringValue).(string)
		return stringValue, nil
	}

	return "", fmt.Errorf("unable to fetch stringvalue from %s", f.String())
}
