SELECT accepted
FROM UserFollows
WHERE userId = ?
  AND followUserId = ?