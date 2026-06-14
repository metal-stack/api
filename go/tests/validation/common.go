package validation

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
)

type (
	prototest struct {
		name             string
		msg              proto.Message
		wantErr          bool
		wantErrorMessage string
	}

	prototests []prototest
)

func createString(n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = 'a'
	}
	return string(s)
}

func createRepeatedString(n int) []string {
	arr := make([]string, n)
	for i := range arr {
		arr[i] = "opt" + string(rune('a'+(i%26)))
	}
	return arr
}

func createDiskDevices(n int) []string {
	devs := make([]string, n)
	for i := range devs {
		devs[i] = "/dev/sd" + string('a'+byte(i%26)) + string('0'+byte(i/26))
	}
	return devs
}

func createStringSlice(n int) []string {
	arr := make([]string, n)
	for i := range arr {
		arr[i] = "tag" + string(rune('a'+i%26))
	}
	return arr
}

func validateProtos(t *testing.T, tests prototests) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := protovalidate.Validate(tt.msg)
			if err != nil && !tt.wantErr {
				t.Errorf("validate error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErr {
				if diff := cmp.Diff(err.Error(), tt.wantErrorMessage); diff != "" {
					t.Errorf("validate error = %v, diff %v", err, diff)
				}
			}
		})
	}
}
