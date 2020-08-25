package GetFileList

import (
    "io/ioutil"
    "strings"
    "web_frame/go/DataStruct"
)

// func main() {
//     _arr := ListFile(`D:\Work`)
//     //fmt.Print(_arr);
//     //var f_obj FileInfoObj
//     for i, val := range _arr{
//         fmt.Print(i)
//         fmt.Print(`--------`)
//         //b, err = json.Marshal(val)
//         fmt.Println(val)
//         fmt.Println(`---------------------------------`)
//     }
// }

// 获取文件列表(全部)
func ListFileAll(myfolder string) []GetFiles.FileInfoObj {
    files, _ := ioutil.ReadDir(myfolder)
    var f_len int = len(files)
    var _arr = make([]GetFiles.FileInfoObj,f_len)
    var f_obj GetFiles.FileInfoObj
    for i, file := range files {
        f_obj.F_dir = myfolder+`\`+file.Name()
        f_obj.F_name = file.Name()
        if file.IsDir() {
            f_obj.F_type = "dir"
            f_obj.F_type2 = ""
            f_obj.F_list = ListFile(f_obj.F_dir)
        } else {
            f_obj.F_type = "file"
            f_obj.F_type2 = strings.Split(f_obj.F_name,".")[len(strings.Split(f_obj.F_name,"."))-1]
        }
        _arr[i] = f_obj
    }
    return _arr
}

// 获取文件列表(一级)
func ListFile(myfolder string) []GetFiles.FileInfoObj {
    files, _ := ioutil.ReadDir(myfolder)
    var f_len int = len(files)
    var _arr = make([]GetFiles.FileInfoObj,f_len)
    var f_obj GetFiles.FileInfoObj
    for i, file := range files {
        f_obj.F_dir = myfolder+`\`+file.Name()
        f_obj.F_name = file.Name()
        if file.IsDir() {
            f_obj.F_type = "dir"
            f_obj.F_type2 = ""
        } else {
            f_obj.F_type = "file"
            f_obj.F_type2 = strings.Split(f_obj.F_name,".")[len(strings.Split(f_obj.F_name,"."))-1]
        }
        _arr[i] = f_obj
    }
    return _arr
}