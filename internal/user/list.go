package user

type ListUsers map[int64]*User

type List struct {
	Users ListUsers `json:"users"`
}

// AddTo add user to list if not already
func (l *List) AddUser(userID int64, userStruct ...*User) error {
	if l == nil {
		l = new(List)
	}
	if l.Users == nil {
		l.Users = map[int64]*User{}
	}
	if l.Exist(userID) {
		return nil
	}
	if len(userStruct) == 1 {
		l.Users[userID] = userStruct[0]
	} else {
		user, err := GetPublic(userID)
		if err != nil {
			return err
		}
		l.Users[userID] = &user
	}
	return nil
}

// Exist userid in map
func (l *List) Exist(userID int64) bool {
	_, ok := l.Users[userID]
	return ok
}
