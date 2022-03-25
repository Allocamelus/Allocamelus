SELECT followUserId
FROM UserFollows
WHERE userId = ?
  AND accepted = 1