package service

import (
	"demo/dao"
	"demo/modle"
)

func Recharge(m modle.Recharge) int {
	x := dao.Recharge(m)
	return x
}
func ChangeAddress(m modle.ChangeAddress) {
	dao.ChangeAddress(m)
}
func Change(m modle.Change, x string) {
	dao.Change(m, x)
}
