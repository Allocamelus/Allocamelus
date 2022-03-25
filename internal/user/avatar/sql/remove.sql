UPDATE UserAvatars
SET active = 0
WHERE userID = ?
  AND active = 1