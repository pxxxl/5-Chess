package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var colorFlag int
	var ipAndPort string
	ipAndPort = "10.19.190.109:8098"
	chess := [15][15]int{}
	var buf [1024]byte
	conn, err1 := net.Dial("tcp", ipAndPort)
	if err1 != nil {
		fmt.Printf("conn server failed, err:%v\n", err1)
		return
	}
	_, err2 :=conn.Read(buf[:])
	if err2 != nil {
		fmt.Printf("read failed, err:%v\n", err2)
		return
	}
	if buf[0] == '2'{
		colorFlag = -1
	}else{
		colorFlag = 1
	}
	//基本的东西设置好了，开始下了
	for{
		_, err3 :=conn.Read(buf[:])
		clearScreen()
		printChess(chess)
		if err3 != nil {
			fmt.Printf("read failed, err:%v\n", err2)
			return
		}
		switch buf[0]{
		case '0':
			if colorFlag == 1{
				ret := getCommand(1,chess)
				_, err4 :=conn.Write([]byte("1"+ret))
				if err4 != nil {
					fmt.Printf("send failed, err:%v\n", err4)
					return
				}
				waitingInformation()
			}else{
				waitingInformation()
			}
		case '1':
			chess[buf[1]-'a'][buf[2]-'A'] = 1
			if colorFlag == -1{
				ret := getCommand(-1,chess)
				_, err4 :=conn.Write([]byte("2"+ret))
				if err4 != nil {
					fmt.Printf("send failed, err:%v\n", err4)
					return
				}
				waitingInformation()
			}else{
				waitingInformation()
			}
		case '2':
			chess[buf[1]-'a'][buf[2]-'A'] = -1
			if colorFlag == 1{
				ret := getCommand(1,chess)
				_, err4 :=conn.Write([]byte("1"+ret))
				if err4 != nil {
					fmt.Printf("send failed, err:%v\n", err4)
					return
				}
				waitingInformation()
			}else{
				waitingInformation()
			}
		case '3':
			chess[buf[1]-'a'][buf[2]-'A'] = 1
			clearScreen()
			printChess(chess)
			fmt.Println("白方胜！")
			os.Exit(0)
		case '4':
			chess[buf[1]-'a'][buf[2]-'A'] = -1
			clearScreen()
			printChess(chess)
			fmt.Println("黑方胜！")
			os.Exit(0)
		}
	}
}

func printChess(chess [15][15]int){
	var str string
	str = "  A   B   C   D   E   F   G   H   I   J   K   L   M   N   O"
	fmt.Println(str)
	str = ""
	for i := 0;i <= 28;i++{
		if i % 2 ==0{
			str += string('a'+ i/2)
		}else{
			str += " "
		}
		str += chessLine(i, chess[i/2])
		fmt.Println(str)
		str = ""
	}

}

func chessLine(line int, chess [15]int) (str string){
	switch{
	case line == 0:
		if chess[0] == 0{
			str+="┌─"
		}else if chess[0] == 1{
			str+="●"
		}else{
			str+="○"
		}
		for i := 1;i < 28;i++{
			switch{
			case i%2==0:
				if chess[i/2]%2 == 0{
					str+="─┬─"
				}else if chess[i/2]%2 == 1{
					str+="─●"
				}else{
					str+="─○"
				}
			case i%2==1:
				str += "─"
			}
		}
		if chess[14] == 0{
			str+="─┐"
		}else if chess[14] == 1{
			str+="─●"
		}else{
			str+="─○"
		}
	case line == 28:
		if chess[0] == 0{
			str+="└─"
		}else if chess[0] == 1{
			str+="●"
		}else{
			str+="○"
		}
		for i := 1;i < 28;i++{
			switch{
			case i%2==0:
				if chess[i/2]%2 == 0{
					str+="─┴─"
				}else if chess[i/2]%2 == 1{
					str+="─●"
				}else{
					str+="─○"
				}
			case i%2==1:
				str += "─"
			}
		}
		if chess[14] == 0{
			str+="─┘"
		}else if chess[14] == 1{
			str+="─●"
		}else{
			str+="─○"
		}
	case line%2 == 0:
		if chess[0] == 0{
			str+="├─"
		}else if chess[0] == 1{
			str+="●"
		}else{
			str+="○"
		}
		for i := 1;i < 28;i++{
			switch{
			case i%2==0:
				if chess[i/2]%2 == 0{
					str+="─┼─"
				}else if chess[i/2]%2 == 1{
					str+="─●"
				}else{
					str+="─○"
				}
			case i%2==1:
				str += "─"
			}
		}
		if chess[14] == 0{
			str+="─┤"
		}else if chess[14] == 1{
			str+="─●"
		}else{
			str+="─○"
		}
	case line%2 == 1:
		str = "│   │   │   │   │   │   │   │   │   │   │   │   │   │   │"
	}
	return
}

func getCommand(color int, chess [15][15]int)string{
	var item rune
	if color == 1{
		item = '●'
	}else{
		item = '○'
	}
	fmt.Printf("请落子,%c", item)
	var str string
	for{
		fmt.Scanln(&str)
		if len(str) != 2{
			continue
		}
		if str[0]<='o'&&str[0]>='a'{
			if str[1]<='O'&&str[1]>='A'{
				if chess[str[0]-'a'][str[1]-'A'] == 0{
					return str
				}
			}
		}
		fmt.Println("输入无效,请重新输入")
	}
}

func clearScreen(){
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n")
}

func waitingInformation(){
	fmt.Println("请等待。。。")
}