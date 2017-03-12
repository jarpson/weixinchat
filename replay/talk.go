package replay 

import ( 
    "encoding/json" 
    "fmt" 
    "io/ioutil"
    "log"
    "net/http"
) 

type TalkApi struct {
    Result   int      `json:"result"`
    Content  string   `json:"content"`
}

func httpGet(url string) (string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("http.Get err:%s\n", url)
        return ""
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("ioutil.ReadAll err:%v\n", err)
        return ""
    }

    var res TalkApi
    err = json.Unmarshal(body, &res) 
    if err != nil {
        log.Printf("json.Unmarshal err:%v\n", err)
        return ""
        if res.Result != 0 {
            log.Printf("api ret:%d\n", res.Result)
            return ""
        }
    }
    return res.Content
}

// query by api
func GetResp(search string) (answer string) {
    if search != "" {
        answer = httpGet(fmt.Sprintf("http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s", search))
    }
    if answer == "" {
        answer = "sorry,I cannot understand"
    }
    return
}


