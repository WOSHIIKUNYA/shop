package service

import (
	"demo/dao"
	"demo/modle"
)

func SendComment(m modle.Comment) {
	dao.SendComment(m)
}
func LookComment(m modle.GoodId) []modle.CommentAll {
	x := dao.LookComment(m)
	return x
}
func Delete(m modle.Writer) bool {
	x := dao.Delete(m)
	return x
}
