package utils

import (
	"net"
	"net/http"
	"io/ioutil"
	"github.com/kdchain/go-kdchain/log"
	"fmt"
	"strings"
)

func GetExternalIP() string{
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		internalIP := GetInternalIP()
		log.Error("Can't get external IP, only get the internal IP!", "GetExternalIP() http.Get", internalIP)
		return internalIP
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	/*if err != nil {
		log.Info("Get external IP", "GetExternalIP", string(body))
		return  string(body)
	}

	internalIP := GetInternalIP()
	log.Error("Can't get external IP, only get the internal IP!", "GetExternalIP() ioutil.ReadAll", internalIP)
	return internalIP*/

	return string(body)
}

func  GetInternalIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return "127.0.0.1"
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func GetInternalIPNew() string{
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		return "127.0.0.1"
	}
	log.Info("Get internal IP", "GetInternalIP", addrSlice)
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return  "127.0.0.1"
}
