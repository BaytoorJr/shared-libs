package auth

import (
	"crypto/rsa"
	"github.com/BaytoorJr/shared-libs/errors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

var (
	authConn  *grpc.ClientConn
	roleConn  *grpc.ClientConn
	publicKey *rsa.PublicKey
)

// Fetch public key from auth api
func getPublicKey(_ *jwt.Token) (interface{}, error) {
	panic("Implement Me")
	return nil, nil
}

// Auth gRPC conn getter
func getAuthGRPCConn() (*grpc.ClientConn, error) {
	if authConn != nil {
		return authConn, nil
	}

	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	conn, err := gRPCDialer(config.AuthConnURL)
	if err != nil {
		return nil, errors.RPCError.SetDevMessage(err.Error())
	}
	authConn = conn

	return authConn, nil
}

// Fetch role actions from role-api
func getRoleActions(roleID int) (interface{}, error) {
	panic("Implement me")
	return nil, nil
}

// Role gRPC conn getter
func getRoleGRPCConn() (*grpc.ClientConn, error) {
	if roleConn != nil {
		return roleConn, nil
	}

	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	conn, err := gRPCDialer(config.RoleConnURL)
	if err != nil {
		return nil, errors.RPCError.SetDevMessage(err.Error())
	}
	roleConn = conn

	return roleConn, nil
}

// gRPC dialer
func gRPCDialer(conn string) (*grpc.ClientConn, error) {
	option := grpc.WithInsecure()

	keepAliveOption := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                60 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	})

	return grpc.Dial(conn, option, keepAliveOption)
}
