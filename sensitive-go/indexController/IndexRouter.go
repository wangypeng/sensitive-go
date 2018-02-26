package indexController

import (
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
)


func InitRoute(session *mgo.Session) *httprouter.Router {

	contoller := Controller{
		session : session,
	}

	route := httprouter.New()
	route.PUT("/add", contoller.addSensitive)
	route.GET("/check", contoller.check)
	route.DELETE("/del", contoller.delSensitive)
	return route

}
