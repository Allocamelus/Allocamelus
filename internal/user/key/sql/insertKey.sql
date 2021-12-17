INSERT INTO UserKeys (
    userId,
    created,
    authKeyHash,
    authKeySalt,
    publicArmored,
    privateArmored,
    recoveryKeyHash,
    recoveryArmored
  )
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
  )