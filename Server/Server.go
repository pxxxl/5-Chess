package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	chess := [15][15]int{}
	roundFlag := 0
	buf := [1024]byte{}
	listen, err1 := net.Listen("tcp", "10.19.190.109:8098")
	if err1 != nil{
		fmt.Println("Listener build failed")
		return
	}
	conn1, err2 := listen.Accept()
	if err2 != nil {
		fmt.Printf("accept failed, err:%v\n", err2)
		return
	}
	fmt.Println("链接完成，链接信息如下：")
	fmt.Println("白方:"+conn1.RemoteAddr().String())
	_, err3 := conn1.Write([]byte("100"))
	if err3 != nil {
		fmt.Printf("accept failed, err:%v\n", err3)
		return
	}
	conn2, err4 := listen.Accept()
	if err4 != nil {
		fmt.Printf("accept failed, err:%v\n", err4)
		return
	}
	fmt.Println("链接完成，链接信息如下：")
	fmt.Println("黑方:"+conn2.RemoteAddr().String())
	_, err5 := conn2.Write([]byte("200"))
	if err5 != nil {
		fmt.Printf("accept failed, err:%v\n", err5)
		return
	}
	fmt.Println("游戏马上开始")
	time.Sleep(3*time.Second)
	roundFlag = 1
	sendMsg("000",conn1, conn2)
	//开始持续接收信息
	for{
		if roundFlag == 1{
			_, err6 := conn1.Read(buf[:])
			if err6 != nil{
				fmt.Printf("accept failed, err:%v\n", err6)
				return
			}
			chess[buf[1]-'a'][buf[2]-'A'] = 1
			win := checkWin(1,int(buf[1]-'a'),int(buf[2]-'A'),chess)
			if win{
				sendMsg("3"+string(buf[1])+string(buf[2]),conn1,conn2)
			}else{
				sendMsg("1"+string(buf[1])+string(buf[2]),conn1,conn2)
			}
		}else{
			_, err6 := conn2.Read(buf[:])
			if err6 != nil{
				fmt.Printf("accept failed, err:%v\n", err6)
				return
			}
			chess[buf[1]-'a'][buf[2]-'A'] = -1
			win := checkWin(-1,int(buf[1]-'a'),int(buf[2]-'A'),chess)
			if win{
				sendMsg("4"+string(buf[1])+string(buf[2]),conn1,conn2)
			}else{
				sendMsg("2"+string(buf[1])+string(buf[2]),conn1,conn2)
			}
		}
		roundFlag = -roundFlag
	}
}

func sendMsg(str string,a net.Conn, b net.Conn){
	_, e1 := a.Write([]byte(str))
	_, e2 := b.Write([]byte(str))
	if e1!=nil{
		fmt.Printf("accept failed, err:%v\n", e1)
		return
	}
	if e2!=nil{
		fmt.Printf("accept failed, err:%v\n", e2)
		return
	}

}

func checkWin(color int,x int, y int, chess [15][15]int)bool{
	winFlag := false
	switch{
	case checkLine(color,x,y,chess,1,0):winFlag = true
	case checkLine(color,x,y,chess,0,1):winFlag = true
	case checkLine(color,x,y,chess,1,1):winFlag = true
	case checkLine(color,x,y,chess,1,-1):winFlag = true
	}
	return winFlag
}

func checkLine(color int, x int, y int, chess[15][15]int, xBios int, yBios int)bool{
	counter := 1
	for{
		if chess[x+xBios][y+yBios] == color{
			counter++
		}else{
			break
		}
	}
	xBios = -xBios
	yBios = -yBios
	for{
		if chess[x+xBios][y+yBios] == color{
			counter++
		}else{
			break
		}
	}
	if counter >= 5{
		return true
	}else{
		return false
	}
}