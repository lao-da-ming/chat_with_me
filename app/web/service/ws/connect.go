package ws

import (
	"bytes"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type WsController struct {
	Logger *zap.Logger
}

func NewWsController(Logger *zap.Logger) *WsController {
	return &WsController{Logger: Logger}
}

var (
	//当前连接数
	curConnCount int32
	//最大连接数
	maxConnCount int32 = 2
	//心跳时间
	heartbeatInterval = time.Second * 15
	//协议升级配置
	upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Second * 20,
		ReadBufferSize:   2048,
		WriteBufferSize:  2048,
		WriteBufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		Subprotocols: nil,
		Error:        nil,
		CheckOrigin: func(r *http.Request) bool { //跨域设置
			return true
		},
		EnableCompression: true,
	}
	//链接池
	ConnectPool sync.Map
)

type ConnObj struct {
	//唯一标识
	uniqueId string
	//链接对象
	conn *websocket.Conn
	//关闭通道
	closeChan chan struct{}
	//收消息缓冲区
	receiveChan chan []byte
	//发消息缓冲区
	sendChan chan []byte
	//单例操作
	once sync.Once
}

// 升级协议连接websocket
func (ws *WsController) Connect(c *gin.Context) {
	c.Writer.Header().Set("Sec-WebSocket-Protocol", strings.Split(c.Request.Header.Get("Sec-WebSocket-Protocol"), ",")[0])
	conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "连接失败" + err.Error()})
		return
	}
	if curConnCount >= maxConnCount {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("超出最大连接数"))
		_ = conn.Close()
		return
	}
	uniqueKey := hex.EncodeToString(uuid.NewV4().Bytes())
	obj := &ConnObj{
		uniqueId:    uniqueKey,
		conn:        conn,
		closeChan:   make(chan struct{}, 1),
		receiveChan: make(chan []byte, 100),
		sendChan:    make(chan []byte, 100),
		once:        sync.Once{},
	}
	ConnectPool.Store(uniqueKey, obj)
	defer obj.close()
	log.Println("连接建立，标识:", uniqueKey)
	//统计+1
	atomic.AddInt32(&curConnCount, 1)
	//统计-1
	defer atomic.AddInt32(&curConnCount, -1)
	//接收客户端信息写入缓冲队列
	go obj.readLoop()
	//读取缓冲队列数据发送到客户端
	go obj.writeLoop()
	//处理缓冲队列数据
	go obj.doTask()
	//心跳检测
	obj.heartbeat()
}

// 处理消息业务
func (co *ConnObj) doTask() {
	for {
		select {
		case <-co.closeChan:
			return
		case msg := <-co.receiveChan:
			//暂时不处理，发回去
			co.WriteToSendChan(msg)
		}
	}
}

// 发送到写出通道
func (co *ConnObj) WriteToSendChan(msg []byte) {
	select {
	case <-co.closeChan:
	case co.sendChan <- msg:
	default:
		co.close()
	}
}

// 从缓冲队列取消息发送到客户端
func (co *ConnObj) writeLoop() {
	for {
		select {
		case <-co.closeChan:
			return
		case msg := <-co.sendChan:
			if err := co.writeMessage(msg); err != nil {
				log.Println("write message err:", err.Error())
				return
			}
		}
	}
}

// 写道客户端
func (co *ConnObj) writeMessage(msg []byte) error {
	if err := co.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		co.close()
		return err
	}
	return nil
}

// 接收客户端信息写入缓冲队列
func (co *ConnObj) readLoop() {
	for {
		msgType, data, err := co.conn.ReadMessage()
		if err != nil {
			log.Println("read msg error:", err)
			co.close()
			return
		}
		select {
		case <-co.closeChan:
			return
		default:
			switch msgType {
			case websocket.TextMessage, websocket.BinaryMessage:
				select {
				case co.receiveChan <- data:
				default:
					log.Println("接收消息队列溢出，关闭链接")
					co.close()
					return
				}
			case websocket.CloseMessage: //关闭消息
				log.Println("receive close msg")
				co.close()
				return
			case websocket.PingMessage:
				log.Println("receive ping msg")
			case websocket.PongMessage: //心跳消息
				log.Println("receive pong msg")
			default:
				log.Println("未知消息类型")
			}
		}
	}
}

// 关闭链接
func (co *ConnObj) close() {
	co.once.Do(func() {
		_ = co.conn.WriteMessage(websocket.TextMessage, []byte("服务端关闭连接"))
		_ = co.conn.Close()
		close(co.closeChan)
		close(co.receiveChan)
		close(co.sendChan)
		ConnectPool.Delete(co.uniqueId)
		log.Println("连接关闭")
	})
}

// 心跳
func (co *ConnObj) heartbeat() {
	for {
		select {
		case <-co.closeChan:
			return
		case <-time.After(heartbeatInterval):
			//这个时间必须大于定时器时间
			if err := co.conn.SetWriteDeadline(time.Now().Add(heartbeatInterval + 4)); err != nil {
				log.Println("设置写超时时间失败，err=", err.Error())
				return
			}
			if err := co.conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				log.Println("发送心跳失败", err.Error())
				return
			}
		}
	}
}
