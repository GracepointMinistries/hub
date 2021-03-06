package settings

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"sync"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const settingsSpecialID = 0x1234

// Serializable ensures that settings implement the json.Marshaler and json.Unmarshaler interfaces
type Serializable interface {
	json.Marshaler
	json.Unmarshaler
}

type wrappedSettings struct {
	*models.Setting
	mutex sync.RWMutex
}

var globalSettings = &wrappedSettings{}

func (s *wrappedSettings) MarshalJSON() ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return json.Marshal(s.Setting)
}

func (s *wrappedSettings) UnmarshalJSON(data []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	inner := &models.Setting{}
	if err := json.Unmarshal(data, inner); err != nil {
		return err
	}
	inner.ID = settingsSpecialID
	s.Setting = inner
	return nil
}

// SerializableSettings returns a reference to the settings
// that can be used (only) for serialization -- the serialization
// functions are already guarded by a mutex, so this should be ok
// to pass around.
func SerializableSettings() Serializable {
	return globalSettings
}

// this should only ever be called while the global settings
// mutex is write locked
func syncSettings(c buffalo.Context) error {
	_, err := globalSettings.Setting.Update(c, modelext.GetTx(c), boil.Infer())
	return err
}

// UpdateSheet updates the sheet global settings
func UpdateSheet(c buffalo.Context, sheet string) error {
	globalSettings.mutex.Lock()
	defer globalSettings.mutex.Unlock()

	globalSettings.Sheet = sheet
	return syncSettings(c)
}

// Sheet returns the current sheet settings
func Sheet() string {
	globalSettings.mutex.RLock()
	defer globalSettings.mutex.RUnlock()

	return globalSettings.Sheet
}

// UpdateScript updates the script global settings
func UpdateScript(c buffalo.Context, script string) error {
	globalSettings.mutex.Lock()
	defer globalSettings.mutex.Unlock()

	globalSettings.Script = script
	return syncSettings(c)
}

// Script returns the current script settings
func Script() string {
	globalSettings.mutex.RLock()
	defer globalSettings.mutex.RUnlock()

	return globalSettings.Script
}

// HasSheet returns whether the current sheet settings are set
func HasSheet() bool {
	globalSettings.mutex.RLock()
	defer globalSettings.mutex.RUnlock()

	return globalSettings.Sheet != ""
}

// Initialize should only be called at the start of the program
func Initialize() error {
	// create or retrieve the global settings singleton from the database
	return modelext.DB.Transaction(func(c *pop.Connection) error {
		settings, err := models.FindSetting(context.Background(), c.TX, settingsSpecialID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if settings != nil {
			globalSettings.Setting = settings
			return nil
		}
		settings = &models.Setting{
			ID: settingsSpecialID,
		}
		err = settings.Insert(context.Background(), c.TX, boil.Infer())
		if err != nil {
			return err
		}
		globalSettings.Setting = settings
		return nil
	})
}
