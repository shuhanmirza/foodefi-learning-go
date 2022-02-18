package db

import (
	"context"
	"database/sql"
	"fmt"
	"foodefi-go/util"
)

//Store provides all functions to execute db queries and transactions
type Store struct {
	queries *Queries
	db      *sql.DB
}

//NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		queries: New(db),
		db:      db,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v ---- rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type EventFieldTx struct {
	fieldType  string
	fieldName  string
	fieldValue string
}

type SubmitEventsTxParams struct {
	blockchainId int64
	blockNumber  int64
	eventName    string
	recorder     string
	fields       []EventFieldTx
}

type SubmitEventTxResult struct {
	blockchain  Blockchains
	event       Events
	eventFields []EventFields
	recorder    Users
}

// SubmitEventTx inserts event data to events and event_fields table
/* Firstly, it checks if the recorder user is a scraper or not
   Secondly, it checks if the blockchain exists or not. throws error if blockchain does not exist.
   Thirdly, It checks event entry exists or not, then enter new event if it does not exist.
   After that, it inserts new event_fields in the event_fields table
*/
func (store *Store) SubmitEventTx(ctx context.Context, arg SubmitEventsTxParams) (SubmitEventTxResult, error) {
	var result SubmitEventTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		// check recorder validity
		result.recorder, err = queries.GetUser(ctx, arg.recorder)
		if err != nil {
			return err
		}
		if result.recorder.Role != util.UserRoleScraper {
			return &util.UserRoleNotPermitted{}
		}

		// check blockchain validity
		result.blockchain, err = queries.GetBlockchain(ctx, arg.blockchainId)
		if err != nil {
			return err
		}

		// check event validity, insert new event if does not exists
		result.event, err = queries.GetEventByAllData(ctx, GetEventByAllDataParams{
			BlockchainID: arg.blockchainId,
			BlockNumber:  arg.blockNumber,
			EventName:    arg.eventName,
		})
		if err == sql.ErrNoRows {
			result.event, err = queries.CreateEvent(ctx, CreateEventParams{
				BlockchainID: arg.blockchainId,
				BlockNumber:  arg.blockNumber,
				EventName:    arg.eventName,
			})
		} else if err != nil {
			return err
		}

		// insert event field
		var eventFieldList []EventFields
		for _, field := range arg.fields {
			var eventField EventFields
			if util.IsValidEventFieldType(field.fieldType) == false {
				return &util.InvalidEventFieldType{}
			}
			eventField, err = queries.CreateEventField(ctx, CreateEventFieldParams{
				EventID:  result.event.ID,
				Name:     field.fieldName,
				Type:     field.fieldType,
				Value:    field.fieldValue,
				Recorder: arg.recorder,
			})
			if err != nil {
				return err
			}

			eventFieldList = append(eventFieldList, eventField)
		}

		result.eventFields = eventFieldList

		return nil
	})

	return result, err
}
