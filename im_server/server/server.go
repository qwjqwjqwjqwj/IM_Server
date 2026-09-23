package server

import (
	"fmt"
	"im_server/user"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*user.User
	mapLock   sync.RWMutex

	//消息广播 channle
	ch_serverMessage chan string
}

// 创建一个server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:               ip,
		Port:             port,
		OnlineMap:        make(map[string]*user.User),
		ch_serverMessage: make(chan string),
	}
	return server

}

// 监听广播消息
func (this *Server) listnBoardcast() {
	for {
		BoardcastMessage := <-this.ch_serverMessage
		this.mapLock.Lock()
		for _, uStore := range this.OnlineMap {
			uStore.Ch_user <- BoardcastMessage
		}
		this.mapLock.Unlock()
	}

}

// 广播消息
func (this *Server) Boardcast(user *user.User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.ch_serverMessage <- sendMsg

}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("Establish contact...")
	user := user.NewUser(conn)

	//用户上线 加入服务器表中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	timer := time.NewTimer(10 * time.Second) //用户活跃
	isLive := make(chan bool)

	//广播上线消息
	this.Boardcast(user, " online")

	//接收用户消息
	go func() {
		buf := make([]byte, 4096)
		for {

			n, err := conn.Read(buf) //n是返回比特长度
			if err != nil {
				if err == io.EOF { //数据流正常结束
					this.Boardcast(user, "offline")
				} else {
					fmt.Println("ERROR:", err)
				}
				fmt.Println("Exxxxxxxxxxxxxxxxxxxxxxxxxxit")

				// //去表
				// this.mapLock.Lock()
				// delete(this.OnlineMap, user.Name)
				// this.mapLock.Unlock()
				// //关闭channel
				// close(user.Ch_user)
				// //关闭连接
				// conn.Close()
				return

			}
			//提取用户消息
			timer.Reset(10 * time.Second)
			msg := strings.TrimRight(string(buf[:n]), "\r\n")
			isLive <- true //用户是活跃的
			if msg == "who" {
				user.GetInformainton("")
				this.mapLock.Lock()
				for name := range this.OnlineMap {
					user.GetInformainton(name + " is online")
				}
				this.mapLock.Unlock()
			} else {
				this.Boardcast(user, msg)
			}

		}
	}()

	//超时踢人
	for {
		select {
		case <-timer.C:
			user.GetInformainton("You haven't responded for a long time and have thus been reported.")
			//从表中删除用户
			this.mapLock.Lock()
			delete(this.OnlineMap, user.Name)
			this.mapLock.Unlock()
			//关闭用户channel
			close(user.Ch_user)
			//关闭用户连接
			conn.Close()
			return
		case <-isLive:

		}

	}
}

// 启动服务器接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Liten err:", err)
		return
	}
	defer listener.Close() //关闭端口监听，不再接受新连接
	fmt.Println("Server is running................")
	//监听广播 一旦有消息发给全部User
	go this.listnBoardcast()
	for {
		//accept
		conn, err := listener.Accept() //用户上线成功 没有用户是Accept一直阻塞
		if err != nil {
			fmt.Println("make connection failed----err:", err)
			continue
		}
		//do handler 用户上线
		go this.Handler(conn)
	}
}
