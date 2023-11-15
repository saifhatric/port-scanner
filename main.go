package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

func main() {
	var URL string
	var num string
	fmt.Print("Enter the website link you want to scan : ")
	fmt.Scan(&URL)
	fmt.Print("Enter the max number of ports u want to scan : ")
	fmt.Scan(&num)

	port := make(chan int, 100)
	result := make(chan int)
	openPorts := []int{}

	max,err:= strconv.Atoi(num)
	if err!=nil{
		fmt.Printf("error : %v",err)
	}
	go func(){
		for i := 1; i<= max; i++ {
			port <- i
		}
	}()
	//running the check for ports
	for i := 0; i < cap(port); i++ {
		go portChecker(URL,port,result)
	}

	for i:=0;i<max;i++{
		port:=<-result
		openPorts = append(openPorts, port)
	}

	close(port)
	close(result)

	sort.Ints(openPorts)
	for _,p:=range openPorts{
		fmt.Printf("port %d is open\n",p)
	}

}

func portChecker(url string,portNum, result chan int) {
	for p := range portNum {
		addr := fmt.Sprintf("%v:%d",url,p)
		conn,err:=net.Dial("tcp",addr)
		if err!=nil{
			return
			// result<-0
			// continue
		}
		conn.Close()
		result<-p
	}
}