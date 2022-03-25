UPDATE UserAvatars
SET active = 0
WHERE userAvatarId IN (
    SELECT userAvatarId
    FROM (
        SELECT userAvatarId
        FROM UserAvatars
        ORDER BY userAvatarId DESC
        LIMIT 1, 18446744073709551615
      ) tmp
  )
  AND userId = ?
  AND active = 1