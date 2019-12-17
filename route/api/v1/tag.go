package v1

import (
	"blog.com/common"

	"blog.com/pkg/models"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit", "15")).MustInt()

	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if name != "" {
		maps["name"] = name
	}
	data, err := models.GetTags(page, limit, maps)
	if err != nil {
		common.ReturnJson(c, 0, "查询失败", nil)
	} else {
		common.ReturnJson(c, 1, "", data)
	}
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			message := err.Key + err.Message
			common.ReturnJson(c, 0, message, nil)
			return
		}
	}
	//添加到数据库
	result, _ := models.AddTag(name)
	if !result {
		common.ReturnJson(c, 0, "添加失败", nil)
		return
	}
	common.ReturnJson(c, 1, "", nil)
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	check, _ := models.CheckTag(map[string]interface{}{"id": id})
	if !check {
		common.ReturnJson(c, 0, "找不到数据", nil)
		return
	}
	name := c.Query("name")
	if name == "" {
		common.ReturnJson(c, 0, "数据不能为空", nil)
		return
	}
	data := make(map[string]interface{})
	data["name"] = name
	update := models.EditTags(id, data)
	if update {
		common.ReturnJson(c, 1, "", nil)
		return
	} else {
		common.ReturnJson(c, 0, "更新失败", nil)
		return
	}
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	check, _ := models.CheckTag(map[string]interface{}{"id": id})
	if !check {
		common.ReturnJson(c, 0, "找不到数据", nil)
		return
	}
	update := models.DeleteTags(id)
	if update {
		common.ReturnJson(c, 1, "", nil)
		return
	} else {
		common.ReturnJson(c, 0, "删除失败", nil)
		return
	}
}
