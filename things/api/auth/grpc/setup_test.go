// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package grpc_test

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/cloustone/pandas/mainflux"
	"github.com/cloustone/pandas/things"
	grpcapi "github.com/cloustone/pandas/things/api/auth/grpc"
	"github.com/cloustone/pandas/things/mocks"
	"github.com/opentracing/opentracing-go/mocktracer"
	"google.golang.org/grpc"
)

const (
	port  = 8080
	token = "token"
	wrong = "wrong"
	email = "john.doe@email.com"
)

var svc things.Service

func TestMain(m *testing.M) {
	startServer()
	code := m.Run()
	os.Exit(code)
}

func startServer() {
	svc = newService(map[string]string{token: email})
	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	server := grpc.NewServer()
	mainflux.RegisterThingsServiceServer(server, grpcapi.NewServer(mocktracer.New(), svc))
	go server.Serve(listener)
}

func newService(tokens map[string]string) things.Service {
	auth := mocks.NewAuthService(tokens)
	conns := make(chan mocks.Connection)
	thingsRepo := mocks.NewThingRepository(conns)
	channelsRepo := mocks.NewChannelRepository(thingsRepo, conns)
	chanCache := mocks.NewChannelCache()
	thingCache := mocks.NewThingCache()
	idp := mocks.NewIdentityProvider()

	return things.New(auth, thingsRepo, channelsRepo, chanCache, thingCache, idp)
}