package dao

import (
	"demo/modle"
	"fmt"
	"os"
)

func Recharge(m modle.Recharge) int {
	x, err := database.Query("select money from user where phone=? ", m.Phone)
	if err != nil {
		fmt.Println(err)
	}
	var money int
	for x.Next() {
		x.Scan(&money)
	}
	money = +m.Money
	_, err1 := database.Exec("update user set money=? where phone=?", money, m.Phone)
	if err1 != nil {
		fmt.Println(err1)
	}
	return money
}
func ChangeAddress(m modle.ChangeAddress) {
	_, err := database.Exec("update user set address=? where phone=?", m.Address, m.Phone)
	if err != nil {
		fmt.Println(err)
	}
}
func Change(m modle.Change, x string) {
	f, err1 := database.Query("select picture from user where phone=?", m.Phone)
	if err1 != nil {
		fmt.Println(err1)
	}
	var z string
	for f.Next() {
		f.Scan(&z)
	}
	if z != x {
		err1 = os.Remove("./" + z)
		if err1 != nil {
			fmt.Println(err1)
		}
	}
	_, err := database.Exec("update user set name=?,picture=?  where phone=?", m.Name, x, m.Phone)
	if err != nil {
		fmt.Println(err)
	}
}
