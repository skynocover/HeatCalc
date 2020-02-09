package main

import (
	"fmt"
	"github.com/skynocover/gocmdtool"
	"strconv"
	"strings"
)

type Torder struct {
	order string
	orarr []string
}

func (order *Torder) Getorder() {
	fmt.Println("==========")
	fmt.Println("What is your order?")
	fmt.Scan(&order.order)
	cmd.CallClear() //清空
}

func (orders *Torder) Dorder(opera *topera) {
	oarr := splito(orders.order)
	for i := 0; i < len(oarr[0]); i++ {
		//oarr[0][i]主命令
		check := 0
		arr := opera.Getinput()

		for j := 0; j < len(arr); j++ {
			if arr[j].coefficient == oarr[0][i] {
				check = 1
				arr[j].value, _ = strconv.ParseFloat(oarr[1][i], 64)
			}
		}

		if check == 0 {
			fmt.Println(oarr[0][i] + " 輸入錯誤")
		}
	}
}

func splito(order string) (oarr [2][]string) {
	osplit := strings.Split(order, ",")
	//var oarr [2][]string
	for i := 0; i < len(osplit); i++ {
		a := strings.Split(osplit[i], "=")
		oarr[0] = append(oarr[0], a[0])
		oarr[1] = append(oarr[1], a[1])
	}
	return
}
