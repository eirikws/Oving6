package main

import (
    "net"
    "fmt"
    "strings"
    "time"
    "strconv"
    "os/exec"
)
type Time struct{


}
func main(){
    port := "20001"
    IP :=GetMyIP()
    num:=0
    var eksint int
    num=ListenerCon(IP,port)
    cmd:=exec.Command("mate-terminal","-x","go","run","main.go")
    cmd.Run()
    for{
        time1:=time.Now().Unix()
        time2:=time.Now().Unix()
        for time1<time2+1{
            time1=time.Now().Unix()    
        }
        num=num+1
        eksint=int(num)
        fmt.Println(num)
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

func ListenerCon(ipAdr string, port string) int {
    serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
    psock, err := net.ListenUDP("udp4", serverAddr)
    var num=0
    if err != nil { 
        fmt.Println(err)
        return 0
    }
    buf := make([]byte,1024)
    for {
        if err != nil { return 0}
        psock.SetDeadline(time.Now().Add(2*time.Second))
        _,_,err:=psock.ReadFromUDP(buf)
        if err!=nil{
            psock.Close()
            fmt.Println("print num",num)
            return num
        }
        readInt:=string(buf)
        fmt.Println("%s",readInt)
        n,_:=strconv.Atoi(readInt)
        fmt.Println(n)
        
    }              
}




