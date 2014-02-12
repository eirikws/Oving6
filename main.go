package main

import (
    "net"
    "fmt"
    "strings"
    "time"
    "strconv"
)
type Time struct{


}



func main(){
    port := "20001"
    IP :=GetMyIP()
    num:=0
    var eksint int
    ListenerCon(IP,port)
    for{
        time1:=time.Now().Unix()
        time2:=time.Now().Unix()
        for time1<time2+2{
            time1=time.Now().Unix()    
        }
        num=num+1
        eksint=int(num)
        go UDPcon(IP,port,eksint)
    }
    return



}


func GetMyIP() string{
    allIPs,err:=net.InterfaceAddrs()
    if err!=nil{
        fmt.Println("IP receiving errors!!!!!!!!\n")
        return ""
    }
    return strings.Split(allIPs[1].String(),"/")[0]
}

func UDPcon(IP string,port string, num int){
    serverAddr,err2:=net.ResolveUDPAddr("udp",IP+":"+port)
    con,err1:=net.DialUDP("udp",nil,serverAddr)
    if err1 != nil {
                fmt.Println("faileee1111")
    }
    if err2 != nil {
                fmt.Println("faileee2222")
    }
    Snum:=strconv.Itoa(num)
    Bmessage:=[]byte(Snum)
    con.Write(Bmessage)
}

func ListenerCon(ipAdr string, port string){
    serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
    psock, err := net.ListenUDP("udp4", serverAddr)
    if err != nil { return }
    buf := make([]byte,1024)
    var readInt strint:=time.Duration
    psock.SetDeadline(t)
    for {
        time1:=time.Now().Unix()
        time2:=time.Now().Unix()
        for time1<time2+5{
            time1=time.Now().Unix()
        }
        if err != nil { return }
        fmt.Println("teststart")
        _,_,err=psock.ReadFromUDP(buf)
        fmt.Println("testStop")
        readInt:=string(buf)
        fmt.Printf("%s\n",readInt)
    }              
}




