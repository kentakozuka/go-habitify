package habitify

import (
	"context"
	"fmt"
	"time"
)

type Journal struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	IsArchived bool      `json:"is_archived"`
	StartDate  time.Time `json:"start_date"`
	TimeOfDay  []string  `json:"time_of_day"`
	Goal       struct {
		UnitType    string `json:"unit_type"`
		Value       int    `json:"value"`
		Periodicity string `json:"periodicity"`
	} `json:"goal"`
	GoalHistoryItems []struct {
		UnitType    string `json:"unit_type"`
		Value       int    `json:"value"`
		Periodicity string `json:"periodicity"`
	} `json:"goal_history_items"`
	LogMethod  string   `json:"log_method"`
	Recurrence string   `json:"recurrence"`
	Remind     []string `json:"remind"`
	Area       struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Priority string `json:"priority"`
	} `json:"area"`
	CreatedDate time.Time `json:"created_date"`
	Priority    float64   `json:"priority"`
	Status      string    `json:"status"`
	Progress    struct {
		CurrentValue  int       `json:"current_value"`
		TargetValue   int       `json:"target_value"`
		UnitType      string    `json:"unit_type"`
		Periodicity   string    `json:"periodicity"`
		ReferenceDate time.Time `json:"reference_date"`
	} `json:"progress"`
}

func (c *Client) ListJournals(ctx context.Context, targetDate time.Time) ([]*Journal, error) {
	var journals []*Journal
	err := c.get(ctx, fmt.Sprintf("%s?target_date=%s", urlJournal, formatTime(targetDate)), &journals)
	if err != nil {
		return nil, err
	}

	return journals, nil
}
