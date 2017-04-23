SELECT
  gems.id         AS Id,
  gems.name       AS name,
  downloads.count AS download_count,
  ls.code         AS codeurl,
  ls.home         AS homeurl
FROM gem_downloads downloads
  JOIN gems
    ON downloads.rubygem_id = gems.id
  JOIN linksets ls
    ON ls.rubygem_id = gems.id
WHERE downloads.version_id = 0
      AND gems.id > {{.}}
ORDER BY download_count DESC
