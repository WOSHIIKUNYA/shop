package dao

import (
	"demo/modle"
	"fmt"
)

func SearchMule(Name modle.SearchName) modle.Goods {
	x, err := database.Query("select *from goods ")
	if err != nil {
		fmt.Println(err)
	}
	var a []modle.Goods
	for x.Next() {
		var A modle.Goods
		x.Scan(&A.Id, &A.Name, &A.Kind, &A.Price, &A.Comment, &A.Picture, &A.Owner)
		a = append(a, A)
	}
	for b := 0; b < len(a); b++ {
		fmt.Println(a[b].Name)
		if Name.Name == a[b].Name {
			return a[b]
		}
	}
	var f modle.Goods
	return f
}
func SearchTag(tag modle.Tag) []modle.Goods {
	x, err := database.Query("select *from goods ")
	if err != nil {
		fmt.Println(err)
	}
	var a []modle.Goods
	for x.Next() {
		var A modle.Goods
		x.Scan(&A.Id, &A.Name, &A.Kind, &A.Price, &A.Comment, &A.Picture, &A.Owner)
		a = append(a, A)
	}
	var z []modle.Goods
	for b := 0; b < len(a); b++ {
		if tag.Kind == a[b].Kind {
			z = append(z, a[b])
		}
	}
	return z
}
func SearchPrice(tag modle.Tag) []modle.Goods {
	x, err := database.Query("select *from goods ")
	if err != nil {
		fmt.Println(err)
	}
	var a []modle.Goods
	for x.Next() {
		var A modle.Goods
		x.Scan(&A.Id, &A.Name, &A.Kind, &A.Price, &A.Comment, &A.Picture, &A.Owner)
		a = append(a, A)
	}
	var z []modle.Goods
	for b := 0; b < len(a); b++ {
		if tag.Kind == a[b].Kind {
			z = append(z, a[b])
		}
	}
	for b := 1; b < len(z); b++ {
		for c := 0; c+1 < len(z); c++ {
			if z[c].Price > z[c+1].Price {
				z[c], z[c+1] = z[c+1], z[c]
			}
		}
	}
	return z
}
func SearchPicture(id modle.Id) string {
	x, err := database.Query("select picture from goods where id=?", id.Id)
	if err != nil {
		fmt.Println(err)
	}
	var z string
	for x.Next() {
		x.Scan(&z)
	}
	return z
}
