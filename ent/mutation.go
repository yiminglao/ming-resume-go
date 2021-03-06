// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go_resume/ent/user"
	"sync"

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
	TypeUser = "User"
)

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	firstname     *string
	lastname      *string
	username      *string
	password      *string
	address       *string
	city          *string
	state         *string
	zipcode       *string
	email         *string
	phone         *string
	photo         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
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
func withUserID(id int) userOption {
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
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetFirstname sets the firstname field.
func (m *UserMutation) SetFirstname(s string) {
	m.firstname = &s
}

// Firstname returns the firstname value in the mutation.
func (m *UserMutation) Firstname() (r string, exists bool) {
	v := m.firstname
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstname returns the old firstname value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldFirstname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFirstname is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFirstname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstname: %w", err)
	}
	return oldValue.Firstname, nil
}

// ResetFirstname reset all changes of the "firstname" field.
func (m *UserMutation) ResetFirstname() {
	m.firstname = nil
}

// SetLastname sets the lastname field.
func (m *UserMutation) SetLastname(s string) {
	m.lastname = &s
}

// Lastname returns the lastname value in the mutation.
func (m *UserMutation) Lastname() (r string, exists bool) {
	v := m.lastname
	if v == nil {
		return
	}
	return *v, true
}

// OldLastname returns the old lastname value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldLastname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastname is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastname: %w", err)
	}
	return oldValue.Lastname, nil
}

// ResetLastname reset all changes of the "lastname" field.
func (m *UserMutation) ResetLastname() {
	m.lastname = nil
}

// SetUsername sets the username field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the username value in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old username value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUsername is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername reset all changes of the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
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

// SetAddress sets the address field.
func (m *UserMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the address value in the mutation.
func (m *UserMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old address value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAddress is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ClearAddress clears the value of address.
func (m *UserMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[user.FieldAddress] = struct{}{}
}

// AddressCleared returns if the field address was cleared in this mutation.
func (m *UserMutation) AddressCleared() bool {
	_, ok := m.clearedFields[user.FieldAddress]
	return ok
}

// ResetAddress reset all changes of the "address" field.
func (m *UserMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, user.FieldAddress)
}

// SetCity sets the city field.
func (m *UserMutation) SetCity(s string) {
	m.city = &s
}

// City returns the city value in the mutation.
func (m *UserMutation) City() (r string, exists bool) {
	v := m.city
	if v == nil {
		return
	}
	return *v, true
}

// OldCity returns the old city value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldCity(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCity is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCity: %w", err)
	}
	return oldValue.City, nil
}

// ClearCity clears the value of city.
func (m *UserMutation) ClearCity() {
	m.city = nil
	m.clearedFields[user.FieldCity] = struct{}{}
}

// CityCleared returns if the field city was cleared in this mutation.
func (m *UserMutation) CityCleared() bool {
	_, ok := m.clearedFields[user.FieldCity]
	return ok
}

// ResetCity reset all changes of the "city" field.
func (m *UserMutation) ResetCity() {
	m.city = nil
	delete(m.clearedFields, user.FieldCity)
}

// SetState sets the state field.
func (m *UserMutation) SetState(s string) {
	m.state = &s
}

// State returns the state value in the mutation.
func (m *UserMutation) State() (r string, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old state value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldState(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldState is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ClearState clears the value of state.
func (m *UserMutation) ClearState() {
	m.state = nil
	m.clearedFields[user.FieldState] = struct{}{}
}

// StateCleared returns if the field state was cleared in this mutation.
func (m *UserMutation) StateCleared() bool {
	_, ok := m.clearedFields[user.FieldState]
	return ok
}

// ResetState reset all changes of the "state" field.
func (m *UserMutation) ResetState() {
	m.state = nil
	delete(m.clearedFields, user.FieldState)
}

// SetZipcode sets the zipcode field.
func (m *UserMutation) SetZipcode(s string) {
	m.zipcode = &s
}

// Zipcode returns the zipcode value in the mutation.
func (m *UserMutation) Zipcode() (r string, exists bool) {
	v := m.zipcode
	if v == nil {
		return
	}
	return *v, true
}

// OldZipcode returns the old zipcode value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldZipcode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldZipcode is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldZipcode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldZipcode: %w", err)
	}
	return oldValue.Zipcode, nil
}

// ClearZipcode clears the value of zipcode.
func (m *UserMutation) ClearZipcode() {
	m.zipcode = nil
	m.clearedFields[user.FieldZipcode] = struct{}{}
}

// ZipcodeCleared returns if the field zipcode was cleared in this mutation.
func (m *UserMutation) ZipcodeCleared() bool {
	_, ok := m.clearedFields[user.FieldZipcode]
	return ok
}

// ResetZipcode reset all changes of the "zipcode" field.
func (m *UserMutation) ResetZipcode() {
	m.zipcode = nil
	delete(m.clearedFields, user.FieldZipcode)
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

// ClearEmail clears the value of email.
func (m *UserMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[user.FieldEmail] = struct{}{}
}

// EmailCleared returns if the field email was cleared in this mutation.
func (m *UserMutation) EmailCleared() bool {
	_, ok := m.clearedFields[user.FieldEmail]
	return ok
}

// ResetEmail reset all changes of the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, user.FieldEmail)
}

// SetPhone sets the phone field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the phone value in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old phone value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPhone is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of phone.
func (m *UserMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[user.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the field phone was cleared in this mutation.
func (m *UserMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[user.FieldPhone]
	return ok
}

// ResetPhone reset all changes of the "phone" field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, user.FieldPhone)
}

// SetPhoto sets the photo field.
func (m *UserMutation) SetPhoto(s string) {
	m.photo = &s
}

// Photo returns the photo value in the mutation.
func (m *UserMutation) Photo() (r string, exists bool) {
	v := m.photo
	if v == nil {
		return
	}
	return *v, true
}

// OldPhoto returns the old photo value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldPhoto(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPhoto is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPhoto requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhoto: %w", err)
	}
	return oldValue.Photo, nil
}

// ClearPhoto clears the value of photo.
func (m *UserMutation) ClearPhoto() {
	m.photo = nil
	m.clearedFields[user.FieldPhoto] = struct{}{}
}

// PhotoCleared returns if the field photo was cleared in this mutation.
func (m *UserMutation) PhotoCleared() bool {
	_, ok := m.clearedFields[user.FieldPhoto]
	return ok
}

// ResetPhoto reset all changes of the "photo" field.
func (m *UserMutation) ResetPhoto() {
	m.photo = nil
	delete(m.clearedFields, user.FieldPhoto)
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
	fields := make([]string, 0, 11)
	if m.firstname != nil {
		fields = append(fields, user.FieldFirstname)
	}
	if m.lastname != nil {
		fields = append(fields, user.FieldLastname)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.address != nil {
		fields = append(fields, user.FieldAddress)
	}
	if m.city != nil {
		fields = append(fields, user.FieldCity)
	}
	if m.state != nil {
		fields = append(fields, user.FieldState)
	}
	if m.zipcode != nil {
		fields = append(fields, user.FieldZipcode)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.photo != nil {
		fields = append(fields, user.FieldPhoto)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldFirstname:
		return m.Firstname()
	case user.FieldLastname:
		return m.Lastname()
	case user.FieldUsername:
		return m.Username()
	case user.FieldPassword:
		return m.Password()
	case user.FieldAddress:
		return m.Address()
	case user.FieldCity:
		return m.City()
	case user.FieldState:
		return m.State()
	case user.FieldZipcode:
		return m.Zipcode()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPhone:
		return m.Phone()
	case user.FieldPhoto:
		return m.Photo()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldFirstname:
		return m.OldFirstname(ctx)
	case user.FieldLastname:
		return m.OldLastname(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldAddress:
		return m.OldAddress(ctx)
	case user.FieldCity:
		return m.OldCity(ctx)
	case user.FieldState:
		return m.OldState(ctx)
	case user.FieldZipcode:
		return m.OldZipcode(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPhone:
		return m.OldPhone(ctx)
	case user.FieldPhoto:
		return m.OldPhoto(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldFirstname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstname(v)
		return nil
	case user.FieldLastname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastname(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case user.FieldCity:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCity(v)
		return nil
	case user.FieldState:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case user.FieldZipcode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetZipcode(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldPhoto:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhoto(v)
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
	if m.FieldCleared(user.FieldAddress) {
		fields = append(fields, user.FieldAddress)
	}
	if m.FieldCleared(user.FieldCity) {
		fields = append(fields, user.FieldCity)
	}
	if m.FieldCleared(user.FieldState) {
		fields = append(fields, user.FieldState)
	}
	if m.FieldCleared(user.FieldZipcode) {
		fields = append(fields, user.FieldZipcode)
	}
	if m.FieldCleared(user.FieldEmail) {
		fields = append(fields, user.FieldEmail)
	}
	if m.FieldCleared(user.FieldPhone) {
		fields = append(fields, user.FieldPhone)
	}
	if m.FieldCleared(user.FieldPhoto) {
		fields = append(fields, user.FieldPhoto)
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
	case user.FieldAddress:
		m.ClearAddress()
		return nil
	case user.FieldCity:
		m.ClearCity()
		return nil
	case user.FieldState:
		m.ClearState()
		return nil
	case user.FieldZipcode:
		m.ClearZipcode()
		return nil
	case user.FieldEmail:
		m.ClearEmail()
		return nil
	case user.FieldPhone:
		m.ClearPhone()
		return nil
	case user.FieldPhoto:
		m.ClearPhoto()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldFirstname:
		m.ResetFirstname()
		return nil
	case user.FieldLastname:
		m.ResetLastname()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldAddress:
		m.ResetAddress()
		return nil
	case user.FieldCity:
		m.ResetCity()
		return nil
	case user.FieldState:
		m.ResetState()
		return nil
	case user.FieldZipcode:
		m.ResetZipcode()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldPhoto:
		m.ResetPhoto()
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
