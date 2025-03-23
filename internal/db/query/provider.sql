-- name: GetProvidersFromLocation :many
SELECT 
    p.provider_id,
    u.first_name,
    u.last_name,
    p.credentials,
    p.specialization
FROM 
    provider_locations pl
JOIN 
    providers p ON pl.provider_id = p.provider_id
JOIN 
    users u ON p.user_id = u.user_id
WHERE 
    pl.location_id = $1
    AND p.is_available = true;

-- name: GetAvailableSlotsForProvider :many
SELECT pa.day_of_week, pa.start_time, pa.end_time
FROM provider_availability pa
WHERE pa.provider_id = $1
  AND pa.day_of_week = $2
  AND NOT EXISTS (
    SELECT 1
    FROM visits v
    WHERE v.provider_id = pa.provider_id
      AND v.scheduled_at::DATE = $3  
      AND v.scheduled_at::TIME = pa.start_time  
  );

-- name: CheckProviderAvailability :one
SELECT EXISTS (
    SELECT 1 FROM provider_availability
    WHERE provider_id = $1 
    AND day_of_week = $2 
    AND start_time = $3
) AS exists;