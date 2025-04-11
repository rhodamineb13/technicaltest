package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"technicaltest/model"
	"technicaltest/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.IService
}

func NewHandler(serv service.IService) *Handler {
	return &Handler{
		Service: serv,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var req model.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	if err := h.Service.Create(c, req); err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	c.JSON(http.StatusOK, &model.ModelResponse{
		StatusCode: http.StatusOK,
		Message:    "new client created",
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	res, err := h.Service.Select(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.ModelResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	c.JSON(http.StatusOK, &model.ModelResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprint("error: ", err),
		Data:       res,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}
	res, err := h.Service.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.ModelResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprint("error: ", err),
		})
	}

	c.JSON(http.StatusOK, &model.ModelResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprint("error: ", err),
		Data:       res,
	})
}

func (h *Handler) Update(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	var req model.UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	if err := h.Service.Update(c, id, req); err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}

	c.JSON(http.StatusOK, &model.ModelResponse{
		StatusCode: http.StatusOK,
		Message:    "client updated",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ModelResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprint("error: ", err),
		})
		return
	}
	err = h.Service.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.ModelResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprint("error: ", err),
		})
	}

	c.JSON(http.StatusOK, &model.ModelResponse{
		StatusCode: http.StatusOK,
		Message:    "client deleted",
	})
}