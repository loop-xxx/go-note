package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func _pull(urls chan string, contexts chan []byte){
	for url := range urls{
		if res, err := http.Get(url); err == nil{
			if context, err := ioutil.ReadAll(res.Body); err == nil{
				contexts <- context
			}else {
				fmt.Printf("[WARING] %v\n", err)
			}
		}else{
			fmt.Printf("[WARING] %v\n", err)
		}
	}
}

func pull(wgp *sync.WaitGroup, urls chan string, contexts chan []byte){
	defer wgp.Done()
	_pull(urls, contexts)
}

const pullerCount = int(0x4)

func puller(urls chan string, contexts chan []byte){
	var wg sync.WaitGroup
	wg.Add(pullerCount)
	for i := int(0); i < pullerCount; i++{
		go pull(&wg, urls, contexts)
	}
	_pull(urls, contexts)

	wg.Wait()
	close(contexts)
}

func _parse( contexts chan []byte, items chan string){
	for context := range contexts{
		items <- string(context)
	}
}

func parse(wgp *sync.WaitGroup, contexts chan []byte, items chan string){
	defer wgp.Done()
	_parse(contexts, items)
}

func parser( contexts chan []byte, items chan string){
	var wg sync.WaitGroup


	wg.Add(parserCount)
	for i := int(0); i < parserCount; i ++{
		go parse(&wg, contexts, items)
	}
	_parse(contexts, items)

	wg.Wait()
	close(items)
}


const parserCount = int(0x4)
func run(page uint){
	var urls = make(chan string)
	var contexts = make (chan []byte)
	var items = make (chan string)

	go parser(contexts, items)

	go puller(urls,contexts)

	go func (){
		for i:= uint(0); i < page ; i++{
			fmt.Printf("pull page %d\n", i+1)
			urls <-fmt.Sprintf("https://sh.58.com/chuzu/pn%d", i+1)
			time.Sleep(time.Second*4)
		}
		close(urls)
	}()

	for item := range items {
		fmt.Println(item)
	}
}

func main(){
	pp := flag.Uint("page", 1, "抓取页面的数量")
	flag.Parse()

	run(*pp)
}



