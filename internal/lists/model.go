package lists

import "time"

type List struct {
	ID        int64
	UserID    int64
	Title     string
	CreatedAt time.Time
}
