package main

import (
	_ "context"
	"github.com/grpc-ecosystem/go-grpc-middleware"

	"log"
	"net"
	_ "runtime/debug"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/status"

	"pb"
)

func LoggingInterceptorfunc(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	log.Printf("gRPC stream method: %s, %v", info.FullMethod, stream)
	err := handler(srv, stream)
	log.Printf("gRPC stream method end: %s, %v", info.FullMethod, stream)
	return err
}

func main() {

	conf := GetConfig()
	log.Printf("gateway start:%+v  server:%+v", conf.Project, conf.Server)

	go listenServer()

	listenClient()
}

//listen client
func listenClient() {
	//unknownOpt := grpc.UnknownServiceHandler(func(srv interface{}, stream grpc.ServerStream) error {
	//	// little bit of gRPC internals never hurt anyone
	//	fullMethodName, ok := grpc.MethodFromServerStream(stream)
	//	if !ok {
	//		return grpc.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	//	}
	//	//md := stream.Context()
	//	//newMD := context.WithValue(md,"uid","111")
	//
	//	f := &pb.SearchRequest{}
	//	if err := stream.RecvMsg(f); err != nil {
	//		return err
	//	}
	//	err2 := stream.SendMsg(&pb.SearchResponse{Response:f.Request})
	//
	//	p,_ := peer.FromContext(stream.Context())
	//	fmt.Printf("unknow func.............addr=%v\n",p.Addr)
	//	fmt.Println("recv:",fullMethodName,f,err2)
	//	return nil
	//})
	//_ = unknownOpt
	clientService := grpc.NewServer(
		grpc.CustomCodec(RawCodec()),
		grpc.UnknownServiceHandler(TransparentHandler(RoundRobinConfigRouter)),
		grpc_middleware.WithStreamServerChain(
			LoggingInterceptorfunc,
		))

	pb.RegisterLoginServiceServer(clientService, &Loginservice{})

	lis, err := net.Listen("tcp", GetConfig().Server.ListenClient)
	if err != nil {
		log.Fatalf("net.Listen client err: %v", err)
	}

	clientService.Serve(lis)

}

//listenServer
func listenServer() {
	server := grpc.NewServer()
	pb.RegisterGateServiceServer(server, &GateService{})
	//listen server
	lis, err := net.Listen("tcp", GetConfig().Server.ListenServer)
	if err != nil {
		log.Fatalf("net.Listen server err: %v", err)
	}

	server.Serve(lis)

}
