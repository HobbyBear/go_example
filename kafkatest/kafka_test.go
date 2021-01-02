package kafkatest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rcrowley/go-metrics"
)

var c sarama.Client

func init() {
	var err error
	c, err = sarama.NewClient([]string{"localhost:9092"}, newSaramaConfig())
	if err != nil{
		log.Fatal(err)
	}
}

func TestProducer(t *testing.T)  {
	p ,err:= sarama.NewSyncProducerFromClient(c)
	if err != nil{
		log.Fatal(err)
	}
	_,_,err = p.SendMessage(&sarama.ProducerMessage{
		Topic:     "test",
		Key:       sarama.StringEncoder("testKey"),
		Value:     sarama.StringEncoder("love"),
		Headers:   nil,
		Metadata:  nil,
		Timestamp: time.Now(),
	})
	if err != nil{
		log.Fatal(err)
	}
}
type groupHandler struct {
}
func (gh *groupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (gh *groupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (gh *groupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println(fmt.Sprintf("topic:%s,value:%s,offset:%d",message.Topic,message.Value,message.Offset))
		return errors.New("测试错误")
	}

	return nil
}

func TestConsumer(t *testing.T)  {
	cg ,err :=sarama.NewConsumerGroupFromClient("group_0",c)
	if err != nil{
		log.Fatal(err)
	}
	var g groupHandler
	for {
		err = cg.Consume(context.TODO(),[]string{"test"},&g)
		if err != nil{
			log.Println(err)
		}
	}
}


func newSaramaConfig() *sarama.Config {
	c := sarama.NewConfig()

	c.Admin.Retry.Max = 5
	c.Admin.Retry.Backoff = 100 * time.Millisecond
	c.Admin.Timeout = 3 * time.Second

	c.Net.MaxOpenRequests = 5
	c.Net.DialTimeout = 30 * time.Second
	c.Net.ReadTimeout = 30 * time.Second
	c.Net.WriteTimeout = 30 * time.Second
	c.Net.SASL.Handshake = true
	c.Net.SASL.Version = sarama.SASLHandshakeV0

	c.Metadata.Retry.Max = 3
	c.Metadata.Retry.Backoff = 250 * time.Millisecond
	c.Metadata.RefreshFrequency = 10 * time.Minute
	c.Metadata.Full = true

	c.Producer.MaxMessageBytes = 1048576
	c.Producer.RequiredAcks = sarama.WaitForLocal
	c.Producer.Timeout = 10 * time.Second
	c.Producer.Return.Successes = true
	c.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	c.Producer.Retry.Max = 3
	c.Producer.Retry.Backoff = 100 * time.Millisecond
	c.Producer.Return.Errors = true
	c.Producer.CompressionLevel = sarama.CompressionLevelDefault

	c.Consumer.Fetch.Min = 1
	c.Consumer.Fetch.Default = 1024 * 1024
	c.Consumer.Retry.Backoff = 2 * time.Second
	c.Consumer.MaxWaitTime = 250 * time.Millisecond
	c.Consumer.MaxProcessingTime = 100 * time.Millisecond
	c.Consumer.Return.Errors = true
	c.Consumer.Offsets.AutoCommit.Enable = true
	c.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	c.Consumer.Offsets.Initial = sarama.OffsetNewest
	c.Consumer.Offsets.Retry.Max = 3

	c.Consumer.Group.Session.Timeout = 10 * time.Second
	c.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	c.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	c.Consumer.Group.Rebalance.Retry.Max = 4
	c.Consumer.Group.Rebalance.Retry.Backoff = 2 * time.Second

	c.ChannelBufferSize = 256
	c.Version = sarama.V2_6_0_0
	c.MetricRegistry = metrics.NewRegistry()

	return c
}
