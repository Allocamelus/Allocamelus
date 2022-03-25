SELECT EXISTS(
    SELECT *
    FROM UserAvatars
    WHERE userId = ?
      AND active = 1
  )