package main

import (
	"demo_rpc/thrift_gen"
	"demo_rpc/vendor/thrift"
	"fmt"
	"net"
	"os"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := thrift_gen.NewDemoServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	for i := 0; i < 10; i++ {
		r1, e1 := client.Hello(fmt.Sprintf(`bughou %d`, i))
		fmt.Println(r1, e1)
	}
}
