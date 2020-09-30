package models

import (
orm "go-admin/global"
"go-admin/tools"

"time" 
)

type DemoElementui struct {

        Id int `json:"id" gorm:"type:int(11);primary_key"` // id
        Input1 string `json:"input1" gorm:"type:varchar(20);"` // 輸入框
        Number1 int64 `json:"number1" gorm:"type:int(11);"` // 計數器
        Select1 string `json:"select1" gorm:"type:varchar(20);"` // 下拉選項1
        Select2 string `json:"select2" gorm:"type:varchar(50);"` // 下拉多選
        Select3s string `json:"select3s" gorm:"type:varchar(50);"` // 下拉多選1
        Select3m string `json:"select3m" gorm:"type:varchar(100);"` // Cascader级联选择器
        Switch1 string `json:"switch1" gorm:"type:varchar(1);"` // Switch开关
        Datepicker1 time.Time `json:"datepicker1" gorm:"type:timestamp;"` // DatePicker
        Timepicker1 string `json:"timepicker1" gorm:"type:time;"` // timepicker1
        Upload1 string `json:"upload1" gorm:"type:varchar(100);"` // Upload 上传
        Tansfer1 string `json:"tansfer1" gorm:"type:varchar(100);"` // Transfer
        CreateBy string `json:"createBy" gorm:"type:varchar(20);"` // 
        UpdateBy string `json:"updateBy" gorm:"type:varchar(20);"` // 
DataScope   string `json:"dataScope" gorm:"-"`
Params      string `json:"params"  gorm:"-"`
BaseModel
}

func (DemoElementui) TableName() string {
return "demo_elementui"
}

// 创建DemoElementui
func (e *DemoElementui) Create() (DemoElementui, error) {
var doc DemoElementui
result := orm.Eloquent.Table(e.TableName()).Create(&e)
if result.Error != nil {
err := result.Error
return doc, err
}
doc = *e
return doc, nil
}


// 获取DemoElementui
func (e *DemoElementui) Get() (DemoElementui, error) {
var doc DemoElementui
table := orm.Eloquent.Table(e.TableName())

    
        if e.Id != 0 {
        table = table.Where("id = ?", e.Id)
        }
    
        if e.Input1 != ""  {
        table = table.Where("input1 = ?", e.Input1)
        }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

if err := table.First(&doc).Error; err != nil {
return doc, err
}
return doc, nil
}

// 获取DemoElementui带分页
func (e *DemoElementui) GetPage(pageSize int, pageIndex int) ([]DemoElementui, int, error) {
var doc []DemoElementui

table := orm.Eloquent.Select("*").Table(e.TableName())

        if e.Input1 != ""  {
        table = table.Where("input1 = ?", e.Input1)
        }
    

// 数据权限控制(如果不需要数据权限请将此处去掉)
dataPermission := new(DataPermission)
dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
table,err := dataPermission.GetDataScope(e.TableName(), table)
if err != nil {
return nil, 0, err
}
var count int

if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
return nil, 0, err
}
table.Where("`deleted_at` IS NULL").Count(&count)
return doc, count, nil
}

// 更新DemoElementui
func (e *DemoElementui) Update(id int) (update DemoElementui, err error) {
if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
return
}

//参数1:是要修改的数据
//参数2:是修改的数据
if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
return
}
return
}

// 删除DemoElementui
func (e *DemoElementui) Delete(id int) (success bool, err error) {
if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&DemoElementui{}).Error; err != nil {
success = false
return
}
success = true
return
}

//批量删除
func (e *DemoElementui) BatchDelete(id []int) (Result bool, err error) {
if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&DemoElementui{}).Error; err != nil {
return
}
Result = true
return
}
