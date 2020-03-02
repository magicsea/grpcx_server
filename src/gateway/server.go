package main

import (
	_ "context"

	mid "github.com/grpc-ecosystem/go-grpc-middleware"

	"log"
	"net"
	_ "runtime/debug"

	"pb"
	"share"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/status"
)

//LoggingInterceptorfunc
func LoggingInterceptorfunc(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	log.Printf("gRPC stream method: %s, %v", info.FullMethod, stream)
	err := handler(srv, stream)
	log.Printf("gRPC stream method end: %s, %v", info.FullMethod, stream)
	return err
}

func main() {
	share.Init("gate1")

	conf := share.GetServerConf("name")
	log.Printf("gateway start:%+v", conf)

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
		mid.WithStreamServerChain(
			LoggingInterceptorfunc,
		))

	pb.RegisterLoginServiceServer(clientService, &Loginservice{})

	addr := share.GetServerConf("listenClient").(string)
	println("client addr:", addr)
	lis, err := net.Listen("tcp", addr)
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
	addr := share.GetServerConf("listenAddr").(string)
	println("server addr:", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net.Listen server err: %v", err)
	}

	server.Serve(lis)

}
