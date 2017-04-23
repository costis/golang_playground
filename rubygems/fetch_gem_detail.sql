SELECT
  g.id   AS id,
  g.name AS name,
  ls.home   AS home,
  ls.code   AS code,
  v.number  AS version,
  v.created_at
FROM
  gems g
  INNER JOIN

  versions v ON v.rubygem_id = g.id
  INNER JOIN linksets ls ON g.id = ls.rubygem_id

WHERE g.name = '{{ . }}'

ORDER BY v.created_at DESC
LIMIT 1;
