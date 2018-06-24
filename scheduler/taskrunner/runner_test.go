package taskrunner

import (
	"testing"
	"log"
	"errors"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i:=0;i<30;i++ {
			dc <- i;
			log.Printf("Dispatacher set : %v",i)
		}
		return nil
	}

	e := func(dc dataChan) error {
		forloop:
			for {
				select {
				case d:=<- dc :
					log.Printf("Excutor received: %v",d)
				default:
					break forloop
				}
			}
		return errors.New("over")
	}

	runner := NewRunner(30,false,d,e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}