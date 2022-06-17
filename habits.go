package habitify

import (
	"context"
	"fmt"
	"time"
)

type getHabitResp struct {
	Message string `json:"message"`
	Habit   *Habit `json:"data"`
	Version string `json:"version"`
	Status  bool   `json:"status"`
}

type listHabitsResp struct {
	Message string   `json:"message"`
	Habits  []*Habit `json:"data"`
	Version string   `json:"version"`
	Status  bool     `json:"status"`
}

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
	resp, err := c.get(ctx, fmt.Sprintf("%s/%s", urlHabits, id))
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	var response getHabitResp
	if err := resp.DecodeJSON(&response); err != nil {
		return nil, err
	}

	return response.Habit, nil
}

func (c *Client) ListHabits(ctx context.Context) ([]*Habit, error) {
	resp, err := c.get(ctx, urlHabits)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	var response listHabitsResp
	if err := resp.DecodeJSON(&response); err != nil {
		return nil, err
	}

	return response.Habits, nil
}
