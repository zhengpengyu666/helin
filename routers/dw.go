package routers

import (
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
	println("已经访问" + dw + "号斗轮机\n")
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
