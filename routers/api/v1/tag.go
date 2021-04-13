package v1

import (
	"github.com/ElijahCYN/Go-gin-api/models"
	"github.com/ElijahCYN/Go-gin-api/pkg/e"
	"github.com/ElijahCYN/Go-gin-api/pkg/setting"
	"github.com/ElijahCYN/Go-gin-api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags
// @Summary Get multiple article tags
// @Produce json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// AddTag
// @Summary Add article tag
// @Produce json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名稱不能為空")
	valid.MaxSize(name, 100, "name").Message("名稱最長為 100 字元")
	valid.Required(createdBy, "created_by").Message("創建人不能為空")
	valid.MaxSize(createdBy, 100, "created_by").Message("創建人名稱最長為 100 字元")
	valid.Range(state, 0, 1, "state").Message("狀態只允許 0 或 1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}


// EditTag
// @Summary Update article tag
// @Produce json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("狀態只允許 0 或 1")
	}

	valid.Required(id, "id").Message("ID 不能為空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能為空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最長為 100 字元")
	valid.MaxSize(name, 100, "name").Message("名稱最長為 100 字元")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag
// @Summary Delete article tag
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID 必須大於 0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
