SELECT meta,
  hash
FROM PostMedia
WHERE postId = ?
  AND active = 1
ORDER BY postMediaId ASC
LIMIT 4