SELECT hash
FROM PostMediaFiles
WHERE hash = ?
  OR newHash = ?
LIMIT 1