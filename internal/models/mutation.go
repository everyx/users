// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/session"
	"github.com/adnaan/users/internal/models/user"
	"github.com/google/uuid"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSession = "Session"
	TypeUser    = "User"
)

// SessionMutation represents an operation that mutate the Sessions
// nodes in the graph.
type SessionMutation struct {
	config
	op            Op
	typ           string
	id            *string
	data          *string
	created_at    *time.Time
	updated_at    *time.Time
	expires_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Session, error)
	predicates    []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows to manage the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for $n.Name.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the id field of the mutation.
func withSessionID(id string) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Session creation.
func (m *SessionMutation) SetID(id string) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *SessionMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetData sets the data field.
func (m *SessionMutation) SetData(s string) {
	m.data = &s
}

// Data returns the data value in the mutation.
func (m *SessionMutation) Data() (r string, exists bool) {
	v := m.data
	if v == nil {
		return
	}
	return *v, true
}

// OldData returns the old data value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldData(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldData is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldData requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldData: %w", err)
	}
	return oldValue.Data, nil
}

// ResetData reset all changes of the "data" field.
func (m *SessionMutation) ResetData() {
	m.data = nil
}

// SetCreatedAt sets the created_at field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *SessionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *SessionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *SessionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetExpiresAt sets the expires_at field.
func (m *SessionMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the expires_at value in the mutation.
func (m *SessionMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old expires_at value of the Session.
// If the Session object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *SessionMutation) OldExpiresAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldExpiresAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ClearExpiresAt clears the value of expires_at.
func (m *SessionMutation) ClearExpiresAt() {
	m.expires_at = nil
	m.clearedFields[session.FieldExpiresAt] = struct{}{}
}

// ExpiresAtCleared returns if the field expires_at was cleared in this mutation.
func (m *SessionMutation) ExpiresAtCleared() bool {
	_, ok := m.clearedFields[session.FieldExpiresAt]
	return ok
}

// ResetExpiresAt reset all changes of the "expires_at" field.
func (m *SessionMutation) ResetExpiresAt() {
	m.expires_at = nil
	delete(m.clearedFields, session.FieldExpiresAt)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.data != nil {
		fields = append(fields, session.FieldData)
	}
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, session.FieldUpdatedAt)
	}
	if m.expires_at != nil {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldData:
		return m.Data()
	case session.FieldCreatedAt:
		return m.CreatedAt()
	case session.FieldUpdatedAt:
		return m.UpdatedAt()
	case session.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldData:
		return m.OldData(ctx)
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case session.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case session.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldData:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetData(v)
		return nil
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case session.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case session.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *SessionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *SessionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(session.FieldExpiresAt) {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	switch name {
	case session.FieldExpiresAt:
		m.ClearExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldData:
		m.ResetData()
		return nil
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case session.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case session.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Session edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op                   Op
	typ                  string
	id                   *uuid.UUID
	email                *string
	password             *string
	confirmed            *bool
	confirmation_sent_at *time.Time
	confirmation_token   *string
	recovery_sent_at     *time.Time
	recovery_token       *string
	otp_sent_at          *time.Time
	otp                  *string
	email_change         *string
	email_change_sent_at *time.Time
	email_change_token   *string
	metadata             *map[string]interface{}
	created_at           *time.Time
	updated_at           *time.Time
	last_signin_at       *time.Time
	clearedFields        map[string]struct{}
	done                 bool
	oldValue             func(context.Context) (*User, error)
	predicates           []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetEmail sets the email field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the email value in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old email value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail reset all changes of the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the password field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the password value in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old password value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword reset all changes of the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetConfirmed sets the confirmed field.
func (m *UserMutation) SetConfirmed(b bool) {
	m.confirmed = &b
}

// Confirmed returns the confirmed value in the mutation.
func (m *UserMutation) Confirmed() (r bool, exists bool) {
	v := m.confirmed
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmed returns the old confirmed value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmed is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmed: %w", err)
	}
	return oldValue.Confirmed, nil
}

// ClearConfirmed clears the value of confirmed.
func (m *UserMutation) ClearConfirmed() {
	m.confirmed = nil
	m.clearedFields[user.FieldConfirmed] = struct{}{}
}

// ConfirmedCleared returns if the field confirmed was cleared in this mutation.
func (m *UserMutation) ConfirmedCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmed]
	return ok
}

// ResetConfirmed reset all changes of the "confirmed" field.
func (m *UserMutation) ResetConfirmed() {
	m.confirmed = nil
	delete(m.clearedFields, user.FieldConfirmed)
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (m *UserMutation) SetConfirmationSentAt(t time.Time) {
	m.confirmation_sent_at = &t
}

// ConfirmationSentAt returns the confirmation_sent_at value in the mutation.
func (m *UserMutation) ConfirmationSentAt() (r time.Time, exists bool) {
	v := m.confirmation_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationSentAt returns the old confirmation_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmationSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmationSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmationSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationSentAt: %w", err)
	}
	return oldValue.ConfirmationSentAt, nil
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (m *UserMutation) ClearConfirmationSentAt() {
	m.confirmation_sent_at = nil
	m.clearedFields[user.FieldConfirmationSentAt] = struct{}{}
}

// ConfirmationSentAtCleared returns if the field confirmation_sent_at was cleared in this mutation.
func (m *UserMutation) ConfirmationSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmationSentAt]
	return ok
}

// ResetConfirmationSentAt reset all changes of the "confirmation_sent_at" field.
func (m *UserMutation) ResetConfirmationSentAt() {
	m.confirmation_sent_at = nil
	delete(m.clearedFields, user.FieldConfirmationSentAt)
}

// SetConfirmationToken sets the confirmation_token field.
func (m *UserMutation) SetConfirmationToken(s string) {
	m.confirmation_token = &s
}

// ConfirmationToken returns the confirmation_token value in the mutation.
func (m *UserMutation) ConfirmationToken() (r string, exists bool) {
	v := m.confirmation_token
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationToken returns the old confirmation_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldConfirmationToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldConfirmationToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldConfirmationToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationToken: %w", err)
	}
	return oldValue.ConfirmationToken, nil
}

// ClearConfirmationToken clears the value of confirmation_token.
func (m *UserMutation) ClearConfirmationToken() {
	m.confirmation_token = nil
	m.clearedFields[user.FieldConfirmationToken] = struct{}{}
}

// ConfirmationTokenCleared returns if the field confirmation_token was cleared in this mutation.
func (m *UserMutation) ConfirmationTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldConfirmationToken]
	return ok
}

// ResetConfirmationToken reset all changes of the "confirmation_token" field.
func (m *UserMutation) ResetConfirmationToken() {
	m.confirmation_token = nil
	delete(m.clearedFields, user.FieldConfirmationToken)
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (m *UserMutation) SetRecoverySentAt(t time.Time) {
	m.recovery_sent_at = &t
}

// RecoverySentAt returns the recovery_sent_at value in the mutation.
func (m *UserMutation) RecoverySentAt() (r time.Time, exists bool) {
	v := m.recovery_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoverySentAt returns the old recovery_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRecoverySentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecoverySentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecoverySentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoverySentAt: %w", err)
	}
	return oldValue.RecoverySentAt, nil
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (m *UserMutation) ClearRecoverySentAt() {
	m.recovery_sent_at = nil
	m.clearedFields[user.FieldRecoverySentAt] = struct{}{}
}

// RecoverySentAtCleared returns if the field recovery_sent_at was cleared in this mutation.
func (m *UserMutation) RecoverySentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldRecoverySentAt]
	return ok
}

// ResetRecoverySentAt reset all changes of the "recovery_sent_at" field.
func (m *UserMutation) ResetRecoverySentAt() {
	m.recovery_sent_at = nil
	delete(m.clearedFields, user.FieldRecoverySentAt)
}

// SetRecoveryToken sets the recovery_token field.
func (m *UserMutation) SetRecoveryToken(s string) {
	m.recovery_token = &s
}

// RecoveryToken returns the recovery_token value in the mutation.
func (m *UserMutation) RecoveryToken() (r string, exists bool) {
	v := m.recovery_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoveryToken returns the old recovery_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRecoveryToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecoveryToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecoveryToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoveryToken: %w", err)
	}
	return oldValue.RecoveryToken, nil
}

// ClearRecoveryToken clears the value of recovery_token.
func (m *UserMutation) ClearRecoveryToken() {
	m.recovery_token = nil
	m.clearedFields[user.FieldRecoveryToken] = struct{}{}
}

// RecoveryTokenCleared returns if the field recovery_token was cleared in this mutation.
func (m *UserMutation) RecoveryTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldRecoveryToken]
	return ok
}

// ResetRecoveryToken reset all changes of the "recovery_token" field.
func (m *UserMutation) ResetRecoveryToken() {
	m.recovery_token = nil
	delete(m.clearedFields, user.FieldRecoveryToken)
}

// SetOtpSentAt sets the otp_sent_at field.
func (m *UserMutation) SetOtpSentAt(t time.Time) {
	m.otp_sent_at = &t
}

// OtpSentAt returns the otp_sent_at value in the mutation.
func (m *UserMutation) OtpSentAt() (r time.Time, exists bool) {
	v := m.otp_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldOtpSentAt returns the old otp_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldOtpSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOtpSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOtpSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtpSentAt: %w", err)
	}
	return oldValue.OtpSentAt, nil
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (m *UserMutation) ClearOtpSentAt() {
	m.otp_sent_at = nil
	m.clearedFields[user.FieldOtpSentAt] = struct{}{}
}

// OtpSentAtCleared returns if the field otp_sent_at was cleared in this mutation.
func (m *UserMutation) OtpSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldOtpSentAt]
	return ok
}

// ResetOtpSentAt reset all changes of the "otp_sent_at" field.
func (m *UserMutation) ResetOtpSentAt() {
	m.otp_sent_at = nil
	delete(m.clearedFields, user.FieldOtpSentAt)
}

// SetOtp sets the otp field.
func (m *UserMutation) SetOtp(s string) {
	m.otp = &s
}

// Otp returns the otp value in the mutation.
func (m *UserMutation) Otp() (r string, exists bool) {
	v := m.otp
	if v == nil {
		return
	}
	return *v, true
}

// OldOtp returns the old otp value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldOtp(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOtp is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOtp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtp: %w", err)
	}
	return oldValue.Otp, nil
}

// ClearOtp clears the value of otp.
func (m *UserMutation) ClearOtp() {
	m.otp = nil
	m.clearedFields[user.FieldOtp] = struct{}{}
}

// OtpCleared returns if the field otp was cleared in this mutation.
func (m *UserMutation) OtpCleared() bool {
	_, ok := m.clearedFields[user.FieldOtp]
	return ok
}

// ResetOtp reset all changes of the "otp" field.
func (m *UserMutation) ResetOtp() {
	m.otp = nil
	delete(m.clearedFields, user.FieldOtp)
}

// SetEmailChange sets the email_change field.
func (m *UserMutation) SetEmailChange(s string) {
	m.email_change = &s
}

// EmailChange returns the email_change value in the mutation.
func (m *UserMutation) EmailChange() (r string, exists bool) {
	v := m.email_change
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChange returns the old email_change value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChange(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChange is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChange requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChange: %w", err)
	}
	return oldValue.EmailChange, nil
}

// ClearEmailChange clears the value of email_change.
func (m *UserMutation) ClearEmailChange() {
	m.email_change = nil
	m.clearedFields[user.FieldEmailChange] = struct{}{}
}

// EmailChangeCleared returns if the field email_change was cleared in this mutation.
func (m *UserMutation) EmailChangeCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChange]
	return ok
}

// ResetEmailChange reset all changes of the "email_change" field.
func (m *UserMutation) ResetEmailChange() {
	m.email_change = nil
	delete(m.clearedFields, user.FieldEmailChange)
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (m *UserMutation) SetEmailChangeSentAt(t time.Time) {
	m.email_change_sent_at = &t
}

// EmailChangeSentAt returns the email_change_sent_at value in the mutation.
func (m *UserMutation) EmailChangeSentAt() (r time.Time, exists bool) {
	v := m.email_change_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeSentAt returns the old email_change_sent_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChangeSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChangeSentAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChangeSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeSentAt: %w", err)
	}
	return oldValue.EmailChangeSentAt, nil
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (m *UserMutation) ClearEmailChangeSentAt() {
	m.email_change_sent_at = nil
	m.clearedFields[user.FieldEmailChangeSentAt] = struct{}{}
}

// EmailChangeSentAtCleared returns if the field email_change_sent_at was cleared in this mutation.
func (m *UserMutation) EmailChangeSentAtCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChangeSentAt]
	return ok
}

// ResetEmailChangeSentAt reset all changes of the "email_change_sent_at" field.
func (m *UserMutation) ResetEmailChangeSentAt() {
	m.email_change_sent_at = nil
	delete(m.clearedFields, user.FieldEmailChangeSentAt)
}

// SetEmailChangeToken sets the email_change_token field.
func (m *UserMutation) SetEmailChangeToken(s string) {
	m.email_change_token = &s
}

// EmailChangeToken returns the email_change_token value in the mutation.
func (m *UserMutation) EmailChangeToken() (r string, exists bool) {
	v := m.email_change_token
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeToken returns the old email_change_token value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldEmailChangeToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmailChangeToken is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmailChangeToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeToken: %w", err)
	}
	return oldValue.EmailChangeToken, nil
}

// ClearEmailChangeToken clears the value of email_change_token.
func (m *UserMutation) ClearEmailChangeToken() {
	m.email_change_token = nil
	m.clearedFields[user.FieldEmailChangeToken] = struct{}{}
}

// EmailChangeTokenCleared returns if the field email_change_token was cleared in this mutation.
func (m *UserMutation) EmailChangeTokenCleared() bool {
	_, ok := m.clearedFields[user.FieldEmailChangeToken]
	return ok
}

// ResetEmailChangeToken reset all changes of the "email_change_token" field.
func (m *UserMutation) ResetEmailChangeToken() {
	m.email_change_token = nil
	delete(m.clearedFields, user.FieldEmailChangeToken)
}

// SetMetadata sets the metadata field.
func (m *UserMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *UserMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old metadata value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldMetadata(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMetadata is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ResetMetadata reset all changes of the "metadata" field.
func (m *UserMutation) ResetMetadata() {
	m.metadata = nil
}

// SetCreatedAt sets the created_at field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetLastSigninAt sets the last_signin_at field.
func (m *UserMutation) SetLastSigninAt(t time.Time) {
	m.last_signin_at = &t
}

// LastSigninAt returns the last_signin_at value in the mutation.
func (m *UserMutation) LastSigninAt() (r time.Time, exists bool) {
	v := m.last_signin_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastSigninAt returns the old last_signin_at value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldLastSigninAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastSigninAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastSigninAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastSigninAt: %w", err)
	}
	return oldValue.LastSigninAt, nil
}

// ClearLastSigninAt clears the value of last_signin_at.
func (m *UserMutation) ClearLastSigninAt() {
	m.last_signin_at = nil
	m.clearedFields[user.FieldLastSigninAt] = struct{}{}
}

// LastSigninAtCleared returns if the field last_signin_at was cleared in this mutation.
func (m *UserMutation) LastSigninAtCleared() bool {
	_, ok := m.clearedFields[user.FieldLastSigninAt]
	return ok
}

// ResetLastSigninAt reset all changes of the "last_signin_at" field.
func (m *UserMutation) ResetLastSigninAt() {
	m.last_signin_at = nil
	delete(m.clearedFields, user.FieldLastSigninAt)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 16)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.confirmed != nil {
		fields = append(fields, user.FieldConfirmed)
	}
	if m.confirmation_sent_at != nil {
		fields = append(fields, user.FieldConfirmationSentAt)
	}
	if m.confirmation_token != nil {
		fields = append(fields, user.FieldConfirmationToken)
	}
	if m.recovery_sent_at != nil {
		fields = append(fields, user.FieldRecoverySentAt)
	}
	if m.recovery_token != nil {
		fields = append(fields, user.FieldRecoveryToken)
	}
	if m.otp_sent_at != nil {
		fields = append(fields, user.FieldOtpSentAt)
	}
	if m.otp != nil {
		fields = append(fields, user.FieldOtp)
	}
	if m.email_change != nil {
		fields = append(fields, user.FieldEmailChange)
	}
	if m.email_change_sent_at != nil {
		fields = append(fields, user.FieldEmailChangeSentAt)
	}
	if m.email_change_token != nil {
		fields = append(fields, user.FieldEmailChangeToken)
	}
	if m.metadata != nil {
		fields = append(fields, user.FieldMetadata)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.last_signin_at != nil {
		fields = append(fields, user.FieldLastSigninAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldConfirmed:
		return m.Confirmed()
	case user.FieldConfirmationSentAt:
		return m.ConfirmationSentAt()
	case user.FieldConfirmationToken:
		return m.ConfirmationToken()
	case user.FieldRecoverySentAt:
		return m.RecoverySentAt()
	case user.FieldRecoveryToken:
		return m.RecoveryToken()
	case user.FieldOtpSentAt:
		return m.OtpSentAt()
	case user.FieldOtp:
		return m.Otp()
	case user.FieldEmailChange:
		return m.EmailChange()
	case user.FieldEmailChangeSentAt:
		return m.EmailChangeSentAt()
	case user.FieldEmailChangeToken:
		return m.EmailChangeToken()
	case user.FieldMetadata:
		return m.Metadata()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldLastSigninAt:
		return m.LastSigninAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldConfirmed:
		return m.OldConfirmed(ctx)
	case user.FieldConfirmationSentAt:
		return m.OldConfirmationSentAt(ctx)
	case user.FieldConfirmationToken:
		return m.OldConfirmationToken(ctx)
	case user.FieldRecoverySentAt:
		return m.OldRecoverySentAt(ctx)
	case user.FieldRecoveryToken:
		return m.OldRecoveryToken(ctx)
	case user.FieldOtpSentAt:
		return m.OldOtpSentAt(ctx)
	case user.FieldOtp:
		return m.OldOtp(ctx)
	case user.FieldEmailChange:
		return m.OldEmailChange(ctx)
	case user.FieldEmailChangeSentAt:
		return m.OldEmailChangeSentAt(ctx)
	case user.FieldEmailChangeToken:
		return m.OldEmailChangeToken(ctx)
	case user.FieldMetadata:
		return m.OldMetadata(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldLastSigninAt:
		return m.OldLastSigninAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldConfirmed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmed(v)
		return nil
	case user.FieldConfirmationSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationSentAt(v)
		return nil
	case user.FieldConfirmationToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationToken(v)
		return nil
	case user.FieldRecoverySentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoverySentAt(v)
		return nil
	case user.FieldRecoveryToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoveryToken(v)
		return nil
	case user.FieldOtpSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtpSentAt(v)
		return nil
	case user.FieldOtp:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtp(v)
		return nil
	case user.FieldEmailChange:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChange(v)
		return nil
	case user.FieldEmailChangeSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeSentAt(v)
		return nil
	case user.FieldEmailChangeToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeToken(v)
		return nil
	case user.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldLastSigninAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastSigninAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldConfirmed) {
		fields = append(fields, user.FieldConfirmed)
	}
	if m.FieldCleared(user.FieldConfirmationSentAt) {
		fields = append(fields, user.FieldConfirmationSentAt)
	}
	if m.FieldCleared(user.FieldConfirmationToken) {
		fields = append(fields, user.FieldConfirmationToken)
	}
	if m.FieldCleared(user.FieldRecoverySentAt) {
		fields = append(fields, user.FieldRecoverySentAt)
	}
	if m.FieldCleared(user.FieldRecoveryToken) {
		fields = append(fields, user.FieldRecoveryToken)
	}
	if m.FieldCleared(user.FieldOtpSentAt) {
		fields = append(fields, user.FieldOtpSentAt)
	}
	if m.FieldCleared(user.FieldOtp) {
		fields = append(fields, user.FieldOtp)
	}
	if m.FieldCleared(user.FieldEmailChange) {
		fields = append(fields, user.FieldEmailChange)
	}
	if m.FieldCleared(user.FieldEmailChangeSentAt) {
		fields = append(fields, user.FieldEmailChangeSentAt)
	}
	if m.FieldCleared(user.FieldEmailChangeToken) {
		fields = append(fields, user.FieldEmailChangeToken)
	}
	if m.FieldCleared(user.FieldLastSigninAt) {
		fields = append(fields, user.FieldLastSigninAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldConfirmed:
		m.ClearConfirmed()
		return nil
	case user.FieldConfirmationSentAt:
		m.ClearConfirmationSentAt()
		return nil
	case user.FieldConfirmationToken:
		m.ClearConfirmationToken()
		return nil
	case user.FieldRecoverySentAt:
		m.ClearRecoverySentAt()
		return nil
	case user.FieldRecoveryToken:
		m.ClearRecoveryToken()
		return nil
	case user.FieldOtpSentAt:
		m.ClearOtpSentAt()
		return nil
	case user.FieldOtp:
		m.ClearOtp()
		return nil
	case user.FieldEmailChange:
		m.ClearEmailChange()
		return nil
	case user.FieldEmailChangeSentAt:
		m.ClearEmailChangeSentAt()
		return nil
	case user.FieldEmailChangeToken:
		m.ClearEmailChangeToken()
		return nil
	case user.FieldLastSigninAt:
		m.ClearLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldConfirmed:
		m.ResetConfirmed()
		return nil
	case user.FieldConfirmationSentAt:
		m.ResetConfirmationSentAt()
		return nil
	case user.FieldConfirmationToken:
		m.ResetConfirmationToken()
		return nil
	case user.FieldRecoverySentAt:
		m.ResetRecoverySentAt()
		return nil
	case user.FieldRecoveryToken:
		m.ResetRecoveryToken()
		return nil
	case user.FieldOtpSentAt:
		m.ResetOtpSentAt()
		return nil
	case user.FieldOtp:
		m.ResetOtp()
		return nil
	case user.FieldEmailChange:
		m.ResetEmailChange()
		return nil
	case user.FieldEmailChangeSentAt:
		m.ResetEmailChangeSentAt()
		return nil
	case user.FieldEmailChangeToken:
		m.ResetEmailChangeToken()
		return nil
	case user.FieldMetadata:
		m.ResetMetadata()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldLastSigninAt:
		m.ResetLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
