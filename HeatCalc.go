package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	bearing = map[string]float64{} //計算用參數

	para    = []string{"dm", "pcs", "v", "rpm", "c0", "fu", "fr", "b"}
	parades = []string{"節圓直徑", "顆數", "黏度", "轉速", "靜額定負荷", "推力荷重", "徑向荷重", "接觸角"}

	result    = []string{"p0", "f1", "g1p0", "mv", "ml", "m", "Q"}
	resultdes = []string{"靜等價荷重", "軸承定數", "荷重常數", "速度項", "荷重項", "動摩擦力矩", "發熱量"}

	state    = map[string]string{} //狀態用參數 軸承樣式,潤滑方式等
	lubearr  = []string{"oilair", "grease", "oiljet"}
	lube     = 0
	btypearr = []string{"angular", "roller", "dgroove"}
	btype    = 0
	f0       = [][]float64{{0.088, 0.28, 0.8}, {0.13, 0.46, 1}} //f0使用lube以及btype計算
)

func init() {
	bearing["dm"] = 90     //節圓直徑
	bearing["pcs"] = 4     //軸承顆數
	bearing["v"] = 68      //黏度
	bearing["rpm"] = 28000 //轉速
	bearing["c0"] = 46000  //徑額定負荷
	bearing["fu"] = 0.5    //推力賀重
	bearing["fr"] = 5      //徑向賀重
	bearing["b"] = 15      //接觸角	                                                                

	state["lube"] = lubearr[lube]    //潤滑方式
	state["btype"] = btypearr[btype] //軸承型式
}

func main() {
	calc() //計算出結果並丟進bearing map裏面
	prt()  //印出結果

	order := getorder() //問問題
	CallClear()         //清空
	dorder(order)       //依照命令判斷參數正確與否並丟進bearing 
	//fmt.Println(bearing["f0"])

	main()
	//show()
}

func dorder(order string) {
	oarr := splito(order)
	for i := 0; i < len(oarr[0]); i++ {
		//fmt.Println(oarr[0][i])	//主命令
		check := 0
		if oarr[0][i] == "btype" {
			check = 1
			btype, _ = strconv.Atoi(oarr[1][i])
		} else if oarr[0][i] == "lube" {
			check = 1
			lube, _ = strconv.Atoi(oarr[1][i])
		}

		for j := 0; j < len(para); j++ {
			if para[j] == oarr[0][i] {
				check = 1
				bearing[oarr[0][i]], _ = strconv.ParseFloat(oarr[1][i], 64)
			}
		}
		if check == 0 {
			fmt.Println(oarr[0][i] + " 輸入錯誤")
		}
	}
	//fmt.Println(oarr)
}
func splito(order string) [2][]string {
	osplit := strings.Split(order, ",")
	var oarr [2][]string
	for i := 0; i < len(osplit); i++ {
		a := strings.Split(osplit[i], "=")
		//fmt.Println(a)
		oarr[0] = append(oarr[0], a[0])
		oarr[1] = append(oarr[1], a[1])
	}
	return oarr
}

func getorder() string { //問問題
	fmt.Println("==========")
	fmt.Println("What is your order?")
	var order string

	fmt.Scan(&order)
	return order
}

func prt() { //印出當前參數
	fmt.Println("基本常數")
	fmt.Println("==========")
	var valarr []string
	for i := 0; i < len(parades); i++ {
		valarr = append(valarr, strconv.FormatFloat(bearing[para[i]], 'f', -1, 32))
	}
	prtable(parades, para, valarr)

	fmt.Println("")
	fmt.Println("選擇常數")
	fmt.Println("==========")
	fmt.Println("潤滑方式, lube," + lubearr[lube] + "  (0=oilair  1=grease 2=oiljet)")
	fmt.Println("軸承型式,btype," + btypearr[btype] + "  (0=angular 1=roller 2=dgroove)")

	fmt.Println("")
	fmt.Println("計算常數")
	fmt.Println("==========")
	var rutarr []string
	for i := 0; i < len(result); i++ {
		rutarr = append(rutarr, strconv.FormatFloat(bearing[result[i]], 'f', -1, 32))
	}
	prtable(resultdes, result, rutarr)
}

func calc() { //計算方程式
	bearing["p0"] = math.Floor(bearing["fu"]*9.81/math.Tan(bearing["b"]*math.Pi/180)*1000+0.5) / 1000 //靜等價賀重
	bearing["f0"] = f0[btype][lube]   //軸承與潤滑定數
	bearing["f1"] = math.Floor(0.001*bearing["pcs"]*math.Pow(bearing["p0"]/bearing["c0"], 0.33)*10000000+0.5) / 10000000
	
	bearing["g1p0"] = math.Floor((0.9*bearing["fu"]/math.Tan(bearing["b"])-0.1*bearing["fr"])*100+0.5) / 100
	//fmt.Println(bearing["g1p0"])
	if bearing["g1p0"] < bearing["fr"] {
		bearing["g1p0"] = bearing["fr"]
	}

	bearing["ml"] = math.Floor(bearing["f1"]*bearing["g1p0"]*bearing["dm"]*math.Pow(10, -3)*1000000+0.5) / 1000000
	bearing["mv"] = math.Floor((bearing["pcs"]*bearing["f0"]*math.Pow(bearing["dm"], 3)*math.Pow(bearing["v"]*bearing["rpm"], (0.6666666667))*math.Pow(10, -11))*100000+0.5) / 100000
	bearing["m"] = bearing["ml"] + bearing["mv"]
	bearing["Q"] = math.Floor(0.00234*math.Pi*bearing["m"]*bearing["rpm"]*60*2*1000+0.5) / 1000
}

//tool------------------------------------------------------------------------
func prtable(a []string, b []string, c []string) { //印出表格
	alen, blen, clen := 0, 0, 0

	for i := 0; i < len(a); i++ {
		if strings.Count(a[i], "")-1 > alen {
			alen = strings.Count(a[i], "") - 1
		}
		if strings.Count(b[i], "")-1 > blen {
			blen = strings.Count(b[i], "") - 1
		}
		if strings.Count(c[i], "")-1 > clen {
			clen = strings.Count(c[i], "") - 1
		}

	}
	for i := 0; i < len(a); i++ {
		len := alen - strings.Count(a[i], "") + 1
		for j := 0; j < len; j++ {
			fmt.Printf("  ")
		}
		fmt.Printf(a[i] + ",")
		len = blen - strings.Count(b[i], "") + 1
		for k := 0; k < len; k++ {
			fmt.Printf(" ")
		}
		fmt.Printf(b[i] + ",")
		len = clen - strings.Count(c[i], "") + 1
		for k := 0; k < len; k++ {
			fmt.Printf(" ")
		}
		fmt.Println(c[i] + ",")
	}
}

func show() {
	for k, v := range bearing {
		fmt.Println(k, v)
	}
}

func CallClear() { //清空用
	var clear map[string]func()     //create a map for storing clear funcs
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
