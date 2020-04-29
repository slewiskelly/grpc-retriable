package retriable

import (
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHasWrapped(t *testing.T) {
	tests := []struct {
		err  error
		want bool
	}{{
		status.Error(codes.OK, "ok"),
		false,
	}, {
		status.Error(codes.Canceled, "canceled"),
		false,
	}, {
		status.Error(codes.Unknown, "unknown"),
		true,
	}, {
		status.Error(codes.InvalidArgument, "invalid argument"),
		false,
	}, {
		status.Error(codes.DeadlineExceeded, "deadline exceeded"),
		false,
	}, {
		status.Error(codes.NotFound, "not found"),
		false,
	}, {
		status.Error(codes.AlreadyExists, "already exists"),
		false,
	}, {
		status.Error(codes.PermissionDenied, "permission denied"),
		false,
	}, {
		status.Error(codes.ResourceExhausted, "resource exhausted"),
		true,
	}, {
		status.Error(codes.FailedPrecondition, "failed precondition"),
		false,
	}, {
		status.Error(codes.Aborted, "aborted"),
		false,
	}, {
		status.Error(codes.OutOfRange, "out of range"),
		false,
	}, {
		status.Error(codes.Unimplemented, "unimplemented"),
		false,
	}, {
		status.Error(codes.Internal, "internal"),
		true,
	}, {
		status.Error(codes.Unavailable, "unavailable"),
		true,
	}, {
		status.Error(codes.DataLoss, "data loss"),
		false,
	}, {
		status.Error(codes.Unauthenticated, "unauthenticated"),
		false,
	}, {
		fmt.Errorf("something went wrong: %w", status.Error(codes.Canceled, "cancelled")),
		false,
	}, {
		fmt.Errorf("something went wrong: %w", status.Error(codes.Unavailable, "unavailable")),
		true,
	}, {
		fmt.Errorf("something went wrong: %w", fmt.Errorf("uh oh: %w", status.Error(codes.Unavailable, "unavailable"))),
		true,
	}, {}, {
		errors.New("no status"),
		false,
	}, {
		nil,
		false,
	}}

	for _, test := range tests {
		if got := Has(test.err); got != test.want {
			t.Errorf("Has(%v) = %v; want %v", test.err, got, test.want)
		}
	}
}
