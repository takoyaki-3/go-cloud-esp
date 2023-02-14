package gocloudesp

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func SetPinMode()error{
	return nil
}

func WriteDigital()error{
	return nil
}

func WriteAnalog()error{
	return nil
}

func WriteServo()error{
	return nil
}

func ReadDigital()(string,error){
	return "",nil
}

func ReadAnalog()(int,error){
	// 0-4095
	return -1,nil
}

func Sleep()error{
	return nil
}

func Readtemperature()(int,error){
	return -1,nil
}

type Config struct {
	Host string
}

// HTTPアクセス
func access(conf Config) {
  url := "http://" + conf.Host
  req, _ := http.NewRequest("GET", url, nil)

  client := new(http.Client)
  resp, _ := client.Do(req)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
}
