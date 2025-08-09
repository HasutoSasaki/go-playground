package loops

import (
	"fmt"
	"time"
)

type Post struct {
	Title     string
	Content   string
	ExpiredAt time.Time
}

func (p Post) expired() bool {
	return time.Now().After(p.ExpiredAt)
}

func MapDelete() {
	m := map[Post]string{
		{"Post 1", "Content 1", time.Now().Add(-time.Hour)}: "Expired",
		{"Post 2", "Content 2", time.Now().Add(time.Hour)}:  "Active",
	}

	for p := range m {
		if p.expired() {
			delete(m, p)
		}
	}

	fmt.Println("After delete:", m)
}
