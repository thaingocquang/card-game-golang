package controller_test

import (
	"bytes"
	"card-game-golang/dao"
	"card-game-golang/model"
	testhelper "card-game-golang/test_helper"
	"card-game-golang/util"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MyProfileSuite ...
type MyProfileSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *MyProfileSuite) SetupSuite() {
	suite.e = testhelper.InitServer()
	testhelper.CreateFakePlayer()
}

// TestMyProfile_Success ...
func (suite *MyProfileSuite) TestMyProfile_Success() {
	var (
		token    string
		response util.Response
	)

	data := map[string]interface{}{
		"id": testhelper.PlayerObjID,
	}

	token, err := util.GenerateUserToken(data)
	if err != nil {
		panic(err)
	}

	// request
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := testhelper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// assert
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "success", response["message"])
}

// TestMyProfile_Fail_TokenInvalid ...
func (suite *MyProfileSuite) TestMyProfile_Fail_TokenInvalid() {
	var (
		response util.Response
		token    string
	)

	// TokenInvalid
	token = "abcxyz"

	// request
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := testhelper.RunAndAssertHTTPUnauthorized(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), "invalid or expired jwt", response["message"])
}

// TestGetMyProfile_Fail_MissingJWT ...
func (suite *MyProfileSuite) TestGetMyProfile_Fail_MissingJWT() {
	var (
		response util.Response
	)

	// request
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)

	// run
	rec := testhelper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), "missing or malformed jwt", response["message"])
}

// =============================================================

// UpdateMyProfile ...
type UpdateMyProfile struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *UpdateMyProfile) SetupSuite() {
	suite.e = testhelper.InitServer()
	testhelper.CreateFakePlayer()
}

// TestUpdateMyProfile_Success ...
func (suite *UpdateMyProfile) TestUpdateMyProfile_Success() {
	var (
		token    string
		response util.Response

		playerDao = dao.Player{}

		// body update only player name
		body = model.Player{
			Name: "update name",
		}
	)

	data := map[string]interface{}{
		"id": testhelper.PlayerObjID,
	}

	token, err := util.GenerateUserToken(data)
	if err != nil {
		panic(err)
	}

	bodyMarshal, _ := json.Marshal(body)

	// request
	req := httptest.NewRequest(http.MethodPut, "/api/me", bytes.NewReader(bodyMarshal))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testhelper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	updatedPlayer, err := playerDao.FindByID(testhelper.PlayerObjID)
	if err != nil {
		panic(err)
	}

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), testhelper.PlayerIDString, response["data"])
	assert.Equal(suite.T(), "success", response["message"])
	assert.Equal(suite.T(), body.Name, updatedPlayer.Name)
}

// TestPlayer ...
func TestPlayer(t *testing.T) {
	suite.Run(t, new(MyProfileSuite))
	suite.Run(t, new(UpdateMyProfile))
}
