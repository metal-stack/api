package errorutil

import (
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect/v2"
	"google.golang.org/grpc/status"

	"github.com/google/go-cmp/cmp"
)

// Convert compares the error and maps it to an appropriate connect.Error.
// If there is no backend-specific wrapped error given to this function, it will return an internal error.
// If there are multiple errors wrapped, this function only cares about the first found error in the error tree.
func Convert(err error) *connect.Error {
	if err == nil {
		return connect.NewError(connect.CodeUnknown, "")
	}

	// Ipam or other connect errors
	if connectErr, ok := errors.AsType[*connect.Error](err); ok {
		// when the connect error is wrapped deeper a tree, connect.Error() calls the string function on
		// the error and adds things like "internal: ..."
		// so we replace the wrapped error message with the direct message
		cleaned := strings.Replace(err.Error(), connectErr.Error(), connectErr.Message(), 1)
		return connect.NewError(connectErr.Code(), cleaned)
	}

	// for other pure grpc apis
	if _, ok := status.FromError(err); ok {
		// when the grpc error is wrapped deeper a tree, status.FromError calls the string function on
		// the error and adds "rpc error: ..."
		// so we unwrap the grpc status on our own and replace the error message with the direct message
		type grpcstatus interface{ GRPCStatus() *status.Status }
		var (
			iterErr = err
		)

		for {
			st, ok := iterErr.(grpcstatus)
			if ok {
				cleaned := strings.Replace(err.Error(), st.GRPCStatus().String(), st.GRPCStatus().Message(), 1)
				return connect.NewError(connect.Code(st.GRPCStatus().Code()), cleaned)
			}

			iterErr = errors.Unwrap(iterErr)
			if iterErr == nil {
				// just a theoretical case
				return connect.NewError(connect.CodeUnknown, err.Error())
			}
		}
	}

	return connect.NewError(connect.CodeInternal, err.Error()).WithCause(err)
}

// newError creates a connect.Error with the given code and formats the message
// using the supplied format and arguments. It supports the %w verb, wrapping
// the referenced error as the connect.Error's cause.
func newError(code connect.Code, format string, args ...any) *connect.Error {
	err := fmt.Errorf(format, args...)
	return connect.NewError(code, err.Error()).WithCause(err)
}

// newErrorFrom creates a connect.Error with the given code from an underlying error.
func newErrorFrom(code connect.Code, err error) error {
	if err == nil {
		return connect.NewError(code, "")
	}
	return connect.NewError(code, err.Error()).WithCause(err)
}

// NotFound creates a new Not Found error (Connect Code 5) with a formatted error message.
func NotFound(format string, args ...any) error {
	return newError(connect.CodeNotFound, format, args...)
}

// Conflict creates a new Already Exists error (Connect Code 6) with a formatted error message.
func Conflict(format string, args ...any) error {
	return newError(connect.CodeAlreadyExists, format, args...)
}

// Internal creates a new Internal error (Connect Code 13) with a formatted error message.
func Internal(format string, args ...any) error {
	return newError(connect.CodeInternal, format, args...)
}

// Aborted creates a new Aborted error (Connect Code 10) with a formatted error message.
func Aborted(format string, args ...any) error {
	return newError(connect.CodeAborted, format, args...)
}

// InvalidArgument creates a new Invalid Argument error (Connect Code 3) with a formatted error message.
func InvalidArgument(format string, args ...any) error {
	return newError(connect.CodeInvalidArgument, format, args...)
}

// FailedPrecondition creates a new Failed Precondition error (Connect Code 9) with a formatted error message.
func FailedPrecondition(format string, args ...any) error {
	return newError(connect.CodeFailedPrecondition, format, args...)
}

// Unauthenticated creates a new Unauthenticated error (Connect Code 16) with a formatted error message.
func Unauthenticated(format string, args ...any) error {
	return newError(connect.CodeUnauthenticated, format, args...)
}

// ResourceExhausted creates a new Resource Exhausted error (Connect Code 8) with a formatted error message.
func ResourceExhausted(format string, args ...any) error {
	return newError(connect.CodeResourceExhausted, format, args...)
}

// PermissionDenied creates a new Permission Denied error (Connect Code 7) with a formatted error message.
func PermissionDenied(format string, args ...any) error {
	return newError(connect.CodePermissionDenied, format, args...)
}

// Unimplemented creates a new Unimplemented error (Connect Code 12) with a formatted error message.
func Unimplemented(format string, args ...any) error {
	return newError(connect.CodeUnimplemented, format, args...)
}

// NewNotFound creates a new Not Found error (Connect Code 5) from an underlying error.
func NewNotFound(err error) error {
	return newErrorFrom(connect.CodeNotFound, err)
}

// NewConflict creates a new Already Exists error (Connect Code 6) from an underlying error.
func NewConflict(err error) error {
	return newErrorFrom(connect.CodeAlreadyExists, err)
}

// NewInternal creates a new Internal error (Connect Code 13) from an underlying error.
func NewInternal(err error) error {
	return newErrorFrom(connect.CodeInternal, err)
}

// NewAborted creates a new Aborted error (Connect Code 10) from an underlying error.
func NewAborted(err error) error {
	return newErrorFrom(connect.CodeAborted, err)
}

// NewInvalidArgument creates a new Invalid Argument error (Connect Code 3) from an underlying error.
func NewInvalidArgument(err error) error {
	return newErrorFrom(connect.CodeInvalidArgument, err)
}

// NewFailedPrecondition creates a new Failed Precondition error (Connect Code 9) from an underlying error.
func NewFailedPrecondition(err error) error {
	return newErrorFrom(connect.CodeFailedPrecondition, err)
}

// NewUnauthenticated creates a new Unauthenticated error (Connect Code 16) from an underlying error.
func NewUnauthenticated(err error) error {
	return newErrorFrom(connect.CodeUnauthenticated, err)
}

// NewResourceExhausted creates a new Resource Exhausted error (Connect Code 8) from an underlying error.
func NewResourceExhausted(err error) error {
	return newErrorFrom(connect.CodeResourceExhausted, err)
}

// NewPermissionDenied creates a new Permission Denied error (Connect Code 7) from an underlying error.
func NewPermissionDenied(err error) error {
	return newErrorFrom(connect.CodePermissionDenied, err)
}

// NewUnimplemented creates a new Unimplemented error (Connect Code 12) from an underlying error.
func NewUnimplemented(err error) error {
	return newErrorFrom(connect.CodeUnimplemented, err)
}

// IsNotFound reports whether the error is a Not Found error (Connect Code 5).
func IsNotFound(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeNotFound
}

// IsConflict reports whether the error is an Already Exists error (Connect Code 6).
func IsConflict(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeAlreadyExists
}

// IsInternal reports whether the error is an Internal error (Connect Code 13).
func IsInternal(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeInternal
}

// IsAborted reports whether the error is an Aborted error (Connect Code 10).
func IsAborted(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeAborted
}

// IsInvalidArgument reports whether the error is an Invalid Argument error (Connect Code 3).
func IsInvalidArgument(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeInvalidArgument
}

// IsFailedPrecondition reports whether the error is a Failed Precondition error (Connect Code 9).
func IsFailedPrecondition(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeFailedPrecondition
}

// IsUnauthenticated reports whether the error is an Unauthenticated error (Connect Code 16).
func IsUnauthenticated(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeUnauthenticated
}

// IsResourceExhausted reports whether the error is a Resource Exhausted error (Connect Code 8).
func IsResourceExhausted(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeResourceExhausted
}

// IsPermissionDenied reports whether the error is a Permission Denied error (Connect Code 7).
func IsPermissionDenied(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodePermissionDenied
}

// IsUnimplemented reports whether the error is a Unimplemented error (Connect Code 12).
func IsUnimplemented(err error) bool {
	connectErr := Convert(err)
	return connectErr.Code() == connect.CodeUnimplemented
}

// WrapConnectErr wraps the error into a connect.Error with the given code if it is not already a connect.Error.
// If the error is already a *connect.Error, it returns it unchanged.
func WrapConnectErr(c connect.Code, underlying error) *connect.Error {
	if connectErr, ok := errors.AsType[*connect.Error](underlying); ok {
		return connectErr
	}

	return connect.NewError(c, underlying.Error()).WithCause(underlying)
}

// ConnectErrorComparer returns a go-cmp Option for comparing *connect.Error values.
// Two errors are considered equal if they have the same error string and the same Connect code.
func ConnectErrorComparer() cmp.Option {
	return cmp.Comparer(func(x, y *connect.Error) bool {
		if x == nil && y == nil {
			return true
		}
		if x == nil && y != nil {
			return false
		}
		if x != nil && y == nil {
			return false
		}
		if x.Error() != y.Error() {
			return false
		}
		return x.Code() == y.Code()
	})
}

// ErrorStringComparer returns a go-cmp Option for comparing error values.
// Two errors are considered equal if their Error() strings are identical.
func ErrorStringComparer() cmp.Option {
	return cmp.Comparer(func(x, y error) bool {
		if x == nil && y == nil {
			return true
		}
		if x == nil && y != nil {
			return false
		}
		if x != nil && y == nil {
			return false
		}
		return x.Error() == y.Error()
	})
}
