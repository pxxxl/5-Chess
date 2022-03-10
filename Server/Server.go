package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var debugger = initDebugger()

func main() {
	chess := [15][15]int{}
	roundFlag := 0
	buf := [1024]byte{}
	listen, err1 := net.Listen("tcp", "127.0.0.1:8098")
	Handle(err1)
	conn1, err2 := listen.Accept()
	Handle(err2)
	fmt.Println("链接完成，链接信息如下：")
	fmt.Println("白方:"+conn1.RemoteAddr().String())
	_, err3 := conn1.Write([]byte("100"))
	Handle(err3)
	conn2, err4 := listen.Accept()
	Handle(err4)
	fmt.Println("链接完成，链接信息如下：")
	fmt.Println("黑方:"+conn2.RemoteAddr().String())
	_, err5 := conn2.Write([]byte("200"))
	Handle(err5)
	fmt.Println("游戏马上开始")
	time.Sleep(1*time.Second)
	roundFlag = 1
	sendMsg("000",conn1, conn2)
	//开始持续接收信息
	for{
		if roundFlag == 1{
			_, err6 := conn1.Read(buf[:])
			Handle(err6)
			chess[buf[1]-'a'][buf[2]-'A'] = 1
			debug(fmt.Sprintf("棋子颜色已设置，现在白方下 %c %c ",rune(buf[1]),rune(buf[2])))
			win := checkWin(1,int(buf[1]-'a'),int(buf[2]-'A'),chess)
			if win{
				sendMsg("3"+string(buf[1])+string(buf[2]),conn1,conn2)
			}else{
				sendMsg("1"+string(buf[1])+string(buf[2]),conn1,conn2)
			}
			debug("检查完成")
		}else{
			_, err6 := conn2.Read(buf[:])
			Handle(err6)
			chess[buf[1]-'a'][buf[2]-'A'] = -1
			debug(fmt.Sprintf("棋子颜色已设置，现在黑方下 %c %c ",rune(buf[1]),rune(buf[2])))
			win := checkWin(-1,int(buf[1]-'a'),int(buf[2]-'A'),chess)
			debug(fmt.Sprintf("棋子颜色已设置，现在黑方下 %c %c ",rune(buf[1]),rune(buf[2])))
			if win{
				sendMsg("4"+string(buf[1])+string(buf[2]),conn1,conn2)
			}else{
				sendMsg("2"+string(buf[1])+string(buf[2]),conn1,conn2)
			}
			debug("检查完成")
		}
		roundFlag = -roundFlag
		debug("更换双方")
	}
}

func sendMsg(str string,a net.Conn, b net.Conn){
	_, e1 := a.Write([]byte(str))
	_, e2 := b.Write([]byte(str))
	if e1!=nil{
		fmt.Printf("sendMsg failed, err:%v\n", e1)
		exitConsole()
	}
	if e2!=nil{
		fmt.Printf("sendMsg failed, err:%v\n", e2)
		exitConsole()
	}

}

func checkWin(color int,x int, y int, chess [15][15]int)bool{
	winFlag := false
	debug("winFlag set")
	switch{
	case checkLine(color,x,y,chess,1,0):winFlag = true
	case checkLine(color,x,y,chess,0,1):winFlag = true
	case checkLine(color,x,y,chess,1,1):winFlag = true
	case checkLine(color,x,y,chess,1,-1):winFlag = true
	}
	debug("check done")
	return winFlag
}

func checkLine(color int, x int, y int, chess[15][15]int, xBios int, yBios int)bool{
	counter := 1
	reach := 1
	for{
		if x+xBios*reach<0||x+xBios*reach>=15||y+yBios*reach<0||y+yBios*reach>=15{
			break
		}
		if chess[x+xBios*reach][y+yBios*reach] == color{
			reach++
			counter++
		}else{
			break
		}
	}
	debug("1checkLine done")
	reach = 1
	xBios = -xBios
	yBios = -yBios
	for{
		if x+xBios*reach<0||x+xBios*reach>=15||y+yBios*reach<0||y+yBios*reach>=15{
			break
		}
		if chess[x+xBios*reach][y+yBios*reach] == color{
			reach++
			counter++
		}else{
			break
		}
	}
	debug("2checkLine done")
	if counter >= 5{
		return true
	}else{
		return false
	}
}

func exitConsole(){
	var exit rune
	for{
		fmt.Scan(&exit)
		if exit == 'q'{
			os.Exit(1)
		}
	}


}

func initDebugger()net.Conn{
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	Handle(err)
	return conn
}

func debug(str string){
	_ ,err := debugger.Write([]byte(str))
	Handle(err)
}

func Handle(err error){
	if err != nil{
		fmt.Println(err)
		exitConsole()
	}
}