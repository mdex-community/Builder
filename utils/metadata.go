package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os/user"
	"time"
)

func Metadata() {
	//Metedata
	timestamp := time.Now().Format(time.RFC850)
	ip := GetIPAdress().String()
	userName := GetUserData().Username
	homeDir := GetUserData().HomeDir

	//Contains a collection of fileds with user's metadata
	userMetaData := AllMetaData{
		UserName:  userName,
		HomeDir:   homeDir,
		IP:        ip,
		Timestamp: timestamp}

	OutputJSONall(&userMetaData)
}

//AllMetaData holds the stuct of all the arguments
type AllMetaData struct {
	UserName  string
	HomeDir   string
	IP        string
	Timestamp string
}

//GetUserData return username and userdir
func GetUserData() *user.User {
	user, err := user.Current()
	if err != nil {
		panic(err)

	}

	return user

}

// GetIPAdress Get preferred outbound ip of this machine
func GetIPAdress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

//OutputJSONall  outputs allMetaData struct in JSON format
func OutputJSONall(allData *AllMetaData) {
	data, _ := json.Marshal(allData)
	err := ioutil.WriteFile("./all.json", data, 0644)

	if err != nil {
		panic(err)
	}
}