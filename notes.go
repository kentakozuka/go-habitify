package habitify

import (
	"context"
	"fmt"
	"time"
)

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
