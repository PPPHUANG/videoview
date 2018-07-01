package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"video_server/scheduler/dbops"
	"log"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	log.Printf("delete vid is : %s", vid)
	if len(vid) == 0 {
		sendResponse(w,400,"Video id should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		log.Printf("error is : %s", err)
		sendResponse(w,500,"Internal server error")
		return
	}
	sendResponse(w,200,"")
	return
}