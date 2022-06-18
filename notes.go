package habitify

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

const dateFormat = "2006-01-02T15:04:05-07:00"

type Note struct {
	ID          string      `json:"id"`
	Content     string      `json:"content"`
	CreatedDate string      `json:"created_date"`
	ImageURL    interface{} `json:"image_url"`
	NoteType    int         `json:"note_type"`
	HabitID     string      `json:"habit_id"`
}

func (c *Client) ListNotes(ctx context.Context, habitID string, from, to time.Time) ([]*Note, error) {
	var notes []*Note
	err := c.get(ctx, fmt.Sprintf("%s/%s?from=%s&to=%s", urlNotes, habitID, formatTime(from), formatTime(to)), &notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func formatTime(t time.Time) string {
	return url.QueryEscape(t.Format(dateFormat))
}
