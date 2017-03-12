package wxmsg 

import ( 
    "log"
    "github.com/weixinchat/replay"
)

// text message deal
func dealText(wxReq *WxMsgReq) ([]byte, error) {
    talk := replay.GetResp(wxReq.Content)
    return RspText(wxReq, talk)
}

// voice message deal
func dealVoice(wxReq *WxMsgReq) ([]byte, error) {
    talk := replay.GetResp(wxReq.Recognition)
    return RspText(wxReq, talk)
}

// event deal
func dealEvent(wxReq *WxMsgReq) ([]byte, error) {
    switch wxReq.Event {
        case "subscribe": // new user
            return RspText(wxReq, varDefaultWelcom)
    }
    return []byte("success"), nil
}

// default message deal
func dealDefault(wxReq *WxMsgReq) ([]byte, error) {
    log.Printf("no spport type:%s\n", wxReq.MsgType)
    return RspText(wxReq, varNoSpportMsg)
}

