package stores

import (
	"context"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/pkg/fiberutil/session"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jackc/pgx/v5"
)

// DBStore struct
type DBStore struct {
	q *db.Queries
}

func NewDBStore(q *db.Queries) *DBStore {
	return &DBStore{q: q}
}

func (dbS *DBStore) Get(key string) (*session.Session, error) {
	if len(key) <= 0 {
		return &session.Session{}, ErrNilKey
	}

	row, err := dbS.q.GetSession(context.Background(), key)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			logger.Error(err)
		}
		return &session.Session{}, session.ErrNoSession
	}

	if row.Expiration < time.Now().Unix() {
		dbS.Delete(key)
		return &session.Session{}, session.ErrNoSession
	}

	s := new(session.Session)
	_, err = s.UnmarshalMsg(row.Data)
	if err != nil {
		logger.Error(err)
		return &session.Session{}, session.ErrNoSession
	}
	return s, nil
}

func (dbS *DBStore) Set(s *session.Session) error {
	if len(s.Key) <= 0 {
		return ErrNilKey
	}

	data, err := s.MarshalMsg(nil)
	// This Error should not happen
	logger.Error(err)

	exist, err := dbS.q.SessionExist(context.Background(), s.Key)
	if err != nil {
		return err
	}
	if exist {
		return dbS.q.UpdateSession(context.Background(), db.UpdateSessionParams{
			Data:       data,
			Expiration: s.Expires.Unix(),
			Key:        s.Key,
		})
	}

	return dbS.q.InsertSession(context.Background(), db.InsertSessionParams{Key: s.Key, Data: data, Expiration: s.Expires.Unix()})

}
func (dbS *DBStore) Delete(key string) error {
	if len(key) <= 0 {
		return ErrNilKey
	}
	return dbS.q.DeleteSession(context.Background(), key)
}

func (dbS *DBStore) Cleanup() error {
	return dbS.q.DeleteOldSessions(context.Background(), time.Now().Unix())
}
