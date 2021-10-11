package routers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"helin/config"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	NETWORK = "tcp"
	DW1     = "DATA-DW-1"
	DW2     = "DATA-DW-2"
	UWB     = "DATA-UWB"
)

func Dw(c echo.Context) error {
	s1 := client(DW1)
	s2 := client(DW2)
	dw1 := dealDw(s1)
	dw2 := dealDw(s2)
	//dw1 := dddd()
	//dw2 := dddd()
	m := make(map[string]interface{})
	if dw1 == "" || dw2 == "" {
		m["code"] = -1
		m["message"] = "定位异常"
		m["data"] = ""
		return c.JSON(http.StatusOK, m)
	}
	dw := dw1 + "," + dw2
	m["code"] = 0
	m["message"] = "成功"
	m["data"] = dw
	return c.JSON(http.StatusOK, m)
}

func Uwb(c echo.Context) error {

	//s := client(UWB)
	m := make(map[string]interface{})
	//if s == "" {
	//	m["code"] = -1
	//	m["message"] = "定位异常"
	//	m["data"] = ""
	//	return c.JSON(http.StatusOK, m)
	//}

	s := "UWB,200.000_22.000_10.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,150.000_13.000_20.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,0.000_0.000_0.000,300.000_50.000_30.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,25.000_-20.000_10.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000," +
		"0.000_0.000_0.000,END"
	data := dealUwb(s)
	m["code"] = 0
	m["message"] = "成功"
	m["data"] = data
	return c.JSON(http.StatusOK, m)
}
func dealUwb(data string) map[string]interface{} {
	//UWB,0.000_0.000_0.000,END  30个坐标，前10个是车后20个是人  长度300 ，宽度+-50 高度30
	data = strings.Replace(data, "_", " ", -1)
	s := strings.Split(data, ",")
	var person []string
	var cart []string
	for i := 1; i < len(s)-1; i++ {
		if strings.EqualFold(s[i], "0.000 0.000 0.000") {
			continue
		}
		if i < 11 {
			person = append(person, s[i])
		} else {
			cart = append(cart, s[i])
		}
	}
	m := make(map[string]interface{})
	m["person"] = person
	m["cart"] = cart
	return m
}

func client(dw string) string {
	addr := config.GetConfig().GetString("ipPort")
	conn, err := net.DialTimeout(NETWORK, addr, 100*time.Millisecond) //创建套接字,连接服务器,设置超时时间
	if err != nil {
		//fmt.Println(err)
		//os.Exit(1)
		return ""
	}
	conn.Write([]byte(dw)) //发送数据给服务器端
	var buff [512]byte
	read, err := conn.Read(buff[0:])
	s := string(buff[:read])
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			println(err)
		}
	}(conn)
	fmt.Println("采集的信息是：" + s)
	return s
}
func dealDw(data string) string {
	//	DW,-1000.000,-1000.000,-1000.000,1.100,2.200,3.300,END  说明：DW,行进距离,回转角度,俯仰角度,无用数据,无用数据,无用数据,END
	s := strings.Split(data, ",")
	dataArr := []string{s[1], s[2], s[3]}
	dataStr := strings.Join(dataArr, " ")
	return dataStr
}

func dddd() string {
	var dw1 [30]string
	dw1[0] = "0.000 -90.000 -10.000"
	dw1[1] = "0.000 -60.000 -0.000"
	dw1[2] = "0.000 -30.000 10.000"
	dw1[3] = "10.000 -90.000 -10.000"
	dw1[4] = "10.000 -90.000 0.000"
	dw1[5] = "10.000 -60.000 10.000"
	dw1[6] = "20.000 -30.000 -10.000"
	dw1[7] = "20.000 90.000 0.000"
	dw1[8] = "20.000 60.000 10.000"
	dw1[9] = "30.000 30.000 -10.000"
	dw1[10] = "30.000 -90.000 0.000"
	dw1[11] = "30.000 60.000 -10.000"
	i := rand.Intn(12)
	return dw1[i]
}
