package main 

import ( 
    "log"
    "net/http"
    "github.com/weixinchat/wxmsg"
) 

func main() { 
    log.Printf("Start serving!\n")
    http.HandleFunc("/wx/recver", wxmsg.ProxyMsg)
    err := http.ListenAndServe("127.0.0.1:8001", nil)
    if err != nil {
        log.Printf("lisetnAndServe err:%v\n", err)
    }
    log.Printf("quit\n")
}

