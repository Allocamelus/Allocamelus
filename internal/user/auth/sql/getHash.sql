SELECT authKeyHash
FROM UserKeys
WHERE userId = ?
ORDER BY userKeyId DESC
LIMIT 1