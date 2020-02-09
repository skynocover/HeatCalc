package main

var (
	opera topera
)

func init() {
	opera.dm.coefficient = "dm"
	opera.dm.name = "節圓直徑"
	opera.dm.value = 140
	opera.dm.unit = "mm"

	opera.pcs.coefficient = "pcs"
	opera.pcs.name = "顆數"
	opera.pcs.value = 2
	opera.pcs.unit = "pcs"

	opera.v.coefficient = "v"
	opera.v.name = "黏度"
	opera.v.value = 25
	opera.v.unit = "cSt"

	opera.rpm.coefficient = "rpm"
	opera.rpm.name = "轉速"
	opera.rpm.value = 600
	opera.rpm.unit = "rpm"

	opera.c0.coefficient = "c0"
	opera.c0.name = "靜額定負荷"
	opera.c0.value = 105000
	opera.c0.unit = "N"

	opera.fu.coefficient = "fu"
	opera.fu.name = "推力荷重"
	opera.fu.value = 8.6
	opera.fu.unit = "kgf"

	opera.fr.coefficient = "fr"
	opera.fr.name = "徑向荷重"
	opera.fr.value = 6.8
	opera.fr.unit = "kgf"

	opera.b.coefficient = "b"
	opera.b.name = "接觸角"
	opera.b.value = 15
	opera.b.unit = "度"

	opera.lube.coefficient = "lube"
	opera.lube.name = "潤滑方式"
	opera.lube.value = 0
	opera.lube.unit = "(0=oilair 1=grease 2=oiljet)"

	opera.btype.coefficient = "btype"
	opera.btype.name = "軸承型式"
	opera.btype.value = 0
	opera.btype.unit = "(0=angular 1=roller)"

	opera.f0.coefficient = "f0"
	opera.f0.name = "軸承潤滑定數"
	opera.f0.unit = ""

	opera.p0.coefficient = "p0"
	opera.p0.name = "靜等價荷重"
	opera.p0.unit = "N"

	opera.f1.coefficient = "f1"
	opera.f1.name = "軸承型式定數"
	opera.f1.unit = ""

	opera.g1p0.coefficient = "g1p0"
	opera.g1p0.name = "荷重常數"
	opera.g1p0.unit = "kgf"

	opera.mv.coefficient = "mv"
	opera.mv.name = "速度項"
	opera.mv.unit = "kgf*m"

	opera.ml.coefficient = "ml"
	opera.ml.name = "荷重項"
	opera.ml.unit = "kgf*m"

	opera.m.coefficient = "m"
	opera.m.name = "動摩擦力矩"
	opera.m.unit = "kgf*m"

	opera.q.coefficient = "Q"
	opera.q.name = "發熱量"
	opera.q.unit = "kcal/h"
}

func main() {
	opera.calc() //計算出結果
	opera.prt()  //印出結果
	var order Torder
	order.Getorder()     //問問題
	order.Dorder(&opera) //依照命令判斷參數正確與否
	main()
}

/*
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
*/
