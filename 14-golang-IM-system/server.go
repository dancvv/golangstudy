package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	// 加锁
	mapLock sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

// 监听message官博消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		fmt.Println(msg)
		// 将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	// fmt.Println("链接建立成功")
	// 用户上限，同时广播
	user := NewUser(conn, this)
	// 用户上线，将用户加入到onlinemap中
	// this.mapLock.Lock()
	// this.OnlineMap[user.Name] = user
	// this.mapLock.Unlock()
	/* 版本更新 */
	// 用户上线
	user.Online()

	// 监听用户是否是活跃的channel
	isLive := make(chan bool)
	// 广播当前用户上线消息
	// this.BroadCast(user, "已上线")

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			// 下线
			// if n == 0 {
			// 	this.BroadCast(user, "下线")
			// 	return
			// }
			// 需要正确的使用下线功能
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}
			// 提取用户的消息，去除'\n'
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)
			// this.BroadCast(user, msg)

			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()
	// 当前handler阻塞
	/* 加入定时器，使得可以进行超时强踢*/
	for {
		select {
		case <-isLive:
			// 当前用户是获取的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 10):
			// 已经超时
			// 将当前的user强制的关闭，利用阻塞的机制
			user.SendMsg("你被踢了")

			// 销毁用的资源
			close(user.C)

			// 关闭链接
			conn.Close()

			// 退出当前handler
			return // runtime.goexist 也可以
		}
	}
	// select {}

}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听message的goroutine
	go this.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listenner accept err:", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}
}
