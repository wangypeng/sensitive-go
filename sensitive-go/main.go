package main

import (
    "fmt"
	"log"
	
	"net/http"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	
    "goSensitive/sensitivefilter/indexController"
    "goSensitive/sensitivefilter/model"
    "goSensitive/sensitivefilter/wordFilter"
    "goSensitive/sensitivefilter/conf"
    "goSensitive/sensitivefilter/constant"
)

var session *mgo.Session

func init (){

    /* init config file */
    conf.InitConf()

    /* init log config */
    // FIXME:log init

    /* init sql session */
    var err error
    session, err = mgo.Dial(conf.ConfigMap["db.ip"]+":"+conf.ConfigMap["db.port"])
    if err != nil {
        panic(err)
    }

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    /* init DFA tree */
    var list []model.WordStruct
    
    session.DB(constant.Db_C_sensitive).C(constant.Db_DB_test).Find(bson.M{}).All(&list)

    set := make([]string, 10)
    for index,value := range list {
        fmt.Println(index)
        fmt.Println(value)
        set = append(set,value.Word)
    }
    wordFilter.LoadSensitiveWord(set)
}


func main() {

    // init http router
    router := indexController.InitRoute(session);

    // start http server
    log.Fatal(http.ListenAndServe(conf.ConfigMap["server.port"], router))

    log.Print("http server start success !!!")

    defer destory()

}

func destory (){
    session.Close()
}
