package v1

import (
	"net/http"

	"blog.com/common"

	"blog.com/pkg/models"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个文章标签
func GetTags(c *gin.Context) {

}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := 200
	message := "成功"
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			code = http.StatusBadRequest
			message = err.Key + err.Message
			common.ReturnJson(c, code, message, nil)
			return
		}
	}
	//添加到数据库
	result, _ := models.AddTag(name, state)
	if !result {
		common.ReturnJson(c, http.StatusBadRequest, "添加失败", nil)
		return
	}
	common.ReturnJson(c, code, message, nil)
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
