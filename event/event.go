package event

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var (
	uploadListener jetstream.Consumer
	Duration       = time.Second * 30
)

func init() {
	initNats()
	n, err := nats.Connect(fmt.Sprintf("nats://127.0.0.1:%d", port), nats.MaxReconnects(5))
	if err != nil {
		panic(err)
	}

	js, err := jetstream.New(n)
	if err != nil {
		log.Print("can not create jetstream connection")
		log.Fatal(err)
	}
	// 创建消费者
	consumerConfig := jetstream.ConsumerConfig{
		Name:           "upload",
		AckPolicy:      jetstream.AckExplicitPolicy,
		DeliverPolicy:  jetstream.DeliverNewPolicy,
		MaxDeliver:     maxDeliver,
		FilterSubjects: []string{"upload"},
	}
	consumer, err := js.CreateOrUpdateConsumer(context.TODO(), "file", consumerConfig)
	if err != nil {
		log.Printf("can not create consumer")
		log.Fatal(err)
	}
	uploadListener = consumer
	go func() {
		for {
			batch, err := uploadListener.Fetch(1)
			if err != nil {
				continue
			}
			msg := batch.Messages()
			for m := range msg {
				meta, _ := m.Metadata()
				if meta.NumDelivered == 1 {
					m.NakWithDelay(Duration)
				} else {
					go DeleteFile(m)
				}
			}
		}
	}()
}

func PushFileUploadEvent(path string) {
	if err := nc.Publish("upload", []byte(path)); err != nil {
		log.Print(err.Error())
	}
}
func OnFileUpload(path string, duration time.Duration, callback func(path string)) {

}

func DeleteFile(msg jetstream.Msg) {
	defer msg.Ack()
	log.Printf("delete file: %s", string(msg.Data()))
	// check file exists
	if _, err := os.Stat(string(msg.Data())); os.IsNotExist(err) {
		return
	}
	if err := os.Remove(string(msg.Data())); err != nil {
		log.Printf("can not delete file: %s", err.Error())
	}
}
