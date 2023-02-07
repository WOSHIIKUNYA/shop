package dao

import (
	"demo/modle"
	"fmt"
	"os"
)

func Sell(good modle.Good, x string) {
	_, err := database.Exec("insert into goods(name,kind,comment,price,owner,picture)values(?,?,?,?,?,?)", good.Name, good.Kind, good.Comment, good.Price, good.Owner, x)
	if err != nil {
		fmt.Println(err)
	}
}
func AddCart(m modle.AddCart) modle.Cart {
	x, err := database.Query("select id,picture,price,name from goods where id=?", m.Id)
	if err != nil {
		fmt.Println(err)
	}
	var c modle.Cart
	for x.Next() {
		x.Scan(&c.Id, &c.Picture, &c.Price, &c.Name)
	}
	c.Phone = m.Phone
	_, err1 := database.Exec("insert into cart (shopid,name,picture,price,phone) values (?,?,?,?,?)", c.Id, c.Name, c.Picture, c.Price, c.Phone)
	if err1 != nil {
		fmt.Println(err1)
	}
	return c
}
func DeleteCart(m int) {
	_, err := database.Exec("delete from cart where id=?", m)
	if err != nil {
		fmt.Println(err)
	}
}
func ClearCart(m modle.ClearCart) bool {
	x, err := database.Query("select picture,price,owner from goods where id=?", m.Id)
	if err != nil {
		fmt.Println(err)
	}
	var z, owner string
	var price int
	for x.Next() {
		x.Scan(&z, &price, &owner)
	}
	p := database.QueryRow("select money from  user where phone=?", m.Phone)
	var number int
	p.Scan(&number)
	number -= price
	if number < 0 {
		return false
	}
	_, err1 := database.Exec("update user set money=? where phone=?", number, m.Phone)
	if err1 != nil {
		fmt.Println(err1)
	}
	q := database.QueryRow("select money from user where phone=?", owner)
	q.Scan(&number)
	number += price
	_, err2 := database.Exec("update user set money=? where phone=?", number, owner)
	if err2 != nil {
		fmt.Println(err2)
	}
	os.Remove("./" + z)
	_, err = database.Exec("delete from goods where id=?", m.Id)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
