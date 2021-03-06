package main

import (
	"fmt"
	"strconv"
)

// Goods 定义接口
type Goods interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name string
	quantity int
	price int
}

type Gift struct {
	name string
	quantity int
	price int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity)+ "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

func (gift Gift) settleAccount() int {
	return gift.quantity * gift.price
}

func (gift Gift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity)+ "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Goods) int {
	var allPrice int
	for _, good :=range goods{
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}


func main() {
	iphone := Phone{
		name: "Iphone",
		quantity: 1,
		price: 5999,
	}

	phoneGift := Gift{
		name: "Nokia",
		quantity: 1,
		price: 2999,
	}

	goods := []Goods{iphone, phoneGift}
	//fmt.Printf("%T", goods)
	totalPrice := calculateAllPrice(goods)
	fmt.Println("total price: ", totalPrice)

	// 空接口
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。

	// 定义一个空接口x
	var x interface{}
	s := "Hello World"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}