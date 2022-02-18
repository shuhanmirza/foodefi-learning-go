package db

import (
	"context"
	"database/sql"
	"encoding/hex"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createShuhanAdminUser(t *testing.T) Users { // :3 olosh ami
	arg := CreateUserParams{
		Username: "shuhan_admin",
		Password: "2860087d2a44a6b61ec5a829a8582bdf0580d270d6d44f05493cd245f5e053ca", //shuhan1234
		Role:     "admin",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)

	require.NotEmpty(t, user.CreatedAt)

	return user
}

func createShuhanScraperUser(t *testing.T) Users { // :3 olosh ami
	arg := CreateUserParams{
		Username: "shuhan_scraper",
		Password: "2860087d2a44a6b61ec5a829a8582bdf0580d270d6d44f05493cd245f5e053ca", //shuhan1234
		Role:     "scraper",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)

	require.NotEmpty(t, user.CreatedAt)

	return user
}

func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username: util.RandomUserName(),
		Password: hex.EncodeToString(util.GetHash([]byte("shuhan1234"))),
		Role:     util.RandomRole(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)

	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	//createShuhanAdminUser(t)
	//createShuhanScraperUser(t)
	createRandomUser(t)

}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
