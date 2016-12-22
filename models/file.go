package models

type FileStruct struct {
    Name string 	`json:"name"`
    IsDir bool	`json:"is_dir"`
    IsDeleteable bool	`json:"is_deleteable"`
    IsExecutable bool	`json:"is_executable"`
    IsReadable bool	`json:"is_readable"`
    IsWritable bool	`json:"is_writable"`
    Path string	`json:"path"`
    Mtime int64	`json:"mtime"`
    Size int64	`json:"size"`
}

