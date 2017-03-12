package wxmsg 

import ( 
    "encoding/xml" 
    "time"
)

func newData(v string) CDATA{
    return CDATA{"<![CDATA[" + v + "]]>"}
}

func RspText(wxReq *WxMsgReq, content string) ([]byte, error) {
    rsp := &WxMsgRspText{
        FromUserName: newData(wxReq.ToUserName),
        ToUserName: newData(wxReq.FromUserName),
        MsgType: newData("text"),
        Content: newData(content),
        CreateTime: time.Duration(time.Now().Unix()),
    }
    return xml.Marshal(rsp)
}

