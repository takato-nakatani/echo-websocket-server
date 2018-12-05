package models

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

type Chat struct {
	ID        bson.ObjectId   `bson:"_id"`
	Name      string          `bson:"name"`
	Contents  string             `bson:"contents"`
}

func SaveChat(contents string) bool {
	chat := &Chat{
		ID:        bson.NewObjectId(),
		Name:      "ゲスト",
		Contents:  contents,
	}
	fmt.Print("chat：%v\n", Db)
	col := Db.C("chat")
	fmt.Print("col：%v\n", col)
	fmt.Print("contents：%v\n", contents)
	fmt.Print("struct：%v\n", chat)
	err := col.Insert(&chat)
	if err != nil {
		log.Fatalln(err)
	}
	return true
}