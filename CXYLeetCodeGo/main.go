package main

import (
	"fmt"
	"sync"
	"time"

	"rsc.io/quote"

	"example/calc"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("startdownload", url)
	time.Sleep(time.Second)
	wg.Done()
}

var ch = make(chan string, 10)

func chanDownloac(url string) {
	fmt.Println("start to download:", url)
	time.Sleep(time.Second)
	ch <- url
}

func downloadTest(){
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))

	}
	wg.Wait()
	fmt.Println("Done!")
}

func chanDownloadTest(){
	for i := 0; i < 100; i++ {
		go chanDownloac("a.com/"+string(i))
	}
	for i := 0; i < 100; i++ {
		mas:=<-ch
		fmt.Println("finish:",mas)
	}
	fmt.Println("done!")
}

func TestDefer(){
	fmt.Println(calc.Get(4))
}

func Test3PartPkg(){
	fmt.Println(calc.Add(1, 2))
	fmt.Println(quote.Hello())
}

type Data struct {
	value int
	isProcessed chan bool
}

// func TestChan(){
// 	inChan:=make(chan *Data,2)
// 	outChanNum:=3
// 	outChans:=make([]chan *Data,0,outChanNum)
// 	for i := 0; i < outChanNum; i++ {
// 		outChans=append(outChans, make(chan *Data, 1))

// 	}

// }

func main() {


}
