package controllers

import (
	"fmt"
	"github.com/omkz/golang-mysql/models"
)

func GetPosts(){
	posts, err := models.PostAll()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range posts {
		fmt.Println(each.Title)
		fmt.Println(each.Content)
	}
}