package errors

import (	
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Helper functions for common errors
func ErrProductNotFound(name string) error {
	return status.Errorf(codes.NotFound, "Product '%s' not found", name)
}

func ErrInvalidRequest(msg string) error {
	return status.Errorf(codes.InvalidArgument, msg)
}

func ErrDatabaseFailure(err error) error {
	return status.Errorf(codes.Internal, "Database error: %v", err)
}

func ErrClientHandshake(reason string, err error) error {
	log.Printf("handshake error: %s %v", reason, err)
	return status.Errorf(codes.Internal, "%s: %v", reason, err)
}

func ErrStoreNameEnv(reason string) error {
	return status.Error(codes.NotFound, reason)
}