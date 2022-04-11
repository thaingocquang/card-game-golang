package controller_test

import (
	"bytes"
	"card-game-golang/model"
	testhelper "card-game-golang/test_helper"
	"card-game-golang/util"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

// RegisterSuite ...
type RegisterSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *RegisterSuite) SetupSuite() {
	suite.e = testhelper.InitServer()
	testhelper.CreateFakePlayer()
}

// TestRegisterSuccess ...
func (suite *RegisterSuite) TestRegister_Success() {
	var (
		body = model.Player{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "success", response["message"])
}

// TestRegister_Fail_EmailRequired ...
func (suite *RegisterSuite) TestRegister_Fail_NameRequired() {
	var (
		body = model.Player{
			Name:     "",
			Email:    "test@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "name: cannot be blank.", response["message"])
}

// TestRegister_Fail_EmailRequired ...
func (suite *RegisterSuite) TestRegister_Fail_EmailRequired() {
	var (
		body = model.Player{
			Name:     "test",
			Email:    "",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "email: cannot be blank.", response["message"])
}

// TestRegister_Fail_PasswordRequired ...
func (suite *RegisterSuite) TestRegister_Fail_PasswordRequired() {
	var (
		body = model.Player{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "password: cannot be blank.", response["message"])
}

// TestRegister_Fail_EmailExisted ...
func (suite *RegisterSuite) TestRegister_Fail_EmailExisted() {
	var (
		body = model.Player{
			Name:     "fake",
			Email:    "fake@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "email already existed", response["message"])
}

// TestRegister_Fail_EmailInValid_InvalidPrefix ...
func (suite *RegisterSuite) TestRegister_Fail_EmailInValid_InvalidPrefix() {
	var (
		body = model.Player{
			Name:     "test",
			Email:    "@mail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "email: must be a valid email address.", response["message"])
}

// TestRegister_Fail_EmailInValid_InvalidDomain ...
func (suite *RegisterSuite) TestRegister_Fail_EmailInValid_InvalidDomain() {
	var (
		body = model.Player{
			Name:     "test",
			Email:    "test@mail#.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "email: must be a valid email address.", response["message"])
}

// TestAuth ...
func TestAuth(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
