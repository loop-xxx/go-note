package session

import (
	"fmt"
	"testing"
	"time"
)

func TestBuildMemManager(t *testing.T){
	manger := BuildMemManager(0x10, time.Second * 5, time.Second*6)
	defer manger.Done()

	sessionID,_ := manger.NewSession(0x10)
	if s, ok := manger.GetSessionByID(sessionID);ok {
		if _, ok := s.Get("user"); !ok{
			fmt.Println("user not exist!")
		}
		s.Set("user", "loop")

		if username, ok := s.Get("user"); ok{
			fmt.Printf("%s\n", username)
		}
		s.Del("user")
		if _, ok := s.Get("user"); !ok{
			fmt.Println("user not exist!")
		}
	}else{
		t.Fail()
	}
	sid, _ := manger.NewSession(0x10)
	fmt.Println(sid)
	time.Sleep(time.Second * 10)

}
