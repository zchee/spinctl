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

// AccountDetails struct for AccountDetails
type AccountDetails struct {
	AccountId                   *string              `json:"accountId,omitempty"`
	AccountType                 *string              `json:"accountType,omitempty"`
	ChallengeDestructiveActions *bool                `json:"challengeDestructiveActions,omitempty"`
	CloudProvider               *string              `json:"cloudProvider,omitempty"`
	Environment                 *string              `json:"environment,omitempty"`
	Name                        *string              `json:"name,omitempty"`
	Permissions                 *map[string][]string `json:"permissions,omitempty"`
	PrimaryAccount              *bool                `json:"primaryAccount,omitempty"`
	ProviderVersion             *string              `json:"providerVersion,omitempty"`
	RequiredGroupMembership     *[]string            `json:"requiredGroupMembership,omitempty"`
	Skin                        *string              `json:"skin,omitempty"`
	Type                        *string              `json:"type,omitempty"`
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountDetails) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountDetails) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountDetails) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountDetails) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetAccountTypeOk() (string, bool) {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret, false
	}
	return *o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountDetails) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *AccountDetails) SetAccountType(v string) {
	o.AccountType = &v
}

// GetChallengeDestructiveActions returns the ChallengeDestructiveActions field value if set, zero value otherwise.
func (o *AccountDetails) GetChallengeDestructiveActions() bool {
	if o == nil || o.ChallengeDestructiveActions == nil {
		var ret bool
		return ret
	}
	return *o.ChallengeDestructiveActions
}

// GetChallengeDestructiveActionsOk returns a tuple with the ChallengeDestructiveActions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetChallengeDestructiveActionsOk() (bool, bool) {
	if o == nil || o.ChallengeDestructiveActions == nil {
		var ret bool
		return ret, false
	}
	return *o.ChallengeDestructiveActions, true
}

// HasChallengeDestructiveActions returns a boolean if a field has been set.
func (o *AccountDetails) HasChallengeDestructiveActions() bool {
	if o != nil && o.ChallengeDestructiveActions != nil {
		return true
	}

	return false
}

// SetChallengeDestructiveActions gets a reference to the given bool and assigns it to the ChallengeDestructiveActions field.
func (o *AccountDetails) SetChallengeDestructiveActions(v bool) {
	o.ChallengeDestructiveActions = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *AccountDetails) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetCloudProviderOk() (string, bool) {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret, false
	}
	return *o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *AccountDetails) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *AccountDetails) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AccountDetails) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetEnvironmentOk() (string, bool) {
	if o == nil || o.Environment == nil {
		var ret string
		return ret, false
	}
	return *o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AccountDetails) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AccountDetails) SetEnvironment(v string) {
	o.Environment = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountDetails) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AccountDetails) GetPermissions() map[string][]string {
	if o == nil || o.Permissions == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetPermissionsOk() (map[string][]string, bool) {
	if o == nil || o.Permissions == nil {
		var ret map[string][]string
		return ret, false
	}
	return *o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AccountDetails) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string][]string and assigns it to the Permissions field.
func (o *AccountDetails) SetPermissions(v map[string][]string) {
	o.Permissions = &v
}

// GetPrimaryAccount returns the PrimaryAccount field value if set, zero value otherwise.
func (o *AccountDetails) GetPrimaryAccount() bool {
	if o == nil || o.PrimaryAccount == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryAccount
}

// GetPrimaryAccountOk returns a tuple with the PrimaryAccount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetPrimaryAccountOk() (bool, bool) {
	if o == nil || o.PrimaryAccount == nil {
		var ret bool
		return ret, false
	}
	return *o.PrimaryAccount, true
}

// HasPrimaryAccount returns a boolean if a field has been set.
func (o *AccountDetails) HasPrimaryAccount() bool {
	if o != nil && o.PrimaryAccount != nil {
		return true
	}

	return false
}

// SetPrimaryAccount gets a reference to the given bool and assigns it to the PrimaryAccount field.
func (o *AccountDetails) SetPrimaryAccount(v bool) {
	o.PrimaryAccount = &v
}

// GetProviderVersion returns the ProviderVersion field value if set, zero value otherwise.
func (o *AccountDetails) GetProviderVersion() string {
	if o == nil || o.ProviderVersion == nil {
		var ret string
		return ret
	}
	return *o.ProviderVersion
}

// GetProviderVersionOk returns a tuple with the ProviderVersion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetProviderVersionOk() (string, bool) {
	if o == nil || o.ProviderVersion == nil {
		var ret string
		return ret, false
	}
	return *o.ProviderVersion, true
}

// HasProviderVersion returns a boolean if a field has been set.
func (o *AccountDetails) HasProviderVersion() bool {
	if o != nil && o.ProviderVersion != nil {
		return true
	}

	return false
}

// SetProviderVersion gets a reference to the given string and assigns it to the ProviderVersion field.
func (o *AccountDetails) SetProviderVersion(v string) {
	o.ProviderVersion = &v
}

// GetRequiredGroupMembership returns the RequiredGroupMembership field value if set, zero value otherwise.
func (o *AccountDetails) GetRequiredGroupMembership() []string {
	if o == nil || o.RequiredGroupMembership == nil {
		var ret []string
		return ret
	}
	return *o.RequiredGroupMembership
}

// GetRequiredGroupMembershipOk returns a tuple with the RequiredGroupMembership field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetRequiredGroupMembershipOk() ([]string, bool) {
	if o == nil || o.RequiredGroupMembership == nil {
		var ret []string
		return ret, false
	}
	return *o.RequiredGroupMembership, true
}

// HasRequiredGroupMembership returns a boolean if a field has been set.
func (o *AccountDetails) HasRequiredGroupMembership() bool {
	if o != nil && o.RequiredGroupMembership != nil {
		return true
	}

	return false
}

// SetRequiredGroupMembership gets a reference to the given []string and assigns it to the RequiredGroupMembership field.
func (o *AccountDetails) SetRequiredGroupMembership(v []string) {
	o.RequiredGroupMembership = &v
}

// GetSkin returns the Skin field value if set, zero value otherwise.
func (o *AccountDetails) GetSkin() string {
	if o == nil || o.Skin == nil {
		var ret string
		return ret
	}
	return *o.Skin
}

// GetSkinOk returns a tuple with the Skin field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetSkinOk() (string, bool) {
	if o == nil || o.Skin == nil {
		var ret string
		return ret, false
	}
	return *o.Skin, true
}

// HasSkin returns a boolean if a field has been set.
func (o *AccountDetails) HasSkin() bool {
	if o != nil && o.Skin != nil {
		return true
	}

	return false
}

// SetSkin gets a reference to the given string and assigns it to the Skin field.
func (o *AccountDetails) SetSkin(v string) {
	o.Skin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountDetails) SetType(v string) {
	o.Type = &v
}

type NullableAccountDetails struct {
	Value        AccountDetails
	ExplicitNull bool
}

func (v NullableAccountDetails) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccountDetails) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
