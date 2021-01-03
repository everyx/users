package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/adnaan/users/internal/models/workspaceinvitation"

	"github.com/adnaan/users/internal/models"
	"github.com/adnaan/users/internal/models/user"
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
	return d.Client.User.UpdateOneID(uid).SetUpdatedAt(time.Now()), nil
}

func (d *DefaultUserStore) New(email, password, role, provider string, meta map[string]interface{}, sendMailFunc func(string, string) error) (string, error) {

	var wrkspace *models.Workspace
	var usr *models.User
	var err error

	if err := withTx(d.Ctx, d.Client, func(tx *models.Tx) error {
		// create user
		createUser := tx.User.Create().
			SetEmail(email).
			SetPassword(password).
			SetRoles([]string{role}).
			SetProvider(provider).
			SetMetadata(meta)

		// send confirmation email
		if sendMailFunc != nil {
			confirmationToken := uuid.New().String()
			err = sendMailFunc(confirmationToken, email)
			if err != nil {
				return fmt.Errorf("err sending confirmation email %v", err)
			}

			createUser = createUser.SetConfirmationToken(confirmationToken).SetConfirmationSentAt(time.Now())

		} else {
			if provider == "email_signup" {
				return fmt.Errorf("provider is email but sendEmailFunc is nil")
			}
			createUser = createUser.SetConfirmed(true)
		}

		// create user
		usr, err = createUser.Save(d.Ctx)
		if err != nil {
			return err
		}

		var name string
		nameVal, ok := meta["name"]
		if ok {
			nameStr, ok := nameVal.(string)
			if ok {
				name = nameStr
			}
		}

		// create default personal workspace
		wrkspace, err = tx.Workspace.Create().
			SetName(fmt.Sprintf("%s's Personal", name)).
			SetDescription(fmt.Sprintf("%s Personal Plan", email)).
			SetOwner(usr).
			SetPlan("default").
			Save(d.Ctx)
		if err != nil {
			return err
		}
		// create role for personal workspace
		_, err = tx.WorkspaceRole.Create().
			SetName("owner").
			SetWorkspaceID(wrkspace.ID).
			SetWorkspaces(wrkspace).
			SetUserID(usr.ID).
			SetUsers(usr).
			Save(d.Ctx)
		if err != nil {
			return err
		}

		invitations, err := tx.WorkspaceInvitation.Query().Where(workspaceinvitation.Email(email)).All(d.Ctx)
		if err == nil {
			for _, invitation := range invitations {
				_, err = d.Client.WorkspaceRole.Create().
					SetName(invitation.Role).
					SetUserID(usr.ID).
					SetWorkspaceID(invitation.WorkspaceID).
					Save(d.Ctx)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserData(id string) (string, string, string, map[string]interface{}, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return "", "", "", nil, err
	}

	return usr.Email, usr.BillingID, usr.APIKey, usr.Metadata, nil
}

func (d *DefaultUserStore) UserIDByEmail(email string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.Email(email)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserIDByConfirmationToken(token string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.ConfirmationToken(token)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserIDByRecoveryToken(token string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.RecoveryToken(token)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserIDByEmailChangeToken(token string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.EmailChangeToken(token)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserIDByOTP(otp string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.Otp(otp)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) UserIDByAPIKey(apiKey string) (string, error) {
	usr, err := d.Client.User.Query().Where(user.APIKey(apiKey)).Only(d.Ctx)
	if err != nil {
		return "", err
	}

	return usr.ID.String(), nil
}

func (d *DefaultUserStore) GetPassword(id string) (string, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return "", err
	}

	return usr.Password, nil
}

func (d *DefaultUserStore) GetAPIKey(id string) (string, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return "", err
	}

	return usr.APIKey, nil
}

func (d *DefaultUserStore) GetEmailChange(id string) (string, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return "", err
	}

	return usr.EmailChange, nil
}

func (d *DefaultUserStore) GetRoles(id string) ([]string, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return nil, err
	}

	return usr.Roles, nil
}

func (d *DefaultUserStore) IsEmailConfirmed(id string) (bool, error) {
	usr, err := d.getUser(id)
	if err != nil {
		return false, err
	}

	return usr.Confirmed, nil
}

func (d *DefaultUserStore) DeleteUser(id string) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	return d.Client.User.DeleteOne(usr).Exec(d.Ctx)
}

func (d *DefaultUserStore) AddRole(id, role string) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	var roles []string
	if usr.Roles == nil {
		roles = usr.Roles
	}
	roles = append(roles, role)
	usrUpdate, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = usrUpdate.SetRoles(roles).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) DeleteRole(id, role string) error {
	usr, err := d.getUser(id)
	if err != nil {
		return err
	}
	if usr.Roles == nil || len(usr.Roles) == 0 {
		return nil
	}
	// filter out the role
	var roles []string
	for _, r := range usr.Roles {
		if r == role {
			continue
		}
		roles = append(roles, role)
	}

	usrUpdate, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = usrUpdate.SetRoles(roles).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) ClearRoles(id string) error {

	usrUpdate, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = usrUpdate.ClearRoles().Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
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
func (d *DefaultUserStore) UpdateAPIKey(id, apiKey string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetAPIKey(apiKey).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) UpdateBillingID(id, billingID string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetBillingID(billingID).Save(d.Ctx)
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

func (d *DefaultUserStore) UpdateProvider(id, provider string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetProvider(provider).Save(d.Ctx)
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
	_, err = u.ClearConfirmationToken().ClearConfirmationSentAt().Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveOTP(id, otp string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetOtp(otp).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) SaveOTPSentAt(id string, otpSentAt time.Time) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.SetOtpSentAt(otpSentAt).Save(d.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultUserStore) DeleteOTP(id string) error {
	u, err := d.updateUserBuilder(id)
	if err != nil {
		return err
	}
	_, err = u.ClearOtp().ClearOtpSentAt().Save(d.Ctx)
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
	_, err = u.ClearRecoveryToken().ClearRecoverySentAt().Save(d.Ctx)
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

	_, err = u.ClearEmailChange().ClearEmailChangeToken().ClearEmailChangeSentAt().Save(d.Ctx)
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

func (d *DefaultUserStore) Close() error {
	return d.Client.Close()
}
