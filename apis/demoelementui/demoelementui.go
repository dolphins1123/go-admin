package demoelementui

import (
"go-admin/models"
"go-admin/tools"
"go-admin/tools/app"
"go-admin/tools/app/msg"

"github.com/gin-gonic/gin"
"github.com/gin-gonic/gin/binding"

)

func GetDemoElementuiList(c *gin.Context) {
var data models.DemoElementui
var err error
var pageSize = 10
var pageIndex = 1

if size := c.Request.FormValue("pageSize"); size != "" {
pageSize = tools.StrToInt(err, size)
}
if index := c.Request.FormValue("pageIndex"); index != "" {
pageIndex = tools.StrToInt(err, index)
}

data.Input1 = c.Request.FormValue("input1")
    

data.DataScope = tools.GetUserIdStr(c)
result, count, err := data.GetPage(pageSize, pageIndex)
tools.HasError(err, "", -1)

app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetDemoElementui(c *gin.Context) {
var data models.DemoElementui
data.Id, _ = tools.StringToInt(c.Param("id"))
result, err := data.Get()
tools.HasError(err, "抱歉未找到相关信息", -1)

app.OK(c, result, "")
}

func InsertDemoElementui(c *gin.Context) {
var data models.DemoElementui
err := c.ShouldBindJSON(&data)
data.CreateBy = tools.GetUserIdStr(c)
tools.HasError(err, "", 500)
result, err := data.Create()
tools.HasError(err, "", -1)
app.OK(c, result, "")
}

func UpdateDemoElementui(c *gin.Context) {
var data models.DemoElementui

err := c.BindWith(&data, binding.JSON)
tools.HasError(err, "数据解析失败!:" , -1)
data.UpdateBy = tools.GetUserIdStr(c)
result, err := data.Update(data.Id)
tools.HasError(err, "", -1)

app.OK(c, result, "")
}

func DeleteDemoElementui(c *gin.Context) {
var data models.DemoElementui
data.UpdateBy = tools.GetUserIdStr(c)

IDS := tools.IdsStrToIdsIntGroup("id", c)
_, err := data.BatchDelete(IDS)
tools.HasError(err, msg.DeletedFail, 500)
app.OK(c, nil, msg.DeletedSuccess)
}