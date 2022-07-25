package delivery

import "time"

type Comment struct {
	ID         int
	ID_Posting int
	ID_Users   int
	Comment    string
	Created_at time.Time
	Deleted_at time.Time
}
