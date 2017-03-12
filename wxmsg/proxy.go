package wxmsg 

import ( 
    "encoding/xml" 
    "io/ioutil"
    "log"
    "net/http"
)

func errDeal(w http.ResponseWriter, msg string) {
        log.Printf("err :%s\n", msg)
        NullResp(w)
}

// Router wx msg 
// spport text and enent now
func ProxyMsg(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        errDeal(w, "read")
        return
    }

    wxReq := &WxMsgReq{}
    err = xml.Unmarshal(body, &wxReq)
    if err != nil {
        errDeal(w, "xml.Unmarshal")
        return
    }
    //log.Printf("type:%s\n", wxReq.MsgType)
    //log.Printf("upk:%v\n", wxReq)

    var rsp []byte
    switch wxReq.MsgType {
        case "text":
            rsp, err = dealText(wxReq);
        case "voice":
            rsp, err = dealVoice(wxReq);
        case "event":
            rsp, err = dealEvent(wxReq);
        default:
            rsp, err = dealDefault(wxReq);
    }
    if err != nil {
        log.Println(err)
        if len(rsp) == 0 {
            rsp = []byte("success")
        }
    }
    w.Write(rsp)
}

