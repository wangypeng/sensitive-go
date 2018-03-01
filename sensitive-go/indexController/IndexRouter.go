package indexController

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
)


func InitRoute(session *mgo.Session) *httprouter.Router {

	contoller := Controller{
		session : session,
	}

	route := httprouter.New()
	route.PUT("/add", log(contoller.addSensitive))
	route.GET("/check", log(contoller.check))
	route.DELETE("/del", contoller.delSensitive)
	return route

}

func log(h httprouter.Handle) httprouter.Handle{
	return func (w http.ResponseWriter, r *http.Request, params httprouter.Params){

		// before time
		beforeTime := time.Now()
		// before time stamp
		befortTimeStamp := beforeTime.Unix()
		fmt.Println(beforeTime)
		fmt.Println(befortTimeStamp)

		fmt.Println("==========  before  ===========")

		h(w,r,params)

		fmt.Println("==========  after  ===========")

		// after time
		afterTime := time.Now()
		// after time stamp
		aftertTimeStamp := afterTime.Unix()
		fmt.Println(afterTime)
		fmt.Println(aftertTimeStamp)

		waitTime := aftertTimeStamp - befortTimeStamp 
		fmt.Println("wait time : " + strconv.FormatInt(waitTime,10) +"ms")
	}
}
