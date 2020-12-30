package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	fr "rpc.tekoapis.com/rpc/fulfillment_rules"
)

func main() {
	// Register kuberesolver to grpc before calling grpc.Dial
	// kuberesolver.RegisterInCluster()

	// it is same as
	// resolver.Register(kuberesolver.NewBuilder(nil /*custom kubernetes client*/, "kubernetes"))
	//\	// grpc.Dial("", grpc.WithInsecure())

	// resolver.SetDefaultScheme("dns")

	// cc, err := grpc.DialContext(context.Background(),
	// 	"127.0.0.1:10443",
	// 	grpc.WithInsecure(),
	// )

	cc, err := grpc.DialContext(context.Background(),
		"dns:///fulfillment-rules-headless-service.orders-management:9090",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	cli := fr.NewSellerServiceClient(cc)
	if cli == nil {
		panic("Cannot init client")
	}

	// cli := orders.NewOrderCapturingServiceClient(cc)
	// res, err := cli.GetInternalOrderByID(context.Background(), &orders.GetInternalOrderByIDRequest{
	// 	OrderId: 2687534084527752248,
	// })

	// // fmt.Printf("%+v\n", res)

	for i := 0; i < 10; i++ {
		res, err := cli.ListSeller(context.Background(), &fr.ListSellerRequest{})
		if err != nil {
			panic(err)
		}

		fmt.Printf("res: %+v\n", res)
		time.Sleep(1 * time.Second)
	}
}
