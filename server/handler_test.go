package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hamza72x/go-gin-gorm/accounts"
	"github.com/hamza72x/go-gin-gorm/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAdminCreateUser(t *testing.T) {
	tt := []struct {
		name          string
		buildRequest  func() *http.Request
		checkResponse func(*testing.T, *http.Response)
	}{
		{
			name: "valid create user",
			buildRequest: func() *http.Request {
				body := util.MustBytesReaderFromMap(gin.H{
					"name": "test",
				})

				req, err := http.NewRequest(http.MethodPost, "/admin/create-account", body)
				require.NoError(t, err)

				return req
			},
			checkResponse: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)
				defer resp.Body.Close()
				acc := &accounts.Account{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(acc))
				require.Equal(t, "test", acc.Name)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			testServer.router.ServeHTTP(w, tc.buildRequest())
			tc.checkResponse(t, w.Result())
		})
	}
}
