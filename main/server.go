package main

import (
	"demo_rpc/demo"
	"demo_rpc/thrift_gen"
	"demo_rpc/vendor/thrift"
	"fmt"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:19090"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	processor := thrift_gen.NewDemoServiceProcessor(&demo.DemoService{})

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
