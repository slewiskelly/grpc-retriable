// Package retriable provides helper functions for determining whether an error
// containing a gRPC status code is retriable.
//
// See https://github.com/grpc/grpc/blob/master/doc/statuscodes.md for more
// information about gRPC status codes.
package retriable

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Is returns true if the error contains one of the following gRPC status codes:
// - Unknown
// - ResourceExhausted
// - Internal
// - Unavailable
//
// False is returned for all other status codes, or if the error does not
// contain a gRPC status code.
//
// Note: The gRPC documentation states that the Internal status code is
// "reserved for serious errors". However, this status code has been adopted as
// a generic code when no other more specific code is appropriate.
func Is(err error) bool {
	if s, ok := status.FromError(err); ok {
		return m[s.Code()]
	}

	return false
}

// Has is the same as Is but will check errors that may have been wrapped. The
// first error in the chain with a status code will be used to determine if the
// error is retriable.
//
// See Is for more information.
func Has(err error) bool {
	if Is(err) {
		return true
	}

	if u := errors.Unwrap(err); u != nil {
		return Has(u)
	}

	return false
}

var m = map[codes.Code]bool{
	// Non-retriable codes:
	//  - OK
	//  - Canceled
	//  - InvalidArgument
	//  - DeadlineExceeded
	//  - NotFound
	//  - AlreadyExists
	//  - PermissionDenied
	//  - FailedPrecondition
	//  - Aborted
	//  - OutOfRange
	//  - Unimplemented
	//  - DataLoss
	//  - Unauthenticated

	// Retriable codes:
	codes.Unknown:           true,
	codes.ResourceExhausted: true,
	codes.Internal:          true,
	codes.Unavailable:       true,
}
