package modelext

import (
	"database/sql"
	"errors"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserWithZgroup is a User model with eagerly loaded Zgroup data
type UserWithZgroup struct {
	models.User
	Zgroup *models.Zgroup `json:"zgroup,omitempty"`
}

// EnsureUserWithOauth finds a user with the given provider id or creates them with the associated name
// and email address if they don't exist associates
func EnsureUserWithOauth(c buffalo.Context, provider, providerID, name, email string) (*models.User, error) {
	tx := getTx(c)
	user, err := models.Users(
		qm.InnerJoin("oauth_users ON oauth_users.user_id = users.id"),
		qm.InnerJoin("oauths ON oauth_users.oauth_id = oauths.id"),
		models.OauthWhere.Provider.EQ(provider),
		models.OauthWhere.ProviderID.EQ(providerID),
	).One(c, tx)
	if user != nil {
		return user, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	oauthProvider := &models.Oauth{
		Provider:   provider,
		ProviderID: providerID,
	}
	user = &models.User{
		Name:  name,
		Email: email,
	}
	if err = oauthProvider.Upsert(c, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		return nil, err
	}
	if err = user.Insert(c, tx, boil.Infer()); err != nil {
		return nil, err
	}
	if err = user.AddOauths(c, tx, false, oauthProvider); err != nil {
		return nil, err
	}
	return user, nil
}

// PaginatedUsers adds pagination to a multi-user query
func PaginatedUsers(c buffalo.Context, queries ...qm.QueryMod) ([]*UserWithZgroup, *Pagination, error) {
	pagination, clauses := getPagination(c)
	clauses = append(clauses, qm.Load(models.UserRels.Zgroups, models.ZgroupWhere.Archived.EQ(false)))
	clauses = append(clauses, queries...)
	users, err := models.Users(
		clauses...,
	).All(c, getTx(c))
	if err != nil {
		return nil, nil, err
	}
	pagination.Cursor = -1
	if len(users) > 0 {
		pagination.Cursor = users[len(users)-1].ID
	}
	enrichedUsers := []*UserWithZgroup{}
	for _, user := range users {
		enrichedUsers = append(enrichedUsers, &UserWithZgroup{
			User:   *user,
			Zgroup: ZgroupForUser(user),
		})
	}
	return enrichedUsers, pagination, nil
}

// FindUser finds the profile for the given user
func FindUser(c buffalo.Context, id int) (*UserWithZgroup, error) {
	user, err := models.Users(
		models.UserWhere.ID.EQ(id),
		qm.Load(models.UserRels.Zgroups, models.ZgroupWhere.Archived.EQ(false)),
	).One(c, getTx(c))
	if err != nil || user == nil {
		return nil, err
	}
	return &UserWithZgroup{
		User:   *user,
		Zgroup: ZgroupForUser(user),
	}, nil
}

// UserStatus returns a human readable status for the user
func UserStatus(blocked bool) string {
	if blocked {
		return "Blocked"
	}
	return "Active"
}
