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

type parameter struct {
	coefficient string
	name        string
	value       float64
	unit        string
}

type topera struct {
	tresult
	tinput
}

type tresult struct {
	f0   parameter
	p0   parameter
	f1   parameter
	g1p0 parameter
	mv   parameter
	ml   parameter
	m    parameter
	q    parameter
}

type tinput struct {
	rpm   parameter
	fu    parameter
	fr    parameter
	dm    parameter
	pcs   parameter
	v     parameter
	c0    parameter
	b     parameter
	lube  parameter
	btype parameter
}

func (opera *topera) Getinput() []*parameter {
	return []*parameter{&opera.dm, &opera.pcs, &opera.v, &opera.c0, &opera.b, &opera.btype, &opera.rpm, &opera.fu, &opera.fr, &opera.lube}
}

func (opera *topera) Getoutput() []*parameter {
	return []*parameter{&opera.f0, &opera.p0, &opera.f1, &opera.g1p0, &opera.mv, &opera.ml, &opera.m, &opera.q}
}

func (opera *topera) Getbearing() []*parameter {
	return []*parameter{&opera.dm, &opera.pcs, &opera.v, &opera.c0, &opera.b, &opera.btype}
}

func (opera *topera) Getopera() []*parameter {
	return []*parameter{&opera.rpm, &opera.fu, &opera.fr, &opera.lube}
}

func findsup(para []*parameter) (sup [3]int) {
	for m := 0; m < len(para); m++ {
		if strings.Count(para[m].name, "") > sup[0] {
			sup[0]=strings.Count(para[m].name, "")
		}
		if strings.Count(para[m].coefficient, "") > sup[1] {
			sup[1]=strings.Count(para[m].coefficient, "")
		}
		if strings.Count(strconv.FormatFloat(para[m].value, 'f', -1, 32), "") > sup[2] {
			sup[2]=strings.Count(strconv.FormatFloat(para[m].value, 'f', -1, 32), "")
		}
	}
	return
}

func prtable(input []*parameter) {
	sup := findsup(input)
	for i := 0; i < len(input); i++ {
		for j := 0; j < sup[0]-strings.Count(input[i].name, ""); j++ {
			fmt.Printf("  ")
		}
		fmt.Printf(input[i].name + ",")
		for k := 0; k < sup[1]-strings.Count(input[i].coefficient, ""); k++ {
			fmt.Printf(" ")
		}
		fmt.Printf(input[i].coefficient + ",")
		for m := 0; m < sup[2]-strings.Count(strconv.FormatFloat(input[i].value, 'f', -1, 32), ""); m++ {
			fmt.Printf(" ")
		}
		fmt.Printf(strconv.FormatFloat(input[i].value, 'f', -1, 32) + " ")
		fmt.Printf(input[i].unit)
		fmt.Println("")
	}
}
func (opera *topera) prt() { //印出當前參數

	fmt.Println("軸承常數")
	fmt.Println("==========")
	bearings := opera.Getbearing()
	prtable(bearings)

	fmt.Println("")
	fmt.Println("運轉常數")
	fmt.Println("==========")
	operas := opera.Getopera()
	prtable(operas)

	fmt.Println("")
	fmt.Println("計算常數")
	fmt.Println("==========")
	output := opera.Getoutput()
	prtable(output)

}
func (opera *topera) GetP0() float64 { //靜等價賀重
	return math.Floor(opera.fu.value*9.81/math.Tan(opera.b.value*math.Pi/180)*100000+0.5) / 100000 //靜等價賀重
}
func (opera *topera) Getf0() float64 { //潤滑定數
	return f0[int(opera.btype.value)][int(opera.lube.value)]
}
func (opera *topera) Getf1() float64 {
	return math.Floor(0.001*opera.pcs.value*math.Pow(opera.p0.value/opera.c0.value, 0.33)*1000000+0.5) / 10000000
}
func (opera *topera) Getg1p0() float64 {
	return math.Floor((0.9*opera.fu.value/math.Tan((opera.b.value)*math.Pi/180)-0.1*opera.fr.value)*10000+0.5) / 10000
}
func (opera *topera) Getml() float64 {
	return math.Floor(opera.f1.value*opera.g1p0.value*opera.dm.value*math.Pow(10, -3)*100000+0.5) / 100000
}
func (opera *topera) Getmv() float64 {
	return math.Floor((opera.pcs.value*opera.f0.value*math.Pow(opera.dm.value, 3)*math.Pow(opera.v.value*opera.rpm.value, (0.6666666667))*math.Pow(10, -11))*100000+0.5) / 100000
}
func (opera *topera) Getq() float64 {
	return math.Floor(0.00234*math.Pi*opera.m.value*opera.rpm.value*60*2*1000+0.5) / 1000
}
func (opera *topera) calc() { //計算方程式
	opera.p0.value = opera.GetP0() //靜等價賀重
	opera.f0.value = opera.Getf0() //潤滑定數

	if opera.btype.value == 0 { //軸承型式定數
		opera.f1.value = opera.Getf1()
	} else {
		opera.f1.value = 0.0003
	}

	opera.g1p0.value = opera.Getg1p0()
	if opera.g1p0.value < opera.fr.value {
		opera.g1p0.value = opera.fr.value
	}

	opera.ml.value = opera.Getml()
	opera.mv.value = opera.Getmv()
	opera.m.value = opera.ml.value + opera.mv.value
	opera.q.value = opera.Getq()
}
