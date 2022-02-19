package db

import (
	"context"
	"database/sql"
	"fmt"
	"foodefi-go/util"
)

//Store provides all functions to execute Db Queries and transactions
type Store struct {
	Queries *Queries
	Db      *sql.DB
}

//NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		Db:      db,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.Db.BeginTx(ctx, nil)
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
	Type  string
	Name  string
	Value string
}

type SubmitEventsTxParams struct {
	BlockchainId int64
	BlockNumber  int64
	EventName    string
	Recorder     string
	Fields       []EventFieldTx
}

type SubmitEventTxResult struct {
	Blockchain  Blockchains
	Event       Events
	EventFields []EventFields
	Recorder    Users
}

// SubmitEventTx inserts Event data to events and event_fields table
/* Firstly, it checks if the Recorder user is a scraper or not
   Secondly, It checks Event entry exists or not, then enter new Event if it does not exist.
   After that, it inserts new event_fields in the event_fields table
*/
func (store *Store) SubmitEventTx(ctx context.Context, arg SubmitEventsTxParams) (SubmitEventTxResult, error) {
	var result SubmitEventTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		// check Recorder validity
		result.Recorder, err = queries.GetUser(ctx, arg.Recorder)
		if err != nil {
			return err
		}
		if result.Recorder.Role != util.UserRoleScraper {
			return &util.UserRoleNotPermitted{}
		}

		// check Event validity, insert new Event if it does not exist
		result.Event, err = queries.GetEventByAllData(ctx, GetEventByAllDataParams{
			BlockchainID: arg.BlockchainId,
			BlockNumber:  arg.BlockNumber,
			EventName:    arg.EventName,
		})
		if err == sql.ErrNoRows {
			result.Event, err = queries.CreateEvent(ctx, CreateEventParams{
				BlockchainID: arg.BlockchainId,
				BlockNumber:  arg.BlockNumber,
				EventName:    arg.EventName,
			})
		} else if err != nil {
			return err
		}

		// insert Event field
		var eventFieldList []EventFields
		for _, field := range arg.Fields {
			var eventField EventFields
			if util.IsValidEventFieldType(field.Type) == false {
				return &util.InvalidEventFieldType{}
			}
			eventField, err = queries.CreateEventField(ctx, CreateEventFieldParams{
				EventID:  result.Event.ID,
				Name:     field.Name,
				Type:     field.Type,
				Value:    field.Value,
				Recorder: arg.Recorder,
			})
			if err != nil {
				return err
			}

			eventFieldList = append(eventFieldList, eventField)
		}

		result.EventFields = eventFieldList

		return nil
	})

	return result, err
}
