package main

import (
	"fmt"
	"net"
	"flag"
	"sync"
	)

func main(){

	ip :=flag.String("t","192.168.0.1","Ip address or hostname to scan")
	ports := flag.Int("p",1000,"Number of Ports to scan")
	flag.Parse()
	var wg sync.WaitGroup
	for i:=1;i<=*ports;i++{
		wg.Add(1)		
		go func(j int){
			defer wg.Done()
			address := fmt.Sprintf("%s:%d",*ip,j)
			conn,err:= net.Dial("tcp",address)
			if err == nil{
				fmt.Println("Port open: ",j)
				conn.Close()			
}
			
		}(i)
	}
	wg.Wait()
}
			
