package v1

import (
	"github.com/ElijahCYN/Go-gin-api/models"
	"github.com/ElijahCYN/Go-gin-api/pkg/e"
	"github.com/ElijahCYN/Go-gin-api/pkg/logging"
	"github.com/ElijahCYN/Go-gin-api/pkg/setting"
	"github.com/ElijahCYN/Go-gin-api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)
// GetArticle
// @Summary Get a single article
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/articles/{id} [get]
// 獲取單一篇文章
func GetArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID 必須大於 0")
	
	code := e.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// GetArticles
// @Summary Get multiple articles
// @Produce json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/articles [get]
// GetArticles 獲取多篇文章
func GetArticles(c *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("狀態只允許 0 或 1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("標籤 ID 必須大於 0")
	}

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// AddArticle
// @Summary Add article
// @Produce json
// @Param tag_id body int true "TagID"
// @Param title body string true "title"
// @Param desc body string true "desc"
// @Param content body string true "content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "state"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context)  {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("標籤 ID 必須大於 0")
	valid.Required(title, "title").Message("標題不能為空")
	valid.Required(desc, "desc").Message("簡述不能為空")
	valid.Required(content, "content").Message("內容不能為空")
	valid.Required(createdBy, "created_by").Message("創建人不能為空")
	valid.Range(state, 0, 1, "state").Message("狀態只允許 0 或 1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface {})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.messag: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditArticle
// @Summary Update article
// @Produce json
// @Param id path int true "ID"
// @Param tag_id body string false "TagID"
// @Param title body string false "Title"
// @Param desc body string false "desc"
// @Param content body string false "Content"
// @Param modified_by body string false "ModifiedBy"
// @Param state body int false "State"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context)  {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo("state").MustInt()
		valid.Range(state, 0, 1, "state").Message("狀態只允許 0 或 1")
	}

	valid.Min(id, 1, "id").Message("ID 必須大於 0")
	valid.MaxSize(title, 100, "title").Message("標題最長為 100 字元")
	valid.MaxSize(desc, 255, "desc").Message("簡述最長為 255 字元")
	valid.MaxSize(content, 65535, "content").Message("內容最長為 65535 字元")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能為空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最長為 100 字元")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle
// @Summary Delete article
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} string string
// @Failure 500 {object} string string
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID 必須大於 0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}