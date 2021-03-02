package actions

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/markbates/goth/gothic"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func getAdmins() []string {
	return strings.Split(envy.Get("ADMINS", ""), ",")
}

func adminCallback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	log.Println(admins)
	for _, admin := range admins {
		if admin == oauthUser.Email {
			c.Session().Set("Admin", admin)
			return c.Redirect(http.StatusSeeOther, "/admin")
		}
	}
	return c.Error(http.StatusUnauthorized, fmt.Errorf("%s is not an admin", oauthUser.Email))
}

func adminLogoutHandler(c buffalo.Context) error {
	c.Session().Clear()
	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

func requireAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("Admin") == nil {
			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}
		c.Set("admin", c.Session().Get("Admin"))

		return next(c)
	}
}

func ensureUserWithOauth(c buffalo.Context, provider, providerID, name, email string) (*models.User, error) {
	tx := getTx(c)
	user, err := models.Users(models.UserWhere.Email.EQ(email)).One(c, tx)
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
	if err = user.AddOauths(c, tx, true, oauthProvider); err != nil {
		return nil, err
	}
	return user, nil
}

func authCallback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	user, err := ensureUserWithOauth(c, c.Param("provider"), oauthUser.UserID, oauthUser.Name, oauthUser.Email)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}

// requireLoggedInUser checks whether or not a user is logged in with an unexpired session cookie
// if the user is not logged in, then the frontend should redirect to /auth/google
func requireLoggedInUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("ID") == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		c.Set("id", c.Session().Get("ID"))
		c.Set("email", c.Session().Get("Email"))

		return next(c)
	}
}

func logoutHandler(c buffalo.Context) error {
	c.Session().Delete("ID")
	c.Session().Delete("Email")
	if c.Session().Get("Admin") != nil {
		return c.Redirect(http.StatusSeeOther, "/admin/users")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
