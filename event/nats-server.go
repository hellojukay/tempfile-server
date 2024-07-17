package event

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var (
	maxDeliver = 5
	port       = 4222
	nc         *nats.Conn
)

func initNats() {
	opts := &server.Options{
		Host:      "localhost",
		Port:      port,
		JetStream: true,
	}
	// Initialize new server with options
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	ns.Start()
	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		log.Printf("NATS server not ready for connections in 4 seconds")
	}

	n, err := nats.Connect(fmt.Sprintf("nats://127.0.0.1:%d", port), nats.MaxReconnects(5))
	if err != nil {
		panic(err)
	}
	nc = n
	js, err := jetstream.New(n)
	if err != nil {
		log.Printf("can init stream connection")
		log.Fatal(err)
	}
	// create nats-stream
	js.CreateOrUpdateStream(context.TODO(), jetstream.StreamConfig{
		Name: "file",
		Subjects: []string{
			"upload",
			"delete",
		},
	})

}
