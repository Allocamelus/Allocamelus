UPDATE UserFollows
SET accepted = 1
WHERE followUserId = ?
  AND accepted = 0