package test

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"

	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

func TestServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"localhost:13000",
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	payload := `{"limit": 100,"next": 11}`

	cli := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	defer cli.Reset()

	descSource := grpcurl.DescriptorSourceFromServer(ctx, cli)

	in := strings.NewReader(payload)
	options := grpcurl.FormatOptions{
		IncludeTextSeparator: true,
	}
	rf, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.Format("json"), descSource, in, options)
	if err != nil {
		panic(err)
	}

	o := new(bytes.Buffer)

	h := &grpcurl.DefaultEventHandler{
		Out:       o,
		Formatter: formatter,
	}
	headers := []string{}

	err = grpcurl.InvokeRPC(ctx, descSource, conn, "grpcserver.proto.v1.ListSomethingService/ListSomething", headers, h, rf.Next)
	if err != nil {
		panic(err)
	}

	t.Log(o.String())
}
