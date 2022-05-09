package v1

import (
	"gohub/app/models/user"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	users := auth.CurrentUser(c)
	response.Data(c, users)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	userPagination := user.Paginate(c, 10)
	response.Data(c, userPagination)
}
