package grpc

import authv1 "github.com/andrey-tsyn/my-giga-wishes/api-contracts/gen/go/auth-service"

type server struct {
	authv1.UnimplementedAuthServer
}
