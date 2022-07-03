package handler

import (
	"net/url"

	"github.com/labstack/echo/v4"
)

type TestAllResponse struct {
	RequestBody    RequestBody
	BodyParseError error
	ParameterOne   string
	ParameterTwo   string
	ParameterAll   []string
	QueryParamOne  string
	QueryParamAll  url.Values
	HeadersAll     map[string]string
	HeadersParser  Headers
}

type RequestBody struct {
	FieldOne   string    `json:"field_one"`
	FieldTwo   int       `json:"field_two"`
	FieldThree []string  `json:"field_three"`
	FieldFour  FieldFour `json:"field_four"`
}

type FieldFour struct {
	ChildField string `json:"child_field"`
}

type QueryParams struct {
	Q1 []string `query:"q1"` // Bind works only for GET/DELETE methods
	Q2 string   `query:"q2"`
}

type Headers struct {
	UserAgent   string `header:"User-Agent"`
	Accept      string `header:"Accept"`
	ContentType string `header:"Content-Type"`
}

func TestAll(c echo.Context) error {
	var requestBody RequestBody
	bodyParseError := c.Bind(&requestBody)

	parameterOne := c.Param("p1")
	parameterTwo := c.Param("p2")
	parameterAll := c.ParamValues()

	queryParamOne := c.QueryParam("q1")
	queryParamAll := c.QueryParams()

	var headersParser Headers
	(&echo.DefaultBinder{}).BindHeaders(c, &headersParser)

	// c.Bind(&headersParser) // doesn't work for header?

	return c.JSON(200, TestAllResponse{
		RequestBody:    requestBody,
		BodyParseError: bodyParseError,
		ParameterOne:   parameterOne,
		ParameterTwo:   parameterTwo,
		ParameterAll:   parameterAll,
		QueryParamOne:  queryParamOne,
		QueryParamAll:  queryParamAll,
		// HeadersAll:     headersAll,
		HeadersParser: headersParser,
	})
}
