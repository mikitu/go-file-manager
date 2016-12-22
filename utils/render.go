package utils

import "github.com/astaxie/beego"

func render(c *beego.Controller, success bool, is_writable bool, results interface{}) {
    if success {
        c.Data["json"] = map[string]interface{}{"success": success, "is_writable" : is_writable, "data": results}
    } else {
        c.Data["json"] = map[string]interface{}{"success": success, "error": map[string]string{"msg": results.(string)}}
    }
    c.ServeJSON()
}

func RenderSuccess(c *beego.Controller, is_writable bool, results interface{}) {
    render(c, true, is_writable, results)
}
func RenderError(c *beego.Controller, results interface{}) {
    render(c, false, false, results)
}

