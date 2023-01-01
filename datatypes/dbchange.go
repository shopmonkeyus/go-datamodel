package datatypes

import "encoding/json"

type ChangeEventOperation string

const (
	ChangeEventInsert ChangeEventOperation = "INSERT"
	ChangeEventUpdate ChangeEventOperation = "UPDATE"
	ChangeEventDelete ChangeEventOperation = "DELETE"
)

type ChangeEventHelper interface {
	GetOperation() ChangeEventOperation
	GetBefore() any
	GetAfter() any
}

type ChangeEvent[T any] struct {
	ID            string               `json:"id"`
	Timestamp     int64                `json:"timestamp"`
	MvccTimestamp string               `json:"mvccTimestamp"`
	Table         string               `json:"table"`
	Key           []string             `json:"key"`
	LocationID    *string              `json:"locationId,omitempty"`
	CompanyID     *string              `json:"companyId,omitempty"`
	UserID        *string              `json:"userId,omitempty"`
	SessionID     *string              `json:"sessionId,omitempty"`
	Version       int64                `json:"version"`
	Region        string               `json:"region"`
	Operation     ChangeEventOperation `json:"operation"`
	Before        *T                   `json:"before,omitempty"`
	After         *T                   `json:"after,omitempty"`
	Diff          []string             `json:"diff,omitempty"`
}

// String returns a JSON stringified version of the ChangeEvent
func (c ChangeEvent[T]) String() string {
	buf, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func (c ChangeEvent[T]) GetOperation() ChangeEventOperation {
	return c.Operation
}

func (c ChangeEvent[T]) GetBefore() any {
	return c.Before
}

func (c ChangeEvent[T]) GetAfter() any {
	return c.After
}
