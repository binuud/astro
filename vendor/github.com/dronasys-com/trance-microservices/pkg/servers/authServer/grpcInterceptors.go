package authServer

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type StreamContextWrapper interface {
	grpc.ServerStream
	SetContext(context.Context)
}

type wrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrapper) Context() context.Context {
	return w.ctx
}

func (w *wrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}

func newStreamContextWrapper(inner grpc.ServerStream) StreamContextWrapper {
	ctx := inner.Context()
	return &wrapper{
		inner,
		ctx,
	}
}

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logrus.Errorf("metadata is not provided")
		return nil, status.Errorf(codes.Unauthenticated, "Access token absent: ") // force relogin on code Unauthenticated
	}

	values := md["authorization"]
	if len(values) == 0 {
		logrus.Errorf("authorization token is not provided")
		return nil, status.Errorf(codes.Unauthenticated, "Access token is invalid:") // force relogin on code Unauthenticated
	}
	logrus.Trace("Token Received ", values)

	// now parse the token
	jwtsToken := values[0]
	jwt := NewTranceJWT(TRANCE_TOKEN_SECRET, 24*time.Hour)
	claims, err := jwt.Verify(jwtsToken)
	if err != nil {
		logrus.Errorf("Access token Expired: (%s) err(%v)", jwtsToken, err)
		return nil, status.Errorf(codes.Unauthenticated, "Access token Expired: (%s) (%v)", jwtsToken, err) // force relogin on code Unauthenticated
	}
	logrus.Trace("Token Claim", claims)

	md.Append("name", claims.Name)
	md.Append("uuid", claims.Uuid)
	newCtx := metadata.NewIncomingContext(ctx, md)

	return handler(newCtx, req)
}

func StreamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	ctx := stream.Context()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logrus.Errorf("metadata is not provided")
		return status.Errorf(codes.Unauthenticated, "Access token is invalid: ") // force relogin on code Unauthenticated
	}

	values := md["authorization"]
	if len(values) == 0 {
		logrus.Errorf("authorization token is not provided")
		return status.Errorf(codes.Unauthenticated, "Access token is invalid:") // force relogin on code Unauthenticated
	}
	logrus.Trace("Token Received ", values)

	// now parse the token
	jwtsToken := values[0]
	jwt := NewTranceJWT(TRANCE_TOKEN_SECRET, 24*time.Hour)
	claims, err := jwt.Verify(jwtsToken)
	if err != nil {
		logrus.Errorf("Cannot parse token, token invalid", err)
		return status.Errorf(codes.Unauthenticated, "Access token is invalid: %v", err) // force relogin on code Unauthenticated
	}
	logrus.Info("Token Claim", claims)

	md.Append("name", claims.Name)
	md.Append("uuid", claims.Uuid)

	newCtx := metadata.NewIncomingContext(ctx, md)
	newStream := newStreamContextWrapper(stream)
	newStream.SetContext(newCtx)
	return handler(srv, newStream)
}

func GetUserUuidFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := fmt.Errorf("No User information in Context")
		logrus.Errorf("Getting User uuid from context (%v)", err)
		return "", err
	}

	values := md["uuid"]
	if len(values) == 0 {

		err := fmt.Errorf("No User UUID information in Context")
		logrus.Errorf("Getting User uuid from context (%v)", err)
		return "", err

	}
	logrus.Trace("UUID  Received ", values[0])
	return values[0], nil
}

func GetNameFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := fmt.Errorf("No Information in Context")
		logrus.Errorf("Getting Name from context (%v)", err)
		return "", err
	}

	values := md["name"]
	if len(values) == 0 {

		err := fmt.Errorf("No Name Info in Context")
		logrus.Errorf("%v", err)
		return "", err

	}
	logrus.Trace("UUID  Received ", values[0])
	return values[0], nil
}
