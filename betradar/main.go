package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badboyd/go-uof-sdk"
	"github.com/badboyd/go-uof-sdk/api"
)

const (
	// EnvToken is API Key for UOF
	envToken = "UOF_TOKEN"
)

var (
	token string
)

func init() {
	token = env(envToken)
}

func env(name string) string {
	e, ok := os.LookupEnv(name)
	if !ok {
		log.Printf("env %s not found", name)
	}
	return e
}

func exitSignal() context.Context {
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		//SIGINT je ctrl-C u shell-u, SIGTERM salje upstart kada se napravi sudo stop ...
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		stop()
	}()
	return ctx
}

func main() {
	exitSig := exitSignal()

	betradarCli, err := api.Staging(exitSig, token)
	if err != nil {
		log.Fatal(err)
	}

	// betradarCli.Ping().Error()
	r, err := betradarCli.Fixture(uof.LangEN, uof.NewEventURN(23216921))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(r))

	type scheduleRsp struct {
		Fixture     uof.Fixture `xml:"fixture,omitempty" json:"fixture,omitempty"`
		GeneratedAt time.Time   `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
	}

	res := scheduleRsp{}
	xml.Unmarshal(r, &res)

	b, _ := json.MarshalIndent(res.Fixture, "", "\t")

	log.Println(string(b))
}
