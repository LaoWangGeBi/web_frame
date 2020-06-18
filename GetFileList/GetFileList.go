package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "encoding/json"
)

// 文件列表
type FileInfoObj struct {
    F_name  string  `json:"F_name"`
    F_dir   string  `json:"F_dir"`
    F_type  string  `json:"F_type"`
    F_type2 string  `json:"F_type2"`
    F_list  []FileInfoObj  `json:"F_list"`
}

func main() {
    _arr := ListFile(`G:\waibao\mdk麦迪科\web`)
    //fmt.Print(_arr);
    var f_obj FileInfoObj
    for i, val := range _arr{
        fmt.Print(i)
        fmt.Print(`--------`)
        b, err = json.Marshal(val)
        fmt.Println(b)
        fmt.Println(`---------------------------------`)
    }
}

// 获取文件列表
func ListFile(myfolder string) []FileInfoObj {
    files, _ := ioutil.ReadDir(myfolder)
    var f_len int = len(files)
    var _arr = make([]FileInfoObj,f_len)
    //var _arr2 = [2]string{}
    f_obj := &FileInfoObj{}
    for i, file := range files {
        f_obj.F_dir = myfolder+`\`+file.Name()
        f_obj.F_name = file.Name()
        if file.IsDir() {
            f_obj.F_type = "dir"
            f_obj.F_type2 = ""
            f_obj.F_list = ListFile(myfolder+`\`+f_obj.F_name)
        } else {
            f_obj.F_type = "file"
            f_obj.F_type2 = strings.Split(f_obj.F_name,".")[1]
        }
        json_str, err := json.Marshal(f_obj)
        if err == nil{
            _arr[i] = f_obj
        }else{
            _arr[i] = ""
        }
    }
    fmt.Println(`||||||||||||||||||||||||||||||||||||||||||`)
    return _arr
}