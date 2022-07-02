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
	Q1 []string
	Q2 string
}

type Headers struct {
	UserAgent   string `reqHeader:"User-Agent"`
	Accept      string `reqHeader:"Accept"`
	ContentType string `reqHeader:"Content-Type"`
}

func TestAll(c *fiber.Ctx) error {
	var requestBody RequestBody
	bodyParseError := c.BodyParser(&requestBody)

	parameterOne := c.Params("p1")
	parameterTwo := c.Params("p2")
	parameterAll := c.AllParams()

	queryParamOne := c.Query("q1")

	var queryParamAll QueryParams
	c.QueryParser(&queryParamAll)

	headersAll := c.GetReqHeaders()

	var headersParser Headers
	c.ReqHeaderParser(&headersParser)

	return c.Status(200).JSON(TestAllResponse{
		RequestBody:    requestBody,
		BodyParseError: bodyParseError,
		ParameterOne:   parameterOne,
		ParameterTwo:   parameterTwo,
		ParameterAll:   parameterAll,
		QueryParamOne:  queryParamOne,
		QueryParamAll:  queryParamAll,
		HeadersAll:     headersAll,
		HeadersParser:  headersParser,
	})
}
