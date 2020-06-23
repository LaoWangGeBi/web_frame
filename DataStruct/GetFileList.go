package GetFiles

// 文件列表
type FileInfoObj struct {
    F_name  string  `json:"F_name"`
    F_dir   string  `json:"F_dir"`
    F_type  string  `json:"F_type"`
    F_type2 string  `json:"F_type2"`
    F_list  []FileInfoObj  `json:"F_list"`
}
