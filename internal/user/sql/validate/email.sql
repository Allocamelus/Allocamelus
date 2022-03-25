SELECT EXISTS(
    SELECT *
    FROM Users
    WHERE email = ?
  )