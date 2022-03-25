UPDATE UserFollows
SET accepted = 1
WHERE userId = ?
  AND followUserId = ?