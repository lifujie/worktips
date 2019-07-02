package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"

	"logger"
)

const (
	// MQName 连接的消息总线类型
	MQName = "rabbitMQ"
	// ExchangeName 消息分发策略名
	ExchangeName = "XMDB_INNER_SPICAL_TOPIC_EXCHANGE"
	// ExchangeCategory exchange发送给绑定queue的策略
	ExchangeCategory = "topic"
	// QueueNamePrefix 从那个队列中接收消息，向那个队列中发送消息， 使用的消息队列真正名字是Prefix+IP
	QueueNamePrefix = "XMDB_INNER_SPICAL_QUEUE_BOOT"
	// BindingKey 发送到那个queue
	BindingKey = "xmdb.neo4j.*"
	// RoutingKey 消息传递给那个exchange
	RoutingKey = "xmdb.inner.special.data"
)

const (
	// SessionScanInterval = 24 * time.Hour
	SessionScanInterval = 1 * time.Minute
)

// Session mq会话结构
type Session struct {
	conn             *amqp.Connection
	ch               *amqp.Channel
	mqName           string
	exchangeName     string
	exchangeCategory string
	queueName        string
	bindingKey       string
	// 当前仅仅使用接收消息功能，所以routingKey不用配置
	routingKey string
	// 其他配置以后补充

	log logger.Logger
}

// Close 关闭mq会话
func (s *Session) Close() error {
	if s.conn == nil {
		return nil
	}
	return s.conn.Close()
}

// InitSessionConfig 初始化MQ链接配置
func InitSessionConfig(log logger.Logger, XmdbSerAddr, exchangeName, queueNamePrefix, bindingKey string) (session *Session) {
	if exchangeName == "" {
		exchangeName = ExchangeName
	}

	//queueNamePrefix = "XMDB_INNER_SPICAL_QUEUE_YZK"
	if queueNamePrefix == "" {
		queueNamePrefix = fmt.Sprintf("%s_%s", QueueNamePrefix, XmdbSerAddr)
	}

	if bindingKey == "" {
		bindingKey = BindingKey
	}

	return &Session{
		mqName:           MQName,
		exchangeName:     exchangeName,
		queueName:        queueNamePrefix,
		bindingKey:       bindingKey,
		exchangeCategory: ExchangeCategory,
		routingKey:       RoutingKey,
		log:              log,
	}
}

// ConstructURL 构造URL
func ConstructURL(username, password, url, port string) (URL string) {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, url, port)
}

// Redial 拨号器
func (s *Session) Redial(url string) *Session {
	conn, err := amqp.Dial(url)
	if err != nil {
		s.log.Debugf("Failed to connect to RabbitMQ: %s\n", err.Error())
		conn.Close()
		return nil
	}

	ch, err := conn.Channel()
	if err != nil {
		s.log.Debugf("Failed to open a channel: %s\n", err.Error())
		ch.Close()
		return nil
	}

	s.conn = conn
	s.ch = ch
	return s
}

// OriginalMessage 消息
type OriginalMessage []byte

// Subscribe 消息订阅
func (s *Session) Subscribe(messages chan OriginalMessage) bool {
	q, err := s.ch.QueueDeclare(
		s.queueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		s.log.Debugf("Failed to declare a queue: %s\n", err.Error())
		return false
	}

	err = s.ch.QueueBind(s.queueName, s.routingKey, s.exchangeName, false, nil)
	if err != nil {
		s.log.Debugf("Failed to bind exchange: %s\n", err.Error())
		return false
	}

	msgs, err := s.ch.Consume(
		q.Name,      // queue
		"cloudboot", // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		s.log.Debugf("Failed to register a consumer: %s\n", err.Error())
		return false
	}

	idleExitTicker := time.NewTicker(SessionScanInterval)

	// 用于测试的变量
	i := 0
	idleExit := false
	for {
		select {
		case <-idleExitTicker.C:
			if idleExit {
				return true
			}
			idleExit = true
		case msg := <-msgs:
			idleExit = false
			var message Message
			_ = json.Unmarshal(msg.Body, &message)
			//s.log.Debugf("%s\n", message.DataType)
			messages <- OriginalMessage(msg.Body)
			//s.log.Debugf("%v\n", msg)
			s.ch.Ack(msg.DeliveryTag, false)

			// 仅仅用于测试
			i++
			if i > 100 {
				time.Sleep(10 * time.Second)
			}
		}
	}
}

var (
	// MsgBrokerIP 消息总线IP
	MsgBrokerIP = "10.0.20.40"
	// MsgBrokerPort 消息总线Port
	MsgBrokerPort = "5672"
)

func main() {
	url := ConstructURL("guest", "guest", MsgBrokerIP, MsgBrokerPort)

	sess := InitSessionConfig(logger.Logger{}, MsgBrokerIP, "", "", "")
	// 与MQ建立链接
	sess = sess.Redial(url)

	exitChan := make(chan int, 1)

	// 监听MQ 获取消息
	messages := make(chan OriginalMessage, 1)
	go func() {
		for {
			select {
			case message := <-messages:
				var mg Message
				_ = json.Unmarshal(message, &mg)
				//fmt.Printf("message: %s\n", mg)
				switch mg.DataType {
				case "CI":
					var mgci CIData
					if err := json.Unmarshal(mg.DataObject.([]byte), &mgci); err != nil {
						fmt.Printf("解析CI数据出错: %s", err.Error())
					} else {
						fmt.Printf("Value: %v\n", mgci)
					}
				case "REL":
				}
			case <-exitChan:
				return
			}

		}
	}()

	// 监听获取消息
	if sess.Subscribe(messages) {
		sess.Close()
	}
	//return err
	time.Sleep(100 * time.Second)
	return
}
