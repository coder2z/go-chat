package websocketConn

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Websocket struct {
	conn      *websocket.Conn
	inChan    chan []byte
	ontChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex
	isClosed bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Websocket, err error) {
	conn = &Websocket{
		conn:      wsConn,
		inChan:    make(chan []byte, 1000), //接受消息管道
		ontChan:   make(chan []byte, 1000), //输出消息管道
		closeChan: make(chan byte, 1),      //关闭通信管道
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

func (w *Websocket) ReadMessage() (data []byte, err error) {
	select {
	case data = <-w.inChan:
	case <-w.closeChan:
		err = errors.New("连接已被关闭")
	}
	return
}

func (w *Websocket) WriteMessage(data []byte) (err error) {
	select {
	case w.ontChan <- data:
	case <-w.closeChan:
		err = errors.New("连接已被关闭")
	}
	return
}

func (w *Websocket) Close() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if !w.isClosed {
		w.conn.Close()
		close(w.closeChan)
		w.isClosed = true
	}
}

// 获取 发现消息管道中的数据，发送消息
func (w *Websocket) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = w.conn.ReadMessage(); err != nil {
			w.Close()
			return
		}
		select {
		case w.inChan <- data:
		case <-w.closeChan:
			w.Close()
			return
		}
	}
}

// 获取 发现消息管道中的数据，发送消息
func (w *Websocket) writeLoop() {
	var (
		data []byte
	)
	for {
		select {
		case data = <-w.ontChan:
		case <-w.closeChan:
			w.Close()
			return
		}
		if err := w.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			w.Close()
			return
		}
	}
}
