package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
    "github.com/mikitu/go-file-manager/utils"
    "github.com/mikitu/go-file-manager/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) ListFiles() {
	fInput := c.Ctx.Input.Query("file")
	fileS := make([]*models.FileStruct, 0)
    baseDir := utils.GetBaseDir()
	files, err := ioutil.ReadDir(baseDir + fInput)
    if err != nil {
        utils.RenderError(&c.Controller, err.Error())
        return
    }
	for _, f := range files {
        path := utils.GetPath(fInput, f)
        filePath := baseDir + path
		fileS = append(fileS, &models.FileStruct{
            Name: f.Name(),
            Mtime: f.ModTime().Unix(),
            IsDir: f.IsDir(),
            Size: f.Size(),
            IsDeleteable: utils.IsDeleteadable(f, filePath),
            IsWritable: utils.IsWritable(filePath),
            IsExecutable: utils.IsExecutable(filePath),
            IsReadable: utils.IsReadable(filePath),
            Path: path,
        })
	}

	utils.RenderSuccess(&c.Controller, utils.IsWritable(fInput), fileS)
}

func (c *MainController) Delete() {

}
func (c *MainController) Download() {
    fInput := c.Ctx.Input.Query("file")
    c.Ctx.Output.Download(utils.GetBaseDir() + "/" + fInput)
}
func (c *MainController) Upload() {

}
func (c *MainController) NewFolder() {

}
func (c *MainController) View() {

}
func (c *MainController) Rename() {

}
