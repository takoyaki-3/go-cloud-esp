package gocloudesp

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	json "github.com/takoyaki-3/go-json"
)

type ESP struct {
	Host string	`json:"host"`
	Port int 		`json:"port"`
	ID string `json:"id"`
	Conn net.Conn
}

func NewESP(path string)(*ESP,error){

	e := ESP{}

	// Load JSON config file
	err := json.LoadFromPath(path, &e)

	// Connect
	e.Conn,err = net.Dial("tcp",e.Host+":"+strconv.Itoa(e.Port))
	if err!=nil{
		return nil, err
	}

	// Send first message
	_,err = e.Conn.Write([]byte("control"))
	if err!=nil{
		return nil, err
	}

	b := make([]byte,1024)
	n,err := e.Conn.Read(b)
	if err!=nil{
		return nil,err
	}
	fmt.Println(string(b[:n]))
	
	return &e,nil
}

// ピンモードの設定
func (e *ESP)SetPinMode(pin int, mode string)error{
	// コマンド送信
	fmt.Println("setPinMode,"+strconv.Itoa(pin)+","+mode)
	e.Conn.Write([]byte("setPinMode,"+strconv.Itoa(pin)+","+mode))

	// 結果受信
	b := make([]byte,1024)
	n,err := e.Conn.Read(b)
	if err!=nil{
		return err
	}
	fmt.Println(string(b[:n]))
	return nil
}

// デジタル出力
func (e *ESP)WriteDigital(pin int, mode string)error{
	// コマンド送信
	e.Conn.Write([]byte("writeDigital,"+strconv.Itoa(pin)+","+mode))

	// 結果受信
	b := make([]byte,1024)
	n,err := e.Conn.Read(b)
	if err!=nil{
		return err
	}
	fmt.Println(string(b[:n]))
	return nil
}

// アナログ出力
func (*ESP)WriteAnalog()error{
	return nil
}

// サーボモータ出力
func (*ESP)WriteServo()error{
	return nil
}

// デジタル入力 HIGH 又は LOW
func (*ESP)ReadDigital()(string,error){
	return "",nil
}

// アナログ入力
func (e *ESP)ReadAnalog(pin int)(int,error){
	//readAnalog
	// 0-4095

	// コマンド送信
	e.Conn.Write([]byte("readAnalog,"+strconv.Itoa(pin)+","))

	// 結果受信
	b := make([]byte,1024)
	n,err := e.Conn.Read(b)
	if err!=nil{
		return -1,err
	}
	arr := strings.Split(string(b[:n]),",")
	n, err = strconv.Atoi(arr[0])

	return n,err
}

// 待機
func (*ESP)Sleep(int)error{
	return nil
}

// センサー内部温度の取得
func (*ESP)Readtemperature()(int,error){
	return -1,nil
}

