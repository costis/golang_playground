SELECT
  id,
  name
FROM gems
WHERE id > {{.}} ORDER BY created_at LIMIT 10000