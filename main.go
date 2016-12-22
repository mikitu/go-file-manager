package main

import (
	_ "github.com/mikitu/go-file-manager/routers"
	"github.com/astaxie/beego"
    "flag"
    "os"
)

func main() {
    baseDir := flag.String("base-dir", "", "set the base dir")
    flag.Parse()
    os.Setenv("APP_BASE_DIR", *baseDir)

	beego.Run()
}

