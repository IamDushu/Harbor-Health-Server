// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CheckProviderAvailability(ctx context.Context, arg CheckProviderAvailabilityParams) (bool, error)
	CheckVisitSlotExists(ctx context.Context, arg CheckVisitSlotExistsParams) (bool, error)
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyRecord(ctx context.Context, arg CreateVerifyRecordParams) (EmailVerification, error)
	CreateVisit(ctx context.Context, arg CreateVisitParams) (Visit, error)
	GetAllPendingVisits(ctx context.Context, memberID uuid.UUID) ([]Visit, error)
	GetAllPendingVisitsWithProviderDetails(ctx context.Context, memberID uuid.UUID) ([]GetAllPendingVisitsWithProviderDetailsRow, error)
	GetAvailableSlotsForProvider(ctx context.Context, arg GetAvailableSlotsForProviderParams) ([]GetAvailableSlotsForProviderRow, error)
	GetLocations(ctx context.Context) ([]Location, error)
	GetMember(ctx context.Context, userID uuid.UUID) (Member, error)
	GetProvidersFromLocation(ctx context.Context, locationID uuid.UUID) ([]GetProvidersFromLocationRow, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (Session, error)
	GetUser(ctx context.Context, email string) (User, error)
	GetVerifyRecord(ctx context.Context, arg GetVerifyRecordParams) (EmailVerification, error)
	GetVerifyRecordOnToken(ctx context.Context, token string) (EmailVerification, error)
	GetVisitInfo(ctx context.Context, visitID uuid.UUID) (GetVisitInfoRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyRecordAttempt(ctx context.Context, verificationID uuid.UUID) (EmailVerification, error)
	UpdateVerifyRecordInvalid(ctx context.Context, verificationID uuid.UUID) (EmailVerification, error)
}

var _ Querier = (*Queries)(nil)
