package handler

import (
	"github.com/gofiber/fiber/v2"
)

type TestAllResponse struct {
	RequestBody    RequestBody
	BodyParseError error
	ParameterOne   string
	ParameterTwo   string
	ParameterAll   map[string]string
	QueryParamOne  string
	QueryParamAll  QueryParams
	HeadersAll     map[string]string
}

type QueryParams struct {
	Q1 []string
	Q2 string
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

func TestAllHandler(c *fiber.Ctx) error {
	requestBody := RequestBody{}
	bodyParseError := c.BodyParser(&requestBody)

	parameterOne := c.Params("p1")
	parameterTwo := c.Params("p2")
	parameterAll := c.AllParams()

	queryParamOne := c.Query("q1")
	queryParamAll := QueryParams{}
	c.QueryParser(&queryParamAll)

	headersAll := c.GetReqHeaders()

	return c.JSON(TestAllResponse{
		RequestBody:    requestBody,
		BodyParseError: bodyParseError,
		ParameterOne:   parameterOne,
		ParameterTwo:   parameterTwo,
		ParameterAll:   parameterAll,
		QueryParamOne:  queryParamOne,
		QueryParamAll:  queryParamAll,
		HeadersAll:     headersAll,
	})
}