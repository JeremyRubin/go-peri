package controllers

import (
	"bufio"
	"github.com/fzzy/sockjs-go/sockjs"
	"log"
	"os"
	"time"
	"encoding/json"
	"fmt"
)
type message struct {
	Name string
	Message string
}
type message_in struct {
	Name string
	Value int
	Message string
	State [2]bool
}
func (m *message_in) Represent() string{
	return fmt.Sprintf("%v \n",m.Value)
}
type device struct {
	clients *sockjs.SessionPool
	delay   time.Duration
}


func NewDevice(pool *sockjs.SessionPool, rate int) *device {
	delay := time.Duration(rate) * time.Millisecond
	dev := device{pool, delay}
	return &dev
}
func (d *device) AddReader(path string) {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	bufReader := bufio.NewReader(fi)
	go func() {
		defer fi.Close()
		for {
			time.Sleep(d.delay)
			if line, err := bufReader.ReadString('\n'); err != nil {
			//if line, err := bufReader.ReadByte(); err != nil {
				log.Print(err)
			} else {
				log.Print(string(line))
				m := message{Name:path, Message:string(line)}
				b, err := json.Marshal(m)
                if err!= nil{
                log.Print(err)
                }
				d.clients.Broadcast(b)
			}
		}
	}()
}

func (d *device) SocketHandler(s sockjs.Session) {
	d.clients.Add(s)
	defer d.clients.Remove(s)
	for {
		time.Sleep(d.delay)
		m := s.Receive()
		if m == nil {
			break
		}
		js := message_in{}
		err := json.Unmarshal(m, &js)
        log.Print(js)
        if err!=nil{
            log.Print(err)
        }
        fi, err := os.OpenFile(js.Name, os.O_WRONLY, 0666)
        if err!=nil{
            log.Print(err)
        }
        //
		_, err = fi.Write([]byte(js.Represent()))
      if err!=nil{
            log.Print(err)
        }
		log.Print(js.Represent())
    	fi.Close()
	}
}
