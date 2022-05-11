package v1

import (
	"gohub/app/models/user"
	"gohub/app/requests"
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
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	userPagination := user.Paginate(c, 10)
	response.Data(c, userPagination)
}
