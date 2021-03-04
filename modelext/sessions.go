package modelext

import (
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// DeleteUserSession deletes the session of a user if it exists
func DeleteUserSession(c buffalo.Context, id int) error {
	_, err := models.Sessions(
		models.SessionWhere.ID.EQ(id),
	).DeleteAll(c, GetTx(c))
	return err
}

// DeleteAdminSession deletes the session of an admin if it exists
func DeleteAdminSession(c buffalo.Context, id int) error {
	_, err := models.AdminSessions(
		models.AdminSessionWhere.ID.EQ(id),
	).DeleteAll(c, GetTx(c))
	return err
}

// ValidateUserSession validates the session for the user is legitimate
func ValidateUserSession(c buffalo.Context, userID, id int) (bool, error) {
	return models.Sessions(
		models.SessionWhere.ID.EQ(id),
		models.SessionWhere.UserID.EQ(userID),
	).Exists(c, GetTx(c))
}

// ValidateAdminSession validates the session for the admin is legitimate
func ValidateAdminSession(c buffalo.Context, admin string, id int) (bool, error) {
	return models.AdminSessions(
		models.AdminSessionWhere.ID.EQ(id),
		models.AdminSessionWhere.Email.EQ(admin),
	).Exists(c, GetTx(c))
}

// CreateUserSession creates a session for the given user
func CreateUserSession(c buffalo.Context, user *models.User, ip string) (*models.Session, error) {
	session := &models.Session{
		IP: ip,
	}
	err := user.AddSessions(c, GetTx(c), true, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// CreateAdminSession creates a session for the given admin
func CreateAdminSession(c buffalo.Context, admin, ip string) (*models.AdminSession, error) {
	session := &models.AdminSession{
		Email: admin,
		IP:    ip,
	}
	err := session.Insert(c, GetTx(c), boil.Infer())
	if err != nil {
		return nil, err
	}
	return session, nil
}
