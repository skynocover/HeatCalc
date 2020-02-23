package main

import (
	"os"
	"os/exec"
	"runtime"
)

var opera Topera

func init() {
	opera.Dm.coefficient = "dm"
	opera.Dm.name = "節圓直徑"
	opera.Dm.Value = 140
	opera.Dm.unit = "mm"

	opera.Pcs.coefficient = "pcs"
	opera.Pcs.name = "顆數"
	opera.Pcs.Value = 2
	opera.Pcs.unit = "pcs"

	opera.V.coefficient = "v"
	opera.V.name = "黏度"
	opera.V.Value = 25
	opera.V.unit = "cst"

	opera.Rpm.coefficient = "rpm"
	opera.Rpm.name = "轉速"
	opera.Rpm.Value = 600
	opera.Rpm.unit = "rpm"

	opera.C0.coefficient = "c0"
	opera.C0.name = "靜額定負荷"
	opera.C0.Value = 105000
	opera.C0.unit = "N"

	opera.Fu.coefficient = "fu"
	opera.Fu.name = "推力荷重"
	opera.Fu.Value = 8.6
	opera.Fu.unit = "kgf"

	opera.Fr.coefficient = "fr"
	opera.Fr.name = "徑向荷重"
	opera.Fr.Value = 6.8
	opera.Fr.unit = "kgf"

	opera.B.coefficient = "b"
	opera.B.name = "接觸角"
	opera.B.Value = 15
	opera.B.unit = "度"

	opera.Lube.coefficient = "lube"
	opera.Lube.name = "潤滑方式"
	opera.Lube.Value = 0
	opera.Lube.unit = "(0=oilair 1=grease 2=oiljet)"

	opera.Btype.coefficient = "btype"
	opera.Btype.name = "軸承型式"
	opera.Btype.Value = 0
	opera.Btype.unit = "(0=angular 1=roller)"

	opera.F0.coefficient = "f0"
	opera.F0.name = "軸承潤滑定數"
	opera.F0.unit = ""

	opera.P0.coefficient = "p0"
	opera.P0.name = "靜等價荷重"
	opera.P0.unit = "N"

	opera.F1.coefficient = "f1"
	opera.F1.name = "軸承型式定數"
	opera.F1.unit = ""

	opera.G1p0.coefficient = "g1p0"
	opera.G1p0.name = "荷重常數"
	opera.G1p0.unit = "kgf"

	opera.Mv.coefficient = "mv"
	opera.Mv.name = "速度項"
	opera.Mv.unit = "kgf*m"

	opera.Ml.coefficient = "ml"
	opera.Ml.name = "荷重項"
	opera.Ml.unit = "kgf*m"

	opera.M.coefficient = "m"
	opera.M.name = "動摩擦力矩"
	opera.M.unit = "kgf*m"

	opera.Q.coefficient = "Q"
	opera.Q.name = "發熱量"
	opera.Q.unit = "kcal/h"
}

func main() {
	opera.calc() //計算出結果
	opera.prt()  //印出結果
	var order Torder
	order.getOrder() //問問題
	CallClear()      //清空
	order.doOrder()  //依照命令判斷參數正確與否
	main()
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
	Value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		Value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
