package service

import (
	"demo/dao"
	"demo/modle"
)

func Sell(good modle.Good, x string) {
	dao.Sell(good, x)
}
func AddCart(z modle.AddCart) modle.Cart {
	x := dao.AddCart(z)
	return x
}
func DeleteCart(m int) {
	dao.DeleteCart(m)
}
func ClearCart(m modle.ClearCart) bool {
	x := dao.ClearCart(m)
	return x
}
