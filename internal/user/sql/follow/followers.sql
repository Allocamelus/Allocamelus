SELECT COUNT(userFollowId)
FROM UserFollows
WHERE followUserId = ?
  AND accepted = 1