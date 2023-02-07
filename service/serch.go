package service

import (
	"demo/dao"
	"demo/modle"
)

func SearchMule(name modle.SearchName) modle.Goods {
	a := dao.SearchMule(name)
	return a
}
func SearchTage(name modle.Tag) []modle.Goods {
	a := dao.SearchTag(name)
	return a
}
func SearchPicture(id modle.Id) string {
	a := dao.SearchPicture(id)
	return a
}
func SearchPrice(name modle.Tag) []modle.Goods {
	a := dao.SearchPrice(name)
	return a
}
