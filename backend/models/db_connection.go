package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"github.com/joho/godotenv"
)

var Db *mgo.Database

//init関数(これはパッケージの呼び出し時にmain関数よりも先に呼び出され、違うパッケージでは、importすれば呼び出される)
func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil{
		//log.Fatalln(err)
		log.Println(err)
	}
	//fmt.Print("connection：%v\n", session)
	//defer session.Close()

	Db = session.DB(os.Getenv("MONGODB_DBNAME"))
	//fmt.Print("connectionDB：%v\n", Db)
	//fmt.Print("connectionCollection：%v\n", Db.C("people"))
}
