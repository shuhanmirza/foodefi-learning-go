package fd_auth

import (
	"bytes"
	"context"
	"encoding/hex"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRandomUser(t *testing.T, store *db.Store) db.Users {
	arg := db.CreateUserParams{
		Username: util.RandomUserName(),
		Password: hex.EncodeToString(util.GetHash([]byte("shuhan1234"))),
		Role:     util.RandomRole(),
	}

	user, err := store.Queries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)

	require.NotEmpty(t, user.CreatedAt)

	return user
}
func TestLoginApi(t *testing.T) {
	store := db.NewStore(testDb)
	server := NewServer(store)

	user := createRandomUser(t, store)

	recorder := httptest.NewRecorder()

	url := "/auth/login"
	var jsonData = []byte(`{
		"username": "` + user.Username + `",
		"password": "shuhan1234"
	}`)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}
