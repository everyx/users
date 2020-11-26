package internal

import (
	"context"
	"time"

	"github.com/adnaan/users/internal/models"
	"github.com/adnaan/users/internal/models/user"
	"github.com/fatih/structs"
	"github.com/google/uuid"
)

type DefaultUserStore struct {
	Driver     string
	DataSource string
	Ctx        context.Context
	Client     *models.Client
}

func (d *DefaultUserStore) getUser(id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return d.Client.User.Get(d.Ctx, uid)
}

func (d *DefaultUserStore) updateUserBuilder(id string) (*models.UserUpdateOne, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return d.Client.User.UpdateOneID(uid), nil
}

func (d *DefaultUserStore) New(email, password string, meta map[string]interface{}) (string, error) {
	usr, err := d.Client.User.Create().
		SetEmail(email).
		SetPassword(password).
		SetMetadata(meta).
		Save(d.Ctx)
	if err != nil {
		return "", err
	}
	return usr.ID.String(), nil
}

func (d *DefaultUserStore) GetUser(id string) (map[string]interface{}, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return nil, err
	}
	usr.Password = ""

	return structs.Map(usr), err
}

func (d *DefaultUserStore) GetUserByEmail(email string) (map[string]interface{}, error) {
	usr, err := d.Client.User.Query().Where(user.Email(email)).Only(d.Ctx)
	if err != nil {
		return nil, err
	}

	return structs.Map(usr), err
}

func (d *DefaultUserStore) GetUserByConfirmationToken(token string) (map[string]interface{}, error) {
	usr, err := d.Client.User.Query().Where(user.ConfirmationToken(token)).Only(d.Ctx)
	if err != nil {
		return nil, err
	}

	return structs.Map(usr), err
}

func (d *DefaultUserStore) GetUserByRecoveryToken(token string) (map[string]interface{}, error) {
	usr, err := d.Client.User.Query().Where(user.RecoveryToken(token)).Only(d.Ctx)
	if err != nil {
		return nil, err
	}

	return structs.Map(usr), err
}

func (d *DefaultUserStore) GetUserByEmailChangeToken(token string) (map[string]interface{}, error) {
	usr, err := d.Client.User.Query().Where(user.EmailChangeToken(token)).Only(d.Ctx)
	if err != nil {
		return nil, err
	}

	return structs.Map(usr), err
}

func (d *DefaultUserStore) DeleteUser(id string) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	return d.Client.User.DeleteOne(usr).Exec(d.Ctx)
}

func (d *DefaultUserStore) UpdatePassword(id, password string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetPassword(password).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) UpdateEmail(id, email string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetEmail(email).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveConfirmationToken(id, token string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetConfirmationToken(token).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveConfirmationTokenSentAt(id string, tokenSentAt time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetConfirmationSentAt(tokenSentAt).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) MarkConfirmed(id string, confirmed bool) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetConfirmed(confirmed).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) DeleteConfirmToken(id string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetNillableConfirmationToken(nil).SetNillableConfirmationSentAt(nil).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveRecoveryToken(id, token string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetRecoveryToken(token).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveRecoveryTokenSentAt(id string, tokenSentAt time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetRecoverySentAt(tokenSentAt).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) DeleteRecoveryToken(id string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetNillableRecoveryToken(nil).SetNillableRecoverySentAt(nil).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveEmailChangeToken(id, email, token string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetEmailChangeToken(token).SetEmailChange(email).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultUserStore) SaveEmailChangeTokenSentAt(id string, tokenSentAt time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetEmailChangeSentAt(tokenSentAt).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultUserStore) DeleteEmailChangeToken(id string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetNillableEmailChangeToken(nil).SetNillableEmailChangeSentAt(nil).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) UpsertMetaData(id string, metaData map[string]interface{}) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	existingMetaData := usr.Metadata
	for k, v := range metaData {
		existingMetaData[k] = v
	}

	_, err = usr.Update().SetMetadata(existingMetaData).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultUserStore) DeleteKeysMetaData(id string, keys []string) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	existingMetaData := usr.Metadata
	for _, k := range keys {
		delete(existingMetaData, k)
	}
	_, err = usr.Update().SetMetadata(existingMetaData).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultUserStore) DeleteAllMetadata(id string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetMetadata(make(map[string]interface{})).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SetUpdatedAt(id string, time time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetUpdatedAt(time).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SetLastSignInAt(id string, time time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetLastSigninAt(time).Save(d.Ctx)
	if err != nil {
		return err
	}
	return nil
}
