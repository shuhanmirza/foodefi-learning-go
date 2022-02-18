package db

import (
	"context"
	"foodefi-go/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubmitEventTx(t *testing.T) {
	store := NewStore(testDb)

	users, _ := testQueries.ListUsersFilterRole(context.Background(), util.UserRoleScraper)
	for len(users) == 0 {
		createRandomUser(t)
		users, _ = testQueries.ListUsersFilterRole(context.Background(), util.UserRoleScraper)
	}
	event := createRandomEvent(t)

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			var eventFieldList []EventFieldTx

			r := util.RandomInt(1, 10)
			for j := int64(0); j < r; j++ {
				eventFieldName := util.RandomEventFieldName()
				eventFieldType := util.RandomEvenFieldType()
				eventFieldValue, _ := util.RandomEventFieldValue(eventFieldType)

				eventField := EventFieldTx{
					fieldType:  eventFieldType,
					fieldName:  eventFieldName,
					fieldValue: eventFieldValue,
				}
				eventFieldList = append(eventFieldList, eventField)
			}

			result, err := store.SubmitEventTx(context.Background(), SubmitEventsTxParams{
				blockchainId: event.BlockchainID,
				blockNumber:  event.BlockNumber,
				eventName:    event.EventName,
				recorder:     users[util.RandomInt(0, int64(len(users)-1))].Username,
				fields:       eventFieldList,
			})

			require.NoError(t, err)
			require.NotEmpty(t, result)

			for _, resultEventField := range result.eventFields {
				eventField, _ := testQueries.GetEventFieldById(context.Background(), resultEventField.EventID)
				require.NotEmpty(t, eventField)
			}

		}()
	}
}
