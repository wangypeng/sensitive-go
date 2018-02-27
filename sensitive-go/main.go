package main

import (
    "fmt"
	"log"
	
	"net/http"

    "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
	
    "sensitive-go/sensitive-go/indexController"
    // "goSensitive/sensitivefilter/model"
    "sensitive-go/sensitive-go/wordFilter"
    "sensitive-go/sensitive-go/conf"
    "sensitive-go/sensitive-go/constant"
)

var session *mgo.Session


type word struct{
	Word string		
	TimeStamp int	
}


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
    var list []word
    
    session.DB(constant.Db_C_sensitive).C(constant.Db_DB_test).Find(nil).All(&list)

    fmt.Println(list)

    set := make([]string,0)
    for index,value := range list {
        fmt.Println(index)
        fmt.Println([]rune(value.Word))
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
