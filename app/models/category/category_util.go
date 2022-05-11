package category

import (
    "gohub/pkg/app"
    "gohub/pkg/database"
    "gohub/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idStr string) (category Category) {
    database.DB.Where("id", idStr).First(&category)
    return
}

func GetBy(field, value string) (category Category) {
    database.DB.Where("? = ?", field, value).First(&category)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}


// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) paginator.PaginationData[Category] {
	return paginator.Paginate[Category](
		c,
		database.DB.Model(Category{}),
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)

}
