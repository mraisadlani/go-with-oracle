package handler

import (
	"Training/go-crud-with-oracle/domain"
	"Training/go-crud-with-oracle/domain/message"
	"Training/go-crud-with-oracle/infrastructure/persistance/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	userCase *usecase.CustomerUsecase
}

func BuildHandler(r *gin.RouterGroup, userCase *usecase.CustomerUsecase) {
	setHandler := &Handler{
		userCase: userCase,
	}

	r.POST("/CreateUser", setHandler.CreateUser)
	r.GET("/GetUser/:id", setHandler.GetUser)
	r.GET("/ViewAllUser", setHandler.ViewAllUser)
	r.POST("/UpdateUser/:id", setHandler.UpdateUser)
	r.POST("/DeleteUser/:id", setHandler.DeleteUser)
	r.POST("/GetDataFromView", setHandler.GetDataFromView)
	r.POST("/GetDataFromFunction", setHandler.GetDataFromFunction)
	r.POST("/SetDataFromProcedure", setHandler.SetDataFromProcedure)
}

// @Summary Create User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Param reqBody body domain.RegisterDTO true "Form Request"
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /CreateUser [post]
func (r *Handler) CreateUser(c *gin.Context) {
	var registerDTO domain.RegisterDTO

	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		message.EntityException(c, err.Error())
		return
	}

	createUser := registerDTO.SetDataUser(registerDTO)

	set, err := r.userCase.CreateUser(createUser)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Create data successfully", set)
	}
}

// @Summary Get User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Param id path string true "ID User"
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /GetUser/{id} [get]
func (r *Handler) GetUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	set, err := r.userCase.FindUser(userID)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Get data successfully", set)
	}
}

// @Summary Get All User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /ViewAllUser [get]
func (r *Handler) ViewAllUser(c *gin.Context) {
	set, err := r.userCase.ViewAll()

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Get data successfully", set)
	}
}

// @Summary Update User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Param id path string true "ID User"
// @Param reqBody body domain.UpdateUserDTO true "Form Request"
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /UpdateUser/{id} [post]
func (r *Handler) UpdateUser(c *gin.Context) {
	var updateDTO domain.UpdateUserDTO

	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		message.EntityException(c, err.Error())
		return
	}

	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	updateUser := updateDTO.SetUpdateUser(updateDTO)

	set, err := r.userCase.UpdateUser(userID, updateUser)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Update data successfully", set)
	}
}

// @Summary Delete User
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Param id path string true "ID User"
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /DeleteUser/{id} [post]
func (r *Handler) DeleteUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	set, err := r.userCase.DeleteUser(userID)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Delete data successfully", set)
	}
}

// @Summary Get Data from View
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /GetDataFromView [post]
func (r *Handler) GetDataFromView(c *gin.Context) {
	set, err := r.userCase.GetDataByView()

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Get data successfully", set)
	}
}

// @Summary Get Data from Function
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /GetDataFromFunction [post]
func (r *Handler) GetDataFromFunction(c *gin.Context) {
	var dateDTO domain.DateRangeDTO

	if err := c.ShouldBindJSON(&dateDTO); err != nil {
		message.EntityException(c, err.Error())
		return
	}

	set, err := r.userCase.GetDataByFunction(&dateDTO)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Get data successfully", set)
	}
}

// @Summary Get Data from Procedure
// @Description REST API User
// @Accept json
// @Produce json
// @Tags User Controller
// @Success 200 {object} message.Response
// @Failure 500,400,404,403 {object} message.Response
// @Router /SetDataFromProcedure [post]
func (r *Handler) SetDataFromProcedure(c *gin.Context) {
	var dateDTO domain.DateRangeDTO

	if err := c.ShouldBindJSON(&dateDTO); err != nil {
		message.EntityException(c, err.Error())
		return
	}

	set, err := r.userCase.SetDataByProcedure(&dateDTO)

	if err != nil {
		message.AppException(c, err.Error())
		return
	} else {
		message.MessageSuccess(c, http.StatusOK, "Set data successfully", set)
	}
}