package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockdb "simplebank/db/mock"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	userRequest := randomUser()
	hashedPassword, err := util.HashPassword(userRequest.Password)
	require.NoError(t, err)

	userResponse := db.User{
		Username:       userRequest.Username,
		FullName:       userRequest.FullName,
		Email:          userRequest.Email,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
	}

	store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(userResponse, nil).Times(1)

	userInJsonBytes, err := json.Marshal(userRequest)
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userInJsonBytes))
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	server := NewServer(store)
	server.router.ServeHTTP(recorder, request)

	requireBodyMatch(t, userRequest, recorder.Body)
}

func randomUser() CreateUserParams {
	return CreateUserParams{
		Username: gofakeit.Username(),
		FullName: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Name(),
	}
}

func requireBodyMatch(t *testing.T, user CreateUserParams, body *bytes.Buffer) {
	bodyInBytes, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser userResponse
	err = json.Unmarshal(bodyInBytes, &gotUser)
	require.NoError(t, err)

	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.True(t, gotUser.PasswordChangedAt.IsZero())
	require.NotEmpty(t, gotUser.CreatedAt)
}
