package db

import (
	"context"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestEvent(t *testing.T) Events {

	arg := CreateEventParams{
		EventName:    "ulala",
		BlockchainID: 2,
		BlockNumber:  1000,
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.EventName, event.EventName)
	require.Equal(t, arg.BlockchainID, event.BlockchainID)
	require.Equal(t, arg.BlockNumber, event.BlockNumber)
	require.NotEmpty(t, event.ID)

	return event
}

func createRandomEvent(t *testing.T) Events {

	blockchain := createRandomBlockchain(t)

	arg := CreateEventParams{
		EventName:    util.RandomEventName(),
		BlockchainID: blockchain.ID,
		BlockNumber:  util.RandomInt(0, 10000),
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.EventName, event.EventName)
	require.Equal(t, arg.BlockchainID, event.BlockchainID)
	require.Equal(t, arg.BlockNumber, event.BlockNumber)
	require.NotEmpty(t, event.ID)

	return event
}

func TestCreateEvent(t *testing.T) {
	//createTestEvent(t)
	createRandomEvent(t)
}
