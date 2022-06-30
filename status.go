package habitify

import (
	"context"
	"fmt"
	"time"
)

const (
	STATUS_INPROGRESS = "in_progress"
	STATUS_COMPLETED  = "completed"
	STATUS_SKIPPED    = "skipped"
	STATUS_FAILED     = "failed"
)

type Status struct {
	Status   string `json:"status"`
	Progress struct {
		CurrentValue  int       `json:"current_value"`
		TargetValue   int       `json:"target_value"`
		UnitType      string    `json:"unit_type"`
		Periodicity   string    `json:"periodicity"`
		ReferenceDate time.Time `json:"reference_date"`
	} `json:"progress"`
}

func (c *Client) GetStatus(ctx context.Context, habitID string, targetDate time.Time) (*Status, error) {
	var status *Status
	err := c.get(ctx, fmt.Sprintf("%s/%s?target_date=%s", urlStatus, habitID, formatTime(targetDate)), &status)
	if err != nil {
		return nil, err
	}

	return status, nil
}
