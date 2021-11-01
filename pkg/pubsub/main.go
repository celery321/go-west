package main

import (
	"context"
	"fmt"
	"github.com/dapr/components-contrib/pubsub"
	pubsub_kafka "github.com/dapr/components-contrib/pubsub/kafka"
	"github.com/dapr/kit/logger"
	"github.com/pkg/errors"
	"time"
)

var (
	logContrib = logger.NewLogger("dapr.contrib")
)


type Service struct {
	conn pubsub.PubSub
	m  pubsub.Metadata
}

func main() {
	s := New()
	request := &pubsub.PublishRequest {
		Data: []byte("aaaa2"),
		PubsubName: "aaa",
		Topic: "first",
		Metadata: s.m.Properties,
	}

	if err := s.pub(request); err != nil {
		panic(errors.Wrap(err, "pub"))
	}

	response := pubsub.SubscribeRequest{
		Topic: "first",
		Metadata: s.m.Properties,
	}

	if err := s.sub(response); err != nil {
		panic(errors.Wrap(err, "pub"))
	}

	select {

	}


}

func New() (s *Service) {
	s = &Service{
	}
	m := pubsub.Metadata{}
	m.Properties = map[string]string{
		"consumerGroup": "a",
		"clientID": "pub",
		"brokers": "10.1.40.61:9092",
		"authRequired": "false",
		"maxMessageBytes": "2048",
		"consumerID" : "sub",
	}
	k := pubsub_kafka.NewKafka(logContrib)
	if err := k.Init(m); err != nil {
		panic(errors.Wrap(err, "init"))
	}
	s.m = m
	s.conn = k
	return s

}
func (s *Service) pub(r *pubsub.PublishRequest)  error {
	 if err := s.conn.Publish(r); err != nil{
	 	return errors.Wrap(err, "pub")
	 }
	 return nil
}

func (s *Service) sub (q pubsub.SubscribeRequest) error {

	time.Sleep(3 * time.Second)
	if err := s.conn.Subscribe(q,
		func(ctx context.Context, msg *pubsub.NewMessage) error {
		fmt.Printf("msg===%s\n", msg.Data)
		return nil
	}); err != nil {
		return errors.Wrap(err, "sub")
	}

	return nil
}

