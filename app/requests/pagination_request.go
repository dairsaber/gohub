package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort     string `valid:"sort" form:"sort"`
	Order    string `valid:"order" form:"order"`
	PageSize string `valid:"page_size" form:"page_size"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":      []string{"in:id,created_at,updated_at"},
		"order":     []string{"in:asc,desc"},
		"page_size": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,created_at,updated_at",
		},
		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"page_size": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
	}
	return validate(data, rules, messages)
}
