// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go_resume/ent/user"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Zipcode holds the value of the "zipcode" field.
	Zipcode string `json:"zipcode,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Photo holds the value of the "photo" field.
	Photo string `json:"photo,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // firstname
		&sql.NullString{}, // lastname
		&sql.NullString{}, // username
		&sql.NullString{}, // password
		&sql.NullString{}, // address
		&sql.NullString{}, // city
		&sql.NullString{}, // state
		&sql.NullString{}, // zipcode
		&sql.NullString{}, // email
		&sql.NullString{}, // phone
		&sql.NullString{}, // photo
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field firstname", values[0])
	} else if value.Valid {
		u.Firstname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field lastname", values[1])
	} else if value.Valid {
		u.Lastname = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[2])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[3])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[4])
	} else if value.Valid {
		u.Address = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field city", values[5])
	} else if value.Valid {
		u.City = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[6])
	} else if value.Valid {
		u.State = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field zipcode", values[7])
	} else if value.Valid {
		u.Zipcode = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[8])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[9])
	} else if value.Valid {
		u.Phone = value.String
	}
	if value, ok := values[10].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field photo", values[10])
	} else if value.Valid {
		u.Photo = value.String
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", firstname=")
	builder.WriteString(u.Firstname)
	builder.WriteString(", lastname=")
	builder.WriteString(u.Lastname)
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", address=")
	builder.WriteString(u.Address)
	builder.WriteString(", city=")
	builder.WriteString(u.City)
	builder.WriteString(", state=")
	builder.WriteString(u.State)
	builder.WriteString(", zipcode=")
	builder.WriteString(u.Zipcode)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", photo=")
	builder.WriteString(u.Photo)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
