package user

import (
	"database/sql"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
)

var (
	//go:embed sql/follow/following.sql
	qFollowing   string
	preFollowing *sql.Stmt
	//go:embed sql/follow/listFollowing.sql
	qListFollowing   string
	preListFollowing *sql.Stmt
	//go:embed sql/follow/follow.sql
	qFollow   string
	preFollow *sql.Stmt
	//go:embed sql/follow/accept.sql
	qAccept   string
	preAccept *sql.Stmt
	//go:embed sql/follow/acceptAll.sql
	qAcceptAll   string
	preAcceptAll *sql.Stmt
	//go:embed sql/follow/unfollow.sql
	qUnfollow   string
	preUnfollow *sql.Stmt
	//go:embed sql/follow/followers.sql
	qFollowers   string
	preFollowers *sql.Stmt
	//go:embed sql/follow/followers.sql
	qListFollowers   string
	preListFollowers *sql.Stmt
	//go:embed sql/follow/listRequests.sql
	qListRequests   string
	preListRequests *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preFollowing, qFollowing)
	data.PrepareQueuer.Add(&preListFollowing, qListFollowing)
	data.PrepareQueuer.Add(&preFollow, qFollow)
	data.PrepareQueuer.Add(&preAccept, qAccept)
	data.PrepareQueuer.Add(&preAcceptAll, qAcceptAll)
	data.PrepareQueuer.Add(&preUnfollow, qUnfollow)
	data.PrepareQueuer.Add(&preFollowers, qFollowers)
	data.PrepareQueuer.Add(&preListFollowers, qListFollowers)
	data.PrepareQueuer.Add(&preListRequests, qListRequests)
}

// Following is userId following followUserId
func Following(userId, followUserId int64) (follow FollowStruct, err error) {
	if compare.EqualInt64(userId, followUserId) {
		follow.Following = true
		return
	}

	err = preFollowing.QueryRow(userId, followUserId).Scan(&follow.Following)
	if err != sql.ErrNoRows {
		if err != nil {
			return
		}
		if !follow.Following {
			follow.Requested = true
		}
	}
	return follow, nil
}

// ListFollowing return a slice of userId(s) that userId follows
func ListFollowing(userId int64) ([]int64, error) {
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

func FollowExt(userId int64, followUserId int64, accepted bool) error {
	following, err := Following(userId, followUserId)
	if err != nil || following.Following {
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
	if t.Public() {
		return FollowExt(userId, followUserId, true)
	}
	// Has the user being followed requested to follow user
	follow, err := Following(followUserId, userId)
	if err != nil {
		return err
	}
	// if so Accept request
	if follow.Requested {
		return Accept(userId, followUserId)
	}
	// else request to follow
	return FollowExt(userId, followUserId, false)
}

// Accept userId Accept followerUserId request
func Accept(userId, followerUserId int64) error {
	// Is follower following already
	follow, err := Following(followerUserId, userId)
	if err != nil || follow.Following {
		// return if following silently
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
		t, err = GetType(followerUserId)
		if err != nil {
			return err
		}
		// Is follower Private? Follow if so
		if t.Private() {
			return FollowExt(userId, followerUserId, true)
		}
	}
	return nil
}

// Decline userId Decline followerUserId request
func Decline(userId, followerUserId int64) error {
	return unfollowDB(followerUserId, userId)
}

// AcceptAll userId Accept all follow request
// Called when accounts goes from private to public
func AcceptAll(userId int64) error {
	_, err := preAcceptAll.Exec(userId)
	return err
}

// unfollowDB userId unfollow followUserId
func unfollowDB(userId, followUserId int64) error {
	follow, err := Following(userId, followUserId)
	if err != nil || (!follow.Following && !follow.Requested) {
		// return if !following silently
		return err
	}

	_, err = preUnfollow.Exec(userId, followUserId)
	return err
}

func Unfollow(userId, unfollowId int64) error {
	if err := unfollowDB(userId, unfollowId); err != nil {
		return err
	}

	// Unfriend if unfollowId and userId are private
	if t, err := GetType(unfollowId); err != nil || !t.Private() {
		return err
	}
	if t, err := GetType(userId); err != nil || !t.Private() {
		return err
	}

	return unfollowDB(unfollowId, userId)
}

// Followers count userId followers
func Followers(userId int64) (followers int64, err error) {
	err = preFollowers.QueryRow(userId).Scan(&followers)
	return
}

// ListFollows return a slice of userId(s) that follow userId
func ListFollowers(userId int64) ([]int64, error) {
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

// ListRequests return a ordered map of userId(s) that request to follow userId
// TODO limit
func ListRequests(userId int64) (map[int64]int64, error) {
	rows, err := preListRequests.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := map[int64]int64{}
	var (
		i         int64
		requester int64
	)
	for rows.Next() {
		if err = rows.Scan(&requester); err != nil {
			return nil, err
		}
		requests[i] = requester
		i++
	}
	return requests, nil
}
