package topic

import (
    "gohub/pkg/app"
    "gohub/pkg/database"
    "gohub/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idStr string) (topic Topic) {
    database.DB.Where("id", idStr).First(&topic)
    return
}

func GetBy(field, value string) (topic Topic) {
    database.DB.Where("? = ?", field, value).First(&topic)
    return
}

func All() (topics []Topic) {
    database.DB.Find(&topics)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}


// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) paginator.PaginationData[Topic] {
	return paginator.Paginate[Topic](
		c,
		database.DB.Model(Topic{}),
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)

}
