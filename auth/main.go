/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-18 23:03:16
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-18 23:06:38
 */
package main

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/wechat"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//logger, err := newZap
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {

	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "",
			AppSecret: "",
		},
		//Logger: ,
	})
}

func newZapLogger() (*zap.Logger, error) {
	return nil, nil
}
