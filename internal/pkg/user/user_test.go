package user_test

import (
	"encoding/json"
	"myapi/internal/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("12")

	// Assertions
	if assert.NoError(t, user.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		gotUserJSON := user.User{}
		json.Unmarshal(rec.Body.Bytes(), &gotUserJSON) // wantJSON, _ := json.Marshal(user.User{IdUser: 12})
		assert.EqualValues(t, user.User{IdUser: 12}, gotUserJSON)
	}

}
