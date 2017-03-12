package wxmsg 

import ( 
    "crypto/sha1"
    "fmt" 
    "io"
    "net/http"
    "sort"
) 

func sha1Slice(strs []string) string {
    sh := sha1.New()
    for i := 0 ; i < len(strs); i++ {
        io.WriteString(sh, strs[i])
    }
    return fmt.Sprintf("%x", sh.Sum(nil))
}

func Verfy(r *http.Request) bool {
    return true
    // r.ParseForm()
    timestamp := r.FormValue("timestamp")
    nonce := r.FormValue("nonce")
    strs := []string{varWxToken,nonce,timestamp}
    sort.Sort(sort.StringSlice(strs))

    sum := sha1Slice(strs)
    signature := r.FormValue("signature")
    if sum == signature {
        return true
    } else {
        return false
    }

}

func NullResp(w http.ResponseWriter) {
    w.Write([]byte("success"))
}
