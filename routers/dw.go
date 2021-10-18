package routers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"helin/config"
	"math/rand"
	"net"
	"net/http"
	"strconv"
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
	//s1 := client(DW1)
	//s2 := client(DW2)
	//dw1 := dealDw(s1)
	//dw2 := dealDw(s2)
	dw1 := dddd()
	dw2 := dddd()
	m := make(map[string]interface{})
	if dw1 == "" || dw2 == "" {
		m["code"] = -1
		m["message"] = "定位异常"
		m["data"] = ""
		return c.JSON(http.StatusOK, m)
	}
	dw := dw2 + "," + dw1
	m["code"] = 0
	m["message"] = "成功"
	m["data"] = dw
	return c.JSON(http.StatusOK, m)
}

func Uwb(c echo.Context) error {
	//s := client(UWB)
	s := test1()
	m := make(map[string]interface{})
	if s == "" {
		m["code"] = -1
		m["message"] = "定位异常"
		m["data"] = ""
		return c.JSON(http.StatusOK, m)
	}
	data := dealUwb(s)
	//data := make(map[string]interface{})
	//data["person"] = uwbData()
	//data["cart"] = uwbData()
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
		if strings.Contains(s[i], "nan") {
			continue
		}
		if strings.EqualFold(s[i], "0.000 0.000 0.000") {
			continue
		}
		if i < 11 {
			cart = append(cart, s[i])
		} else {
			person = append(person, s[i])
		}
	}
	m := make(map[string]interface{})
	m["person"] = person
	m["cart"] = cart
	return m
}

func test1() string {
	s1 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,134.991_29.018_14.125,127.629_38.033_18.787,104.305_34.191_12.861,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s2 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,137.625_29.031_12.949,129.427_38.115_10.353,106.825_33.954_11.827,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s3 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,140.956_29.047_13.821,138.279_37.572_12.629,117.243_32.617_11.887,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s4 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,142.879_-nan(ind)_-nan(ind),143.052_9.216_-nan(ind),124.547_31.832_10.969,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s5 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,165.318_28.818_11.751,107.243_37.987_14.145,151.445_29.423_12.700,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s6 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,171.418_28.858_11.382,118.393_37.207_14.305,164.302_30.344_11.030,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	s7 := "UWB,0.000_0.000_0.000,0.000_0.000_0.000,150.156_37.646_14.703,1.865_47.999_16.443,-1.730_47.846_19.388,0.000_0.000_0.000,0.000_0.000_0.000,6.141_-55.257_5.304,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,0.000_0.000_0.000,END"
	i := []string{s1, s2, s3, s4, s5, s6, s7}
	i2 := rand.Intn(6)
	return i[i2]
}

func uwbData() [5]string {
	//UWB,0.000_0.000_0.000,END  30个坐标，前10个是车后20个是人  长度300 ，宽度+-50 高度30
	var uwb [5]string
	for i := 0; i < 5; i++ {
		x := rand.Intn(300)
		y := rand.Intn(50)
		if i%2 != 0 {
			y = -y
		}
		z := rand.Intn(30)
		uwb[i] = strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z)
	}
	return uwb
}

func client(dw string) string {
	addr := config.GetConfig().GetString("ipPort")
	conn, err := net.DialTimeout(NETWORK, addr, 200*time.Millisecond) //创建套接字,连接服务器,设置超时时间
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
	//行进距离 0-30    回转-90-90  俯仰-10- 10
	x := rand.Intn(30)
	y := rand.Intn(90)
	if y%2 != 0 {
		y = -y
	}
	z := rand.Intn(10)
	if z%2 != 0 {
		z = -z
	}
	result := strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z)
	fmt.Println("采集的信息是：" + result)
	return result
}
