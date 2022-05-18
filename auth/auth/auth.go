/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-18 22:25:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-18 22:54:34
 */
package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}

	return &authpb.LoginResponse{
		//openID不能给客户看，调试的时候可以用，正式的时候要改回状态码
		AccessToken: "token for open id" + openID,
		ExpiresIn:   7200,
	}, nil
}
