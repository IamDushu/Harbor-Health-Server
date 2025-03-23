-- name: CheckVisitSlotExists :one
SELECT EXISTS (
    SELECT 1 FROM visits
    WHERE provider_id = $1 AND scheduled_at = $2
) AS exists;

-- name: CreateVisit :one
INSERT INTO visits (visit_id, provider_id, member_id, location_id, scheduled_at, status, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;