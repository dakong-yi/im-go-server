package client

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	conn2 "github.com/dakong-yi/im-go-server/pkg/conn"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //判断当前client的模式
}

func newClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999, // 设置flay默认值，否则flag默认为int整型
	}
	//创建链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	//返回客户端
	return client
}

// client菜单栏的输出，并获取flag输入
func (client *Client) menu() bool {
	var flag int

	fmt.Println("input 1 into public chat")
	fmt.Println("input 2 into private chat")
	fmt.Println("input 3 into rename")
	fmt.Println("input 0 into exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("invalid input integer")
		return false
	}
}

// DealResponse 监听server回应的消息，直接显示到标准输出
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn) // 永久阻塞监听
	/*
		上面一句相当于如下for循环一直从conn中读取，然后输出到终端
		//for {
		//	buf := make([]byte, 4096)
		//	client.conn.Read(buf)
		//	fmt.Println(string(buf))
		//}
	*/

}

func (client *Client) Run() {

}

// 尝试从终端命令行解析IP和Port创建客户端
var serverIp string
var serverPort int

// 文件的初始化函数
// 命令的格式  ./client.exe -ip 127.0.0.1 -port 8888
func init() {
	//属于初始化工作，一般放在init中
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "set server port(default:8888)")
}

func main() {
	//通过命令行解析
	flag.Parse()
	client := newClient(serverIp, serverPort)
	//client := newClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println("------- connect server error------")
		return
	}

	fmt.Println("-------- connect server success ------")
	//buf := make([]byte, 4096)
	//n,err:= client.conn.Read(buf)
	//fmt.Println(n)
	//fmt.Println(err)
	//fmt.Println(string(buf))
	msg := conn2.Message{Data: "{\"user_id\":123,\"device_id\":123}", Type: 1, RequestId: 1}
	//str1 := []byte(`{"type":1,"request_id":1,"data":"{\"user_id\":123,\"device_id\":123}"}`)
	str1, _ := json.Marshal(msg)
	client.conn.Write(str1)
	//var str string
	for {
		//buffer := make([]byte, 1024)
		//n, err := client.conn.Read(buffer)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println(string(buffer[0:n]))

		type reply struct {
			Name string
		}
		rep := reply{}
		d := json.NewDecoder(client.conn)
		err := d.Decode(&rep)
		//err = json.Unmarshal(buffer, &rep)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(rep.Name)
		//fmt.Scanln(&str)
		//client.conn.Write([]byte(str))
	}

	//按理说启动client.Run()方法之后，服务器返回相应的处理结果，
	//主go程会阻塞在Run方法，如果使用主go程中的Run方法接受返回消息，就会变成串行执行
	//无法同一时刻满足其他的业务，而run应该跟dealResponse应该是并行的
	//所以提供一个新的go程只处理server回应的信息
	//client.DealResponse()
}
