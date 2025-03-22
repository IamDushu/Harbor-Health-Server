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