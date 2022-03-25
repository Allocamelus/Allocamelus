SELECT COUNT(userEventId)
FROM UserEvents
WHERE eventType = ?
  AND userId = ?
  AND created > ?