/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-18 22:36:25
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-18 22:44:11
 */
package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true,
			OrigName:    true,
		},
	))

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, "localhost:8081",
		[]grpc.DialOption{
			grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal("cannot register auth service: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
