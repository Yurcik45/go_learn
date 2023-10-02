package controllers

import (
	"fmt"
)

type Item struct {
	id int8
	name string
	age int8
}

func AddItem(data Item) {}

func GetItem(id int8) Item {}

func ChangeItem(data Item) {}

func DeleteItem(id int8) {}