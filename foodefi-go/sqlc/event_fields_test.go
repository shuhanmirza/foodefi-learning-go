package db

import (
	"context"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEventField(t *testing.T) EventFields {

	eventFieldType := util.RandomEvenFieldType()
	eventFieldValue, _ := util.RandomEventFieldValue(eventFieldType)
	users, _ := testQueries.ListUsers(context.Background())

	arg := CreateEventFieldParams{
		EventID:  util.RandomInt(1, 10),
		Name:     util.RandomEventFieldName(),
		Type:     eventFieldType,
		Value:    eventFieldValue,
		Recorder: users[util.RandomInt(0, int64(len(users)))-1].Username, //TODO: clean
	}

	eventField, err := testQueries.CreateEventField(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eventField)

	require.Equal(t, arg.EventID, eventField.EventID)
	require.Equal(t, arg.Name, eventField.Name)
	require.Equal(t, arg.Type, eventField.Type)
	require.Equal(t, arg.Value, eventField.Value)
	require.Equal(t, arg.Recorder, eventField.Recorder)

	require.NotEmpty(t, eventField.ID)

	return eventField
}

func TestCreateEventField(t *testing.T) {
	createRandomEventField(t)
}
