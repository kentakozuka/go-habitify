package habitify

import (
	"context"
	"fmt"
	"time"
)

type Habit struct {
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
}

func (c *Client) GetHabit(ctx context.Context, id string) (*Habit, error) {
	var habit Habit
	err := c.get(ctx, fmt.Sprintf("%s/%s", urlHabits, id), &habit)
	if err != nil {
		return nil, err
	}

	return &habit, nil
}

func (c *Client) ListHabits(ctx context.Context) ([]*Habit, error) {
	var habits []*Habit
	err := c.get(ctx, urlHabits, &habits)
	if err != nil {
		return nil, err
	}

	return habits, nil
}
