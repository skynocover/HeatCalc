package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Torder struct {
	order string
	orarr [2][]string
}

func (order *Torder) getOrder() {
	fmt.Println("==========")
	fmt.Println("What is your order?")
	fmt.Scan(&order.order)
}

func (orders *Torder) doOrder() {
	orders.splitOrder()
	for i := 0; i < len(orders.orarr[0]); i++ {
		//oarr[0][i]主命令
		check := 0 //確認輸入是否正確

		input := []*Parameter{&opera.Dm, &opera.Pcs, &opera.V, &opera.C0, &opera.B, &opera.Btype, &opera.Rpm, &opera.Fu, &opera.Fr, &opera.Lube}

		switch orders.orarr[0][i] {
		case "save":
			err := Save(orders.orarr[1][i]+".ht", &opera)
			Check(err)
			check = 1
		case "load":
			err := Load(orders.orarr[1][i]+".ht", &opera)
			Check(err)
			check = 1
		default:
			for j := 0; j < len(input); j++ {
				if input[j].coefficient == orders.orarr[0][i] {
					check = 1
					input[j].Value, _ = strconv.ParseFloat(orders.orarr[1][i], 64)
				}
			}
		}

		if check == 0 {
			fmt.Println(orders.orarr[0][i] + " 輸入錯誤")
		}
	}
}

func (orders *Torder) splitOrder() {
	osplit := strings.Split(orders.order, ",")
	for i := 0; i < len(osplit); i++ {
		a := strings.Split(osplit[i], "=")
		orders.orarr[0] = append(orders.orarr[0], a[0])
		orders.orarr[1] = append(orders.orarr[1], a[1])
	}
}

/*
func splito(order string) (oarr [2][]string) { //將命令拆開命回傳陣列
	osplit := strings.Split(order, ",")
	for i := 0; i < len(osplit); i++ {
		a := strings.Split(osplit[i], "=")
		oarr[0] = append(oarr[0], a[0])
		oarr[1] = append(oarr[1], a[1])
	}
	return
}
*/
