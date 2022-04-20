SELECT PM.alt,
  PMF.width,
  PMF.height,
  PMF.hash
FROM PostMedia PM
  JOIN PostMediaFiles PMF ON (PM.postMediaFileId = PMF.postMediaFileId)
WHERE postId = ?
  AND active = 1
ORDER BY postMediaId ASC
LIMIT 4