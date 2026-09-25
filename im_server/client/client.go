package client

import "net"

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func (this *Client) NewClient(ServerIP string, ServerPort int) *Client {
	client := &Client{
		ServerIP:   ServerIP,
		ServerPort: ServerPort,
	}

	net.Dial("tcp", client.ServerIP+string(client.ServerPort))

}
