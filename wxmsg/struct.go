package wxmsg 

import ( 
    "encoding/xml" 
    "time"
)

var (
    varDefaultWelcom = "欢迎关注逻辑作者，主页君在此恭候多时了!\n"
    varNoSpportMsg = "你可以发文字或语音给我"
    varWxToken = "you token"
)

// weixin server send
type WxMsgReq struct {
    XMLName     	xml.Name `xml:"xml"`
    ToUserName  	string
    FromUserName	string
    CreateTime  	time.Duration
    MsgId       	int
    MsgType     	string

    MediaId         string      // image voice video .etc

    // event
    Event           string

    // text 
    Content     	string

    // image 
    PicUrl          string

    // voice
    Format          string
    Recognition     string

    // video/shortvideo
    ThumbMediaId    string

    // location
    Location_X      float64
    Location_Y      float64
    Scale           float64
    Label           string

    // link
    Title           string
    Description     string
    Url             string
}

type CDATA struct {
    Text            string `xml:",innerxml"`
}

// reply text to weixin server
type WxMsgRspText struct {
    XMLName         xml.Name `xml:"xml"`
    ToUserName      CDATA
    FromUserName    CDATA
    CreateTime      time.Duration
    MsgType         CDATA
    Content         CDATA
}

