package user

import "net"

type User struct {
	Name    string
	Addr    string
	Ch_user chan string
	conn    net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Ch_user: make(chan string),
		conn:    conn,
	}
	//启动监听当前User的channel
	go user.ListenMessage()
	return user
}

// 监听收到的消息 Ch_user
func (this *User) ListenMessage() {
	for {
		msg := <-this.Ch_user
		this.conn.Write([]byte(msg + "\r\n")) //write 写给客户端
	}
}
func (this *User) GetInformainton(infor string) {
	this.conn.Write([]byte(infor + "\r\n"))
}
