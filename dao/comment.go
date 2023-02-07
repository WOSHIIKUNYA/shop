package dao

import (
	"demo/modle"
	"fmt"
)

func SendComment(m modle.Comment) {
	_, err := database.Exec("insert into comment(name,score,goodid,body)values(?,?,?,?)", m.Name, m.Score, m.GoodsId, m.Body)
	if err != nil {
		fmt.Println(err)
	}
}
func LookComment(m modle.GoodId) []modle.CommentAll {
	x, err := database.Query("select id,name,score,goodid,body from comment where goodid=?", m.GoodsId)
	if err != nil {
		fmt.Println(err)
	}
	var z []modle.CommentAll
	for x.Next() {
		var a modle.CommentAll
		x.Scan(&a.Id, &a.Name, &a.Score, &a.GoodsId, &a.Body)
		z = append(z, a)
	}
	var score int = 0
	for b := 0; b < len(z); b++ {
		score += z[b].Score
	}
	score /= len(z)
	score = score + 1
	fmt.Println(score, m.GoodsId)
	database.Exec("update goods set score=? where id=? ", score, m.GoodsId)
	return z
}
func Delete(m modle.Writer) bool {
	x, err := database.Query("select name from comment where id=?", m.Id)
	if err != nil {
		fmt.Println(err)
	}
	var a string
	for x.Next() {
		x.Scan(&a)
	}
	if a != m.Name {
		return false
	}
	_, err1 := database.Exec("delete from comment where id=?", m.Id)
	if err1 != nil {
		fmt.Println(err)
	}
	return true
}
