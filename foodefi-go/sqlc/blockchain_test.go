package db

import (
	"context"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestBlockchain(t *testing.T) Blockchains { // :3 olosh ami
	name := "Alpine"

	blockchain, err := testQueries.CreateBlockchain(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, blockchain)

	require.Equal(t, name, blockchain.Name)
	require.NotEmpty(t, blockchain.ID)
	require.NotEmpty(t, blockchain.CreatedAt)

	return blockchain
}

func createRandomBlockchain(t *testing.T) Blockchains {
	name := util.RandomBlockchainName()

	blockchain, err := testQueries.CreateBlockchain(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, blockchain)

	require.Equal(t, name, blockchain.Name)
	require.NotEmpty(t, blockchain.ID)
	require.NotEmpty(t, blockchain.CreatedAt)

	return blockchain
}

func TestCreateBlockchain(t *testing.T) {
	//createTestBlockchain(t)
	createRandomBlockchain(t)
}
