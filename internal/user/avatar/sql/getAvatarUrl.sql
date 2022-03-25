SELECT hash
FROM UserAvatars
WHERE userId = ?
  AND active = 1
ORDER BY userAvatarId DESC
LIMIT 1