SELECT userKeyId,
  publicArmored
FROM UserKeys
WHERE userId = ?
  AND (
    replaced > ?
    OR replaced = 0
  )
ORDER BY userKeyId DESC