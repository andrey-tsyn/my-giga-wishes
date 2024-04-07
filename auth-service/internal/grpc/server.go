package grpc

import (
	"context"

	authv1 "github.com/andrey-tsyn/my-giga-wishes/api-contracts/gen/go/auth-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authv1.UnimplementedAuthServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Register(context.Context, *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (s *Server) LoginViaEmail(context.Context, *authv1.LoginViaEmailRequest) (*authv1.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginViaEmail not implemented")
}

func (s *Server) LoginViaUsername(context.Context, *authv1.LoginViaUsernameRequest) (*authv1.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginViaUsername not implemented")
}

func (s *Server) Logout(context.Context, *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
