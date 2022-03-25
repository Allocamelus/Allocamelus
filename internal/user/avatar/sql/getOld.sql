SELECT UA.fileType,
  UA.hash
FROM UserAvatars UA
WHERE userId = ?
  AND active = 0
  AND NOT EXISTS (
    SELECT *
    FROM UserAvatars
    WHERE UA.hash = hash
      AND active = 1
  )