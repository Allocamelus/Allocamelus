package follow

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var preFollowing *sql.Stmt

// Following is userId following followUserId
func Following(userId, followUserId int64) (following bool, err error) {
	if preFollowing == nil {
		preFollowing = g.Data.Prepare(`SELECT EXISTS(SELECT * FROM UserFollows WHERE userId = ? AND followUserId = ?)`)
	}
	err = preFollowing.QueryRow(userId, followUserId).Scan(following)
	return
}

var preListFollowing *sql.Stmt

// ListFollowing return a slice of userId(s) that userId follows
func ListFollowing(userId int64) ([]int64, error) {
	if preListFollowing == nil {
		preListFollowing = g.Data.Prepare(`SELECT followUserId FROM UserFollows WHERE userId = ?`)
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

// Follow userId follow followUserId
func Follow(userId, followUserId int64) error {
	if preFollow == nil {
		preFollow = g.Data.Prepare(`INSERT INTO UserFollows (userId, followUserId, created) VALUES (?, ?, ?)`)
	}

	following, err := Following(userId, followUserId)
	if err != nil || following {
		// return if following silently
		return err
	}

	_, err = preFollow.Exec(userId, followUserId, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
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
	if err != nil {
		return err
	}

	return nil
}

var preFollowers *sql.Stmt

// Followers count userId followers
func Followers(userId int64) (followers int64, err error) {
	if preFollowers == nil {
		preFollowers = g.Data.Prepare(`SELECT COUNT(userFollowId) FROM UserFollows WHERE followUserId = ?`)
	}
	err = preFollowers.QueryRow(userId).Scan(followers)
	return
}

var preListFollowers *sql.Stmt

// ListFollowing return a slice of userId(s) that follow userId
func ListFollowers(userId int64) ([]int64, error) {
	if preListFollowers == nil {
		preListFollowers = g.Data.Prepare(`SELECT userId FROM UserFollows WHERE followUserId = ?`)
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
