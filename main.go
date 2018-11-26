package main

import (
	"time"
	"net/http"
	"sort"
	"fmt"
)

func web() {
	http.HandleFunc("/", func(
		reponse http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(reponse, "<h1>Hello world.</h1>")
	})
	http.ListenAndServe(":8888", nil)
}

func web_concurrent(){
	ch := make(chan string)
	for i:=0;i<1000;i++ {
		go printString(i,ch)
	}
	for{
		msg := <-ch
		fmt.Println(msg)
	}
	time.Sleep(time.Second)
}

func printString(i int, ch chan string){
	for {
		ch <- fmt.Sprintf("Hello world from routeline %d", i)
	}
}

func st() {
	a := []int{5,7,3,2,8}

	sort.Ints(a)
	for _,v := range a{
		fmt.Print(v)
	}
}

func main() {

	web()
	web_concurrent()

	//st()
}
