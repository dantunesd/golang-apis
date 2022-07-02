package handler

import (
	"github.com/gin-gonic/gin"
)

type TestAllResponse struct {
	RequestBody    RequestBody
	BodyParseError string
	ParameterOne   string
	ParameterTwo   string
	ParameterAll   gin.Params
	QueryParamOne  string
	QueryParamAll  QueryParams
	HeadersAll     Headers
}

type RequestBody struct {
	FieldOne   string    `json:"field_one" binding:"required"`
	FieldTwo   int       `json:"field_two"`
	FieldThree []string  `json:"field_three"`
	FieldFour  FieldFour `json:"field_four"`
}

type FieldFour struct {
	ChildField string `json:"child_field"`
}

type QueryParams struct {
	Q1 []string `form:"q1"`
	Q2 string   `form:"q2"`
}

type Headers struct {
	UserAgent   string `header:"User-Agent"`
	Accept      string `header:"Accept"`
	ContentType string `header:"Content-Type"`
}

func TestAll(c *gin.Context) {
	var errorString string
	var requestBody RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		errorString = err.Error()
	}

	parameterOne := c.Param("p1")
	parameterTwo := c.Param("p2")
	parameterAll := c.Params

	queryParamOne := c.Query("q1")

	var queryParamAll QueryParams
	c.ShouldBindQuery(&queryParamAll)

	var headersAll Headers
	c.ShouldBindHeader(&headersAll)

	c.JSON(200, TestAllResponse{
		RequestBody:    requestBody,
		BodyParseError: errorString,
		ParameterOne:   parameterOne,
		ParameterTwo:   parameterTwo,
		ParameterAll:   parameterAll,
		QueryParamOne:  queryParamOne,
		QueryParamAll:  queryParamAll,
		HeadersAll:     headersAll,
	})
}
