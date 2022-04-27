package controller_test

import (
	"bytes"
	"card-game-golang/dto"
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "success", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "name: cannot be blank.", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "email: cannot be blank.", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "password: cannot be blank.", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "email already existed", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "email: must be a valid email address.", response.Message)
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
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "email: must be a valid email address.", response.Message)
}

// ===========================================

// LoginSuite ...
type LoginSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *LoginSuite) SetupSuite() {
	suite.e = testhelper.InitServer()
	testhelper.CreateFakePlayer()
}

// TestLogin_Success ...
func (suite *LoginSuite) TestLogin_Success() {
	var (
		body = dto.PlayerLogin{
			Email:    "fake@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response.Data)
}

// TestLogin_Fail_EmailNotExistInDB ...
func (suite *LoginSuite) TestLogin_Fail_EmailNotExistInDB() {
	var (
		body = dto.PlayerLogin{
			Email:    "adsdasd@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "email not existed in db", response.Message)
}

// TestLogin_Fail_WrongPassword ...
func (suite *LoginSuite) TestLogin_Fail_WrongPassword() {
	var (
		body = dto.PlayerLogin{
			Email:    "fake@gmail.com",
			Password: "asdasdaa",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "wrong password", response.Message)
}

// ===========================================

// LoginSuite ...
type AdminLoginSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *AdminLoginSuite) SetupSuite() {
	suite.e = testhelper.InitServer()
	testhelper.CreateFakePlayer()
}

// TestAdminLogin_Success ...
func (suite *LoginSuite) TestAdminLogin_Success() {
	var (
		body = dto.Admin{
			Username: "admin",
			Password: "123456",
		}
		response util.Response
	)

	bodyJSON, _ := json.Marshal(body)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// assert
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response.Data)
	assert.Equal(suite.T(), "success", response.Message)
}

// TestAuth ...
func TestAuth(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
	suite.Run(t, new(LoginSuite))
	suite.Run(t, new(AdminLoginSuite))
}
