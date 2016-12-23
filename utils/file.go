package utils

import (
    "golang.org/x/sys/unix"
    "os"
    "github.com/astaxie/beego"
    "log"
    "os/user"
    "strings"
)

func IsWritable(path string) bool {

    return unix.Access(path, unix.W_OK) == nil
}
func IsExecutable(path string) bool {
    return unix.Access(path, unix.X_OK) == nil
}
func IsReadable(path string) bool {
    return unix.Access(path, unix.R_OK) == nil
}
func IsDeleteadable(f os.FileInfo, path string) bool {
    allow_delete, _ := beego.AppConfig.Bool("allow_delete")
    if allow_delete && ((! f.IsDir() && IsWritable(path)) || (f.IsDir() && IsWritable(path) && IsDeleteableRecursive(path))) {
        return true
    }
    return false
}

func GetPath(fInput string, f os.FileInfo) string {
    var path string
    if  fInput == "" {
        path = f.Name()
    } else {
        path = fInput + "/" + f.Name()
    }
    return path
}

func GetBaseDir() string {
    base_dir := beego.AppConfig.String("base_dir")
    if dir := os.Getenv("APP_BASE_DIR"); dir != "" {
        base_dir = dir
    }
    if strings.HasPrefix(base_dir, "~/") {
        base_dir = GetHomeDir() + "/" + strings.Replace(base_dir, "~/", "", -1)
    }
    return base_dir
}

func GetHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}

func IsDeleteableRecursive(path string) bool {
    //TODO: implement
    return false
}