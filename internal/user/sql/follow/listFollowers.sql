SELECT userId
FROM UserFollows
WHERE followUserId = ?
  AND accepted = 1