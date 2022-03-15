package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// Service implements auth service.
type Service struct {
	OpenIDResolver 	OpenIDResolver
	TokenGenerator	TokenGenerator
	TokenExpire		time.Duration
	Mongo			*dao.Mongo
	Logger 			*zap.Logger
}

// OpenIDResolver 输入code，输出openID（即用户唯一标识）
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// TokenGenerator 由accountID产生相应的Token
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

// Login logs a user in.
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIDResolver.Resolve(req.Code)	// 由小程序端的code，从微信的接口获得用户唯一标识openID
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot open id: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(c, openID)	// 由openID获得数据库中的accountID
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	tkn, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)	// 获得token（JWT方法）
	if err != nil {
		s.Logger.Error("cannot generate token for openID.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{AccessToken: tkn, ExpiresIn: int32(s.TokenExpire.Seconds())}, nil
}

