SELECT
  rg.name as name,
  ls.code as repository,
  ls.home as home,
  ls.created_at as created,
  ls.updated_at as last_update,
  dl.version_id as version,
  dl.count as dl_count

FROM gems rg
  INNER JOIN linksets ls ON rg.id = ls.rubygem_id
  INNER JOIN gem_downloads dl ON rg.id = dl.rubygem_id
LIMIT 10;

