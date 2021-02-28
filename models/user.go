package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gobuffalo/pop"
)

// User defines a person who has signed up at the hub
type User struct {
	ID        int       `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

// EnsureUserWithOAuth returns the user with the associated provider id or creates them and associates
// them with the given email
func EnsureUserWithOAuth(ctx context.Context, provider, providerID, email string) (*User, error) {
	db := ctx.Value("tx").(*pop.Connection)

	// first see if we have someone associated with the given provider
	user := &User{}
	err := db.RawQuery(`
		SELECT users.* FROM users
		JOIN oauth_users ON users.id = oauth_users.user_id
		JOIN oauths ON oauths.id = oauth_users.oauth_id
		WHERE oauths.provider = ? AND oauths.provider_id = ?`, provider, providerID).First(user)
	if err == nil {
		return user, nil
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// create the oauth ids and the user associations
	err = db.RawQuery(`INSERT INTO oauths(provider, provider_id) VALUES(?, ?)`, provider, providerID).Exec()
	if err != nil {
		return nil, err
	}
	// we may already have a user with the same email but a different provider
	err = db.RawQuery(`INSERT INTO users(email) VALUES(?) ON CONFLICT DO NOTHING`, email).Exec()
	if err != nil {
		return nil, err
	}
	err = db.RawQuery(`
		INSERT INTO oauth_users(oauth_id, user_id) VALUES(
			(SELECT id FROM oauths WHERE provider = ? AND provider_id = ?),
			(SELECT id FROM users WHERE email = ?)
		) ON CONFLICT DO NOTHING
	`, provider, providerID, email).Exec()
	if err != nil {
		return nil, err
	}
	err = db.RawQuery(`SELECT * FROM users WHERE email = ?`, email).First(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser returns the user with the given id
func GetUser(ctx context.Context, id int) (*User, error) {
	user := &User{}
	db := ctx.Value("tx").(*pop.Connection)
	err := db.RawQuery(`SELECT * FROM users WHERE id = ?`, id).First(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
