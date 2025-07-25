package mqtt

import (
	"crypto/tls"
	"fmt"
	"github.com/google/wire"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ProviderSet = wire.NewSet(NewMQTT)

// 定义 MQTT 客户端配置
type MQTTConfig struct {
	Broker   string
	Port     int
	ClientID string
	Username string
	Password string
}

// MQTTClient 结构体封装 MQTT 功能
type MQTTClient struct {
	client mqtt.Client
	config MQTTConfig
}

// NewMQTTClient 创建新的 MQTT 客户端
func NewMQTT() *MQTTClient {
	config := MQTTConfig{
		Broker:   "broker.emqx.io", // 公共 MQTT 代理
		Port:     1883,
		ClientID: "go-mqtt-client-" + fmt.Sprint(time.Now().Unix()),
		Username: "emqx",   // 公共代理的用户名
		Password: "public", // 公共代理的密码
	}
	opts := mqtt.NewClientOptions()
	brokerURL := fmt.Sprintf("tcp://%s:%d", config.Broker, config.Port)
	opts.AddBroker(brokerURL)
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	// 设置 TLS 配置（可选）
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts.SetTLSConfig(tlsConfig)

	// 设置连接回调
	opts.OnConnect = func(c mqtt.Client) {
		log.Printf("成功连接到 MQTT 代理: %s", brokerURL)
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("MQTT 连接丢失: %v", err)
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("MQTT连接失败", token.Error())
	}
	return &MQTTClient{
		client: client,
		config: config,
	}
}

// Publish 发布消息到指定主题
func (m *MQTTClient) Publish(topic, message string, qos byte) error {
	token := m.client.Publish(topic, qos, false, message)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("发布失败: %v", token.Error())
	}
	log.Printf("已发布消息到 %s: %s", topic, message)
	return nil
}

// Subscribe 订阅主题并设置消息处理函数
func (m *MQTTClient) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) error {
	if token := m.client.Subscribe(topic, qos, handler); token.Wait() && token.Error() != nil {
		return fmt.Errorf("订阅失败: %v", token.Error())
	}
	log.Printf("已订阅主题: %s", topic)
	return nil
}

// Subscribe 取消订阅主题
func (m *MQTTClient) UnSubscribe(topic ...string) error {
	if token := m.client.Unsubscribe(topic...); token.Wait() && token.Error() != nil {
		return fmt.Errorf("取消订阅失败: %v", token.Error())
	}
	log.Printf("已取消订阅主题: %s", topic)
	return nil
}
