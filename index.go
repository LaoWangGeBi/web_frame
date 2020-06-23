package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "web_frame/GetFileList"
    //"encoding/json"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    // fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println("URL", r.URL)
    // fmt.Println(r.Form["url_long"])
    _arr := GetFileList.ListFile(`D:\Work`)
    //b, _ := json.Marshal(_arr);
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    for k, v := range _arr {
        fmt.Println("key:", k)
        fmt.Println("val:", v)
        fmt.Println("$$$$$$$$$$$$$$$$$$$$")
    }
    fmt.Fprintf(w, "asdfsadf") //这个写入到w的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
    err := http.ListenAndServe(":5454", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}