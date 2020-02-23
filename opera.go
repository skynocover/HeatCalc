package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	f0 = [][]float64{{0.088, 0.28, 0.8}, {0.13, 0.46, 1}} //f0使用lube以及btype計算
)

type Parameter struct {
	coefficient string  //指令名
	name        string  //指令說明
	Value       float64 //數值 只有Value需要儲存
	unit        string  // 單位
}

type Topera struct {
	Tresult
	Tinput
}

type Tresult struct {
	F0   Parameter
	P0   Parameter
	F1   Parameter
	G1p0 Parameter
	Mv   Parameter
	Ml   Parameter
	M    Parameter
	Q    Parameter
}

type Tinput struct {
	Rpm   Parameter
	Fu    Parameter
	Fr    Parameter
	Dm    Parameter
	Pcs   Parameter
	V     Parameter
	C0    Parameter
	B     Parameter
	Lube  Parameter
	Btype Parameter
}

func prtable(input []Parameter) {
	//確認空格的數量
	var sup [3]int
	for m := 0; m < len(input); m++ {
		if strings.Count(input[m].name, "") > sup[0] {
			sup[0] = strings.Count(input[m].name, "")
		}
		if strings.Count(input[m].coefficient, "") > sup[1] {
			sup[1] = strings.Count(input[m].coefficient, "")
		}
		if strings.Count(strconv.FormatFloat(input[m].Value, 'f', -1, 32), "") > sup[2] {
			sup[2] = strings.Count(strconv.FormatFloat(input[m].Value, 'f', -1, 32), "")
		}
	}
	//利用空格數量去填空
	for i := 0; i < len(input); i++ {
		for j := 0; j < sup[0]-strings.Count(input[i].name, ""); j++ {
			fmt.Printf("  ")
		}
		fmt.Printf(input[i].name + ",")
		for k := 0; k < sup[1]-strings.Count(input[i].coefficient, ""); k++ {
			fmt.Printf(" ")
		}
		fmt.Printf(input[i].coefficient + ",")
		for m := 0; m < sup[2]-strings.Count(strconv.FormatFloat(input[i].Value, 'f', -1, 32), ""); m++ {
			fmt.Printf(" ")
		}
		fmt.Printf(strconv.FormatFloat(input[i].Value, 'f', -1, 32) + " ")
		fmt.Printf(input[i].unit)
		fmt.Println("")
	}
}
func (opera *Topera) prt() { //印出當前參數
	fmt.Println("軸承常數")
	fmt.Println("==========")
	prtable([]Parameter{opera.Dm, opera.Pcs, opera.V, opera.C0, opera.B, opera.Btype})

	fmt.Println("")
	fmt.Println("運轉常數")
	fmt.Println("==========")
	prtable([]Parameter{opera.Rpm, opera.Fu, opera.Fr, opera.Lube})

	fmt.Println("")
	fmt.Println("計算常數")
	fmt.Println("==========")
	prtable([]Parameter{opera.F0, opera.P0, opera.F1, opera.G1p0, opera.Mv, opera.Ml, opera.M, opera.Q})
}

func (opera *Topera) calc() { //計算方程式
	opera.P0.Value = math.Floor(opera.Fu.Value*9.81/math.Tan(opera.B.Value*math.Pi/180)*100000+0.5) / 100000 //靜等價賀重
	opera.F0.Value = f0[int(opera.Btype.Value)][int(opera.Lube.Value)]                                       //潤滑定數

	//軸承型式定數
	if opera.Btype.Value == 0 {
		opera.F1.Value = math.Floor(0.001*opera.Pcs.Value*math.Pow(opera.P0.Value/opera.C0.Value, 0.33)*1000000+0.5) / 10000000
	} else {
		opera.F1.Value = 0.0003
	}
	//荷重常數
	opera.G1p0.Value = math.Floor((0.9*opera.Fu.Value/math.Tan((opera.B.Value)*math.Pi/180)-0.1*opera.Fr.Value)*10000+0.5) / 10000
	if opera.G1p0.Value < opera.Fr.Value {
		opera.G1p0.Value = opera.Fr.Value
	}

	opera.Ml.Value = math.Floor(opera.F1.Value*opera.G1p0.Value*opera.Dm.Value*math.Pow(10, -3)*100000+0.5) / 100000
	opera.Mv.Value = math.Floor((opera.Pcs.Value*opera.F0.Value*math.Pow(opera.Dm.Value, 3)*math.Pow(opera.V.Value*opera.Rpm.Value, (0.6666666667))*math.Pow(10, -11))*100000+0.5) / 100000
	opera.M.Value = opera.Ml.Value + opera.Mv.Value
	opera.Q.Value = math.Floor(0.00234*math.Pi*opera.M.Value*opera.Rpm.Value*60*2*1000+0.5) / 1000
}
