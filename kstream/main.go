package main

import (
	"errors"
	"fmt"
	"time"

	"git.chotot.org/go-kafka-consumer/logger"

	"git.chotot.org/data/kstream"
)

func main() {
	log := logger.GetLogger("test-script")
	kstConfig := kstream.ProcessConfig{
		AppID:           "dat1",
		ProcessName:     "test",
		Brokers:         []string{"10.60.3.187:9092", "10.60.3.188:9092"},
		Registry:        "http://10.60.3.131:32100",
		Subscribe:       []string{"AdStatsStringJSON"},
		ISerder:         kstream.AVRO,
		Log:             log,
		RewindEvent:     false,
		NormalizeRecord: true,
	}
	ops := &kstream.SimpleOp{
		Name: "Printer",
		Op: func(i *kstream.Record) ([]*kstream.Record, error) {
			fmt.Print(i)
			if _, ok := i.Val.(map[string]interface{}); !ok {
				return nil, errors.New("Malform data")
			}
			return nil, nil
		}}

	proc, err := kstream.NewProcessor(kstConfig, ops)
	if err != nil {
		panic(err)
	}
	if err := proc.Init(nil); err != nil {
		panic(err)
	}

	// Debug code because Run() will not return unless process.Close() is invoked
	go func() {
		time.Sleep(30000 * time.Millisecond)
		proc.Stop()
	}()
	// end debug
	proc.Run()
}
