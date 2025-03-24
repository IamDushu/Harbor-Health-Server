-- name: CheckVisitSlotExists :one
SELECT EXISTS (
    SELECT 1 FROM visits
    WHERE provider_id = $1 AND scheduled_at = $2
) AS exists;

-- name: CreateVisit :one
INSERT INTO visits (visit_id, provider_id, member_id, location_id, scheduled_at, status, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAllPendingVisits :many
SELECT * 
FROM visits
WHERE member_id = $1
  AND status = 'pending';

-- name: GetAllPendingVisitsWithProviderDetails :many
SELECT 
    v.visit_id,
    v.scheduled_at,
    u.image_url AS provider_image_url,
    u.first_name AS provider_firstName,
    u.last_name AS provider_lastName,
    pr.credentials AS provider_credentials
FROM visits v
JOIN providers pr ON v.provider_id = pr.provider_id
JOIN users u ON pr.user_id = u.user_id
WHERE v.member_id = $1
  AND v.status = 'pending';

-- name: GetVisitInfo :one
SELECT 
    v.visit_id,
    l.phone AS location_phone,
    l.latitude,
    l.longitude,
    l.name AS location_name,
    l.address AS location_address,
    l.image_url AS location_image,
    u.image_url AS provider_image,
    u.first_name || ' ' || u.last_name AS provider_name,
    p.credentials AS provider_credentials,
    v.scheduled_at AS visit_time
FROM visits v
JOIN locations l ON v.location_id = l.location_id
JOIN providers p ON v.provider_id = p.provider_id
JOIN users u ON p.user_id = u.user_id
WHERE v.visit_id = $1;
