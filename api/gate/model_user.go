/*
 * Spinnaker API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.18.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gate

import (
	"bytes"
	"encoding/json"
)

// User struct for User
type User struct {
	AccountNonExpired     *bool               `json:"accountNonExpired,omitempty"`
	AccountNonLocked      *bool               `json:"accountNonLocked,omitempty"`
	AllowedAccounts       *[]string           `json:"allowedAccounts,omitempty"`
	Authorities           *[]GrantedAuthority `json:"authorities,omitempty"`
	CredentialsNonExpired *bool               `json:"credentialsNonExpired,omitempty"`
	Email                 *string             `json:"email,omitempty"`
	Enabled               *bool               `json:"enabled,omitempty"`
	FirstName             *string             `json:"firstName,omitempty"`
	LastName              *string             `json:"lastName,omitempty"`
	Roles                 *[]string           `json:"roles,omitempty"`
	Username              *string             `json:"username,omitempty"`
}

// GetAccountNonExpired returns the AccountNonExpired field value if set, zero value otherwise.
func (o *User) GetAccountNonExpired() bool {
	if o == nil || o.AccountNonExpired == nil {
		var ret bool
		return ret
	}
	return *o.AccountNonExpired
}

// GetAccountNonExpiredOk returns a tuple with the AccountNonExpired field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountNonExpiredOk() (bool, bool) {
	if o == nil || o.AccountNonExpired == nil {
		var ret bool
		return ret, false
	}
	return *o.AccountNonExpired, true
}

// HasAccountNonExpired returns a boolean if a field has been set.
func (o *User) HasAccountNonExpired() bool {
	if o != nil && o.AccountNonExpired != nil {
		return true
	}

	return false
}

// SetAccountNonExpired gets a reference to the given bool and assigns it to the AccountNonExpired field.
func (o *User) SetAccountNonExpired(v bool) {
	o.AccountNonExpired = &v
}

// GetAccountNonLocked returns the AccountNonLocked field value if set, zero value otherwise.
func (o *User) GetAccountNonLocked() bool {
	if o == nil || o.AccountNonLocked == nil {
		var ret bool
		return ret
	}
	return *o.AccountNonLocked
}

// GetAccountNonLockedOk returns a tuple with the AccountNonLocked field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountNonLockedOk() (bool, bool) {
	if o == nil || o.AccountNonLocked == nil {
		var ret bool
		return ret, false
	}
	return *o.AccountNonLocked, true
}

// HasAccountNonLocked returns a boolean if a field has been set.
func (o *User) HasAccountNonLocked() bool {
	if o != nil && o.AccountNonLocked != nil {
		return true
	}

	return false
}

// SetAccountNonLocked gets a reference to the given bool and assigns it to the AccountNonLocked field.
func (o *User) SetAccountNonLocked(v bool) {
	o.AccountNonLocked = &v
}

// GetAllowedAccounts returns the AllowedAccounts field value if set, zero value otherwise.
func (o *User) GetAllowedAccounts() []string {
	if o == nil || o.AllowedAccounts == nil {
		var ret []string
		return ret
	}
	return *o.AllowedAccounts
}

// GetAllowedAccountsOk returns a tuple with the AllowedAccounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAllowedAccountsOk() ([]string, bool) {
	if o == nil || o.AllowedAccounts == nil {
		var ret []string
		return ret, false
	}
	return *o.AllowedAccounts, true
}

// HasAllowedAccounts returns a boolean if a field has been set.
func (o *User) HasAllowedAccounts() bool {
	if o != nil && o.AllowedAccounts != nil {
		return true
	}

	return false
}

// SetAllowedAccounts gets a reference to the given []string and assigns it to the AllowedAccounts field.
func (o *User) SetAllowedAccounts(v []string) {
	o.AllowedAccounts = &v
}

// GetAuthorities returns the Authorities field value if set, zero value otherwise.
func (o *User) GetAuthorities() []GrantedAuthority {
	if o == nil || o.Authorities == nil {
		var ret []GrantedAuthority
		return ret
	}
	return *o.Authorities
}

// GetAuthoritiesOk returns a tuple with the Authorities field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAuthoritiesOk() ([]GrantedAuthority, bool) {
	if o == nil || o.Authorities == nil {
		var ret []GrantedAuthority
		return ret, false
	}
	return *o.Authorities, true
}

// HasAuthorities returns a boolean if a field has been set.
func (o *User) HasAuthorities() bool {
	if o != nil && o.Authorities != nil {
		return true
	}

	return false
}

// SetAuthorities gets a reference to the given []GrantedAuthority and assigns it to the Authorities field.
func (o *User) SetAuthorities(v []GrantedAuthority) {
	o.Authorities = &v
}

// GetCredentialsNonExpired returns the CredentialsNonExpired field value if set, zero value otherwise.
func (o *User) GetCredentialsNonExpired() bool {
	if o == nil || o.CredentialsNonExpired == nil {
		var ret bool
		return ret
	}
	return *o.CredentialsNonExpired
}

// GetCredentialsNonExpiredOk returns a tuple with the CredentialsNonExpired field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCredentialsNonExpiredOk() (bool, bool) {
	if o == nil || o.CredentialsNonExpired == nil {
		var ret bool
		return ret, false
	}
	return *o.CredentialsNonExpired, true
}

// HasCredentialsNonExpired returns a boolean if a field has been set.
func (o *User) HasCredentialsNonExpired() bool {
	if o != nil && o.CredentialsNonExpired != nil {
		return true
	}

	return false
}

// SetCredentialsNonExpired gets a reference to the given bool and assigns it to the CredentialsNonExpired field.
func (o *User) SetCredentialsNonExpired(v bool) {
	o.CredentialsNonExpired = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *User) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnabledOk() (bool, bool) {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret, false
	}
	return *o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *User) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *User) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (string, bool) {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret, false
	}
	return *o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (string, bool) {
	if o == nil || o.LastName == nil {
		var ret string
		return ret, false
	}
	return *o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *User) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret, false
	}
	return *o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *User) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *User) SetRoles(v []string) {
	o.Roles = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (string, bool) {
	if o == nil || o.Username == nil {
		var ret string
		return ret, false
	}
	return *o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

type NullableUser struct {
	Value        User
	ExplicitNull bool
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
