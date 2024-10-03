package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CommonResponse struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}

	Paging struct {
		Paging      string `json:"paging"`
		RowsPerPage string `json:"rowsPerPage"`
		TotalRows   string `json:"totalRows"`
		TotalPage   string `json:"totalPage"`
	}

	CommonResponseWithPaging struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Paging     Paging      `json:"paging"`
	}
)

func NewResponseCreated(c *gin.Context, result interface{}) {
	c.JSON(http.StatusCreated, CommonResponse{
		StatusCode: http.StatusCreated,
		Message:    http.StatusText(http.StatusCreated),
		Data:       result,
	})
}

func NewResponseOK(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       result,
	})
}

func NewResponseOKWithPaging(c *gin.Context, result interface{}, paging, rowsPerPage, totalRows, totalPage string) {
	c.JSON(http.StatusOK, CommonResponseWithPaging{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       result,
		Paging: Paging{
			Paging:      paging,
			RowsPerPage: rowsPerPage,
			TotalRows:   totalRows,
			TotalPage:   totalPage,
		},
	})
}

func NewResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, CommonResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func NewResponseError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, CommonResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	})
}

func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, CommonResponse{
		StatusCode: http.StatusForbidden,
		Message:    message,
	})
}

func NewResponseValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, CommonResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "Validation Error",
		Data:       errors,
	})
}
