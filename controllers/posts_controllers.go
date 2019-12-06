package controllers

import (
	"fmt"
	"github.com/omkz/golang-mysql/models"
	"github.com/omkz/golang-mysql/config"
)

func GetPosts() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from posts")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result models.Posts

	for rows.Next() {
		var each = models.Post{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.Title)
		fmt.Println(each.Content)
	}

}