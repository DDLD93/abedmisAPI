package routes

import (
	"fmt"
	"mime/multipart"
	"net/http"

	controller "github.com/ddld93/abedmis/src/controllers"
	"github.com/ddld93/abedmis/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

	// or for multiple files
	// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}
type BeneRoute struct {
	BenenCtrl *controller.DB_Connect
}
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}

var validate *validator.ValidationErrors

// func (ctrl *BeneRoute) FormTest() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var form ProfileForm
// 		// in this case proper binding will be automatically selected
// 		if err := c.ShouldBind(&form); err != nil {
// 			c.String(http.StatusBadRequest, "bad request")
// 			return
// 		}

// 		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
// 		if err != nil {
// 			c.String(http.StatusInternalServerError, "unknown error")
// 			return
// 		}

// 		// db.Save(&form)

// 		c.String(http.StatusOK, "ok")

// 	}
// }

func (ctrl *BeneRoute) AddMany() gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate := validator.New()
		var beneficiaries []models.Beneficiary
		//validate the request body
		if err := c.BindJSON(&beneficiaries); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Status:  http.StatusBadRequest,
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}
		// if err := validate.Struct(beneficiaries); err != nil {
		// 	c.JSON(http.StatusBadRequest, Response{
		// 		Status:  http.StatusBadRequest,
		// 		Message: "Invalid request Body",
		// 		Payload: nil,
		// 		Error:   err})
		// 	return
		// }
		res, err := ctrl.BenenCtrl.AddMany(beneficiaries)

		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Status:  http.StatusInternalServerError,
				Message: "An Error occured adding to database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusCreated, Response{
			Status:  http.StatusCreated,
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}

func (ctrl *BeneRoute) AddOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		var beneficiaries models.Beneficiary
		//validate the request body
		if err := c.BindJSON(&beneficiaries); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Status:  http.StatusBadRequest,
				Message: "Error parsing request body",
				Payload: nil,
				Error:   err.Error(),
			})
			return
		}
		beneficiaries.PSP = c.GetHeader("psp")
		beneficiaries.Mode = c.GetHeader("mode")
		fmt.Print(beneficiaries)
		res, err := ctrl.BenenCtrl.AddOne(beneficiaries)

		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Status:  http.StatusInternalServerError,
				Message: "An Error occured adding to database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusCreated, Response{
			Status:  http.StatusCreated,
			Message: "success",
			Payload: res,
			Error:   "",
		})
		return
	}
}

func (ctrl *BeneRoute) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pspKey := c.GetHeader("psp")
		mode := c.GetHeader("mode")
		res, err := ctrl.BenenCtrl.GetAll(pspKey,mode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Status:  http.StatusInternalServerError,
				Message: "An error occured querying database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{
			Status:  http.StatusOK,
			Message: "success",
			Payload: res,
			Error:   "",
		})
		
	}
}

func (ctrl *BeneRoute) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(Id)
		res, err := ctrl.BenenCtrl.GetOne(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Status:  http.StatusInternalServerError,
				Message: "An error occured querying database",
				Payload: nil,
				Error:   err.Error()})
			return
		}
		c.JSON(http.StatusCreated, Response{
			Status:  http.StatusCreated,
			Message: "success",
			Payload: res,
			Error:   "",
		})

	}
}
