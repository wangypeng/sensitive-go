package dao

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"sensitive-go/sensitive-go/model"
)	

type Dao struct{
	S *mgo.Session
}

func (dao *Dao) Query (limit int , skip int) []model.WordStruct{

	session := dao.S.Copy()
	defer session.Close()

	var p []model.WordStruct
	session.DB("test").C("people").Find(bson.M{}).Limit(limit).Skip(skip).All(&p)

	return p
}

func (dao *Dao) FindOne (queryWord string) (bool) {

	session := dao.S.Copy()
	defer session.Close()

	var p *model.WordStruct

	err := session.DB("test").C("people").Find(bson.M{"word":queryWord}).One(p)

	if err != nil {
		fmt.Println("find word to Db fail")
	}

	var resultData bool
	if p != nil || err != nil{ 
		resultData = false
	} else {
		resultData = true
	}
	return resultData;
}

func (dao *Dao) Insert (data *model.WordStruct) error{

	session := dao.S.Copy()
	defer session.Close()

	err := session.DB("test").C("people").Insert(data)

	if err != nil {
		fmt.Println("insert word to Db fail")
	}
	
	return err;
}

func (dao *Dao) Delete (word string) error {

	session := dao.S.Copy()
	defer session.Close()

	err := session.DB("test").C("people").Remove(bson.M{"word":word})

	if err != nil {
		fmt.Println("delete word to Db fail")
	}

	return err;
}
