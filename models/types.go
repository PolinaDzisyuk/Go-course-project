package models

import "time"

type Note struct {
	ID      int
	Title   string
	Content string
	Tags    []string
	Date    time.Time
}
