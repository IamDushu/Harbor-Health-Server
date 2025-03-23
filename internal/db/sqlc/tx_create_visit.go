package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateVisitArgs struct {
	ProviderID  uuid.UUID `json:"provider_id"`
	LocationID  uuid.UUID `json:"location_id"`
	MemberID    uuid.UUID `json:"member_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Notes       string    `json:"notes"`
}

func (s *Store) CreateVisitTx(ctx context.Context, visitArgs CreateVisitArgs) (Visit, error) {

	var newVisit Visit

	err := s.execTx(ctx, func(q *Queries) error {
		visitExists, err := q.CheckVisitSlotExists(ctx, CheckVisitSlotExistsParams{
			ProviderID:  visitArgs.ProviderID,
			ScheduledAt: visitArgs.ScheduledAt,
		})
		if err != nil {
			return err
		}

		if visitExists {
			return fmt.Errorf("time slot is already booked")
		}
		newVisit, err = q.CreateVisit(ctx, CreateVisitParams{
			VisitID:     uuid.New(),
			ProviderID:  visitArgs.ProviderID,
			MemberID:    visitArgs.MemberID,
			LocationID:  visitArgs.LocationID,
			ScheduledAt: visitArgs.ScheduledAt,
			Status:      "pending",
			Notes:       visitArgs.Notes,
		})
		return err
	})

	return newVisit, err
}
