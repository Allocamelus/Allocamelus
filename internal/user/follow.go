package user

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var preFollowing *sql.Stmt

// Following is userId following followUserId
func Following(userId, followUserId int64) (following bool, err error) {
	if preFollowing == nil {
		preFollowing = g.Data.Prepare(`SELECT EXISTS(SELECT * FROM UserFollows WHERE userId = ? AND followUserId = ? AND accepted = 1)`)
	}
	err = preFollowing.QueryRow(userId, followUserId).Scan(&following)
	return
}

var preListFollowing *sql.Stmt

// ListFollowing return a slice of userId(s) that userId follows
func ListFollowing(userId int64) ([]int64, error) {
	if preListFollowing == nil {
		preListFollowing = g.Data.Prepare(`SELECT followUserId FROM UserFollows WHERE userId = ? AND accepted = 1`)
	}

	rows, err := preListFollowing.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		following    []int64
		followUserId int64
	)
	for rows.Next() {
		if err = rows.Scan(&followUserId); err != nil {
			return nil, err
		}
		following = append(following, followUserId)
	}
	return following, nil
}

var preFollow *sql.Stmt

func FollowExt(userId int64, followUserId int64, accepted bool) error {
	if preFollow == nil {
		preFollow = g.Data.Prepare(`INSERT INTO UserFollows (userId, followUserId, accepted, created) VALUES (?, ?, ?, ?)`)
	}

	following, err := Following(userId, followUserId)
	if err != nil || following {
		// return if following silently
		return err
	}

	_, err = preFollow.Exec(userId, followUserId, accepted, time.Now().Unix())
	return err
}

// Follow userId follow followUserId
func Follow(userId int64, followUserId int64) error {
	t, err := GetType(followUserId)
	if err != nil {
		return err
	}

	// Only Public Users are auto followed
	accepted := t.Public()

	return FollowExt(userId, followUserId, accepted)
}

var preAccept *sql.Stmt

// Accept userId Accept followerUserId request
func Accept(userId, followerUserId int64) error {
	if preAccept == nil {
		preAccept = g.Data.Prepare(`UPDATE UserFollows SET accepted = 1 WHERE userId = ? AND followUserId = ?`)
	}

	// Is follower following already
	following, err := Following(followerUserId, userId)
	if err != nil || !following {
		// return if !following silently
		return err
	}

	_, err = preAccept.Exec(followerUserId, userId)
	if err != nil {
		return err
	}

	t, err := GetType(userId)
	if err != nil {
		return err
	}

	// Auto Follow back is like fb friends
	if t.Private() {
		return FollowExt(userId, followerUserId, true)
	}
	return nil
}

var preAcceptAll *sql.Stmt

// AcceptAll userId Accept follow request
func AcceptAll(userId int64) error {
	if preAcceptAll == nil {
		preAcceptAll = g.Data.Prepare(`UPDATE UserFollows SET accepted = 1 WHERE followUserId = ? AND accepted = 0`)
	}

	_, err := preAcceptAll.Exec(userId)
	return err
}

var preUnfollow *sql.Stmt

// Unfollow userId unfollow followUserId
func Unfollow(userId, followUserId int64) error {
	if preUnfollow == nil {
		preUnfollow = g.Data.Prepare(`DELETE FROM UserFollows WHERE userId = ? AND followUserId = ?`)
	}

	following, err := Following(userId, followUserId)
	if err != nil || !following {
		// return if !following silently
		return err
	}

	_, err = preUnfollow.Exec(userId, followUserId)
	return err
}

var preFollowers *sql.Stmt

// Followers count userId followers
func Followers(userId int64) (followers int64, err error) {
	if preFollowers == nil {
		preFollowers = g.Data.Prepare(`SELECT COUNT(userFollowId) FROM UserFollows WHERE followUserId = ? AND accepted = 1`)
	}
	err = preFollowers.QueryRow(userId).Scan(&followers)
	return
}

var preListFollowers *sql.Stmt

// ListFollowing return a slice of userId(s) that follow userId
func ListFollowers(userId int64) ([]int64, error) {
	if preListFollowers == nil {
		preListFollowers = g.Data.Prepare(`SELECT userId FROM UserFollows WHERE followUserId = ? AND accepted = 1`)
	}

	rows, err := preListFollowing.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		followers []int64
		follower  int64
	)
	for rows.Next() {
		if err = rows.Scan(&follower); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	return followers, nil
}
