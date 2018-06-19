package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"video_server/api/defs"
	"encoding/json"
	"video_server/api/dbops"
	"video_server/api/session"
	"fmt"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCreatedential{}
	if err := json.Unmarshal(res,ubody);err != nil {
		fmt.Print(err)
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd);err != nil {
		sendErrorResponse(w,defs.ErrorDBError)
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp,err := json.Marshal(su);err != nil {
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w,string(resp),201)
	}
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}