# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**ChallengeDestructiveActions** | Pointer to **bool** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**PrimaryAccount** | Pointer to **bool** |  | [optional] 
**ProviderVersion** | Pointer to **string** |  | [optional] 
**RequiredGroupMembership** | Pointer to **[]string** |  | [optional] 
**Skin** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### GetAccountId

`func (o *AccountDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountDetails) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *AccountDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *AccountDetails) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetAccountType

`func (o *AccountDetails) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountDetails) GetAccountTypeOk() (string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountType

`func (o *AccountDetails) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountType

`func (o *AccountDetails) SetAccountType(v string)`

SetAccountType gets a reference to the given string and assigns it to the AccountType field.

### GetChallengeDestructiveActions

`func (o *AccountDetails) GetChallengeDestructiveActions() bool`

GetChallengeDestructiveActions returns the ChallengeDestructiveActions field if non-nil, zero value otherwise.

### GetChallengeDestructiveActionsOk

`func (o *AccountDetails) GetChallengeDestructiveActionsOk() (bool, bool)`

GetChallengeDestructiveActionsOk returns a tuple with the ChallengeDestructiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChallengeDestructiveActions

`func (o *AccountDetails) HasChallengeDestructiveActions() bool`

HasChallengeDestructiveActions returns a boolean if a field has been set.

### SetChallengeDestructiveActions

`func (o *AccountDetails) SetChallengeDestructiveActions(v bool)`

SetChallengeDestructiveActions gets a reference to the given bool and assigns it to the ChallengeDestructiveActions field.

### GetCloudProvider

`func (o *AccountDetails) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AccountDetails) GetCloudProviderOk() (string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCloudProvider

`func (o *AccountDetails) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProvider

`func (o *AccountDetails) SetCloudProvider(v string)`

SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.

### GetEnvironment

`func (o *AccountDetails) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AccountDetails) GetEnvironmentOk() (string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnvironment

`func (o *AccountDetails) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironment

`func (o *AccountDetails) SetEnvironment(v string)`

SetEnvironment gets a reference to the given string and assigns it to the Environment field.

### GetName

`func (o *AccountDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDetails) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AccountDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AccountDetails) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPermissions

`func (o *AccountDetails) GetPermissions() map[string][]string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccountDetails) GetPermissionsOk() (map[string][]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissions

`func (o *AccountDetails) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissions

`func (o *AccountDetails) SetPermissions(v map[string][]string)`

SetPermissions gets a reference to the given map[string][]string and assigns it to the Permissions field.

### GetPrimaryAccount

`func (o *AccountDetails) GetPrimaryAccount() bool`

GetPrimaryAccount returns the PrimaryAccount field if non-nil, zero value otherwise.

### GetPrimaryAccountOk

`func (o *AccountDetails) GetPrimaryAccountOk() (bool, bool)`

GetPrimaryAccountOk returns a tuple with the PrimaryAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimaryAccount

`func (o *AccountDetails) HasPrimaryAccount() bool`

HasPrimaryAccount returns a boolean if a field has been set.

### SetPrimaryAccount

`func (o *AccountDetails) SetPrimaryAccount(v bool)`

SetPrimaryAccount gets a reference to the given bool and assigns it to the PrimaryAccount field.

### GetProviderVersion

`func (o *AccountDetails) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *AccountDetails) GetProviderVersionOk() (string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProviderVersion

`func (o *AccountDetails) HasProviderVersion() bool`

HasProviderVersion returns a boolean if a field has been set.

### SetProviderVersion

`func (o *AccountDetails) SetProviderVersion(v string)`

SetProviderVersion gets a reference to the given string and assigns it to the ProviderVersion field.

### GetRequiredGroupMembership

`func (o *AccountDetails) GetRequiredGroupMembership() []string`

GetRequiredGroupMembership returns the RequiredGroupMembership field if non-nil, zero value otherwise.

### GetRequiredGroupMembershipOk

`func (o *AccountDetails) GetRequiredGroupMembershipOk() ([]string, bool)`

GetRequiredGroupMembershipOk returns a tuple with the RequiredGroupMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequiredGroupMembership

`func (o *AccountDetails) HasRequiredGroupMembership() bool`

HasRequiredGroupMembership returns a boolean if a field has been set.

### SetRequiredGroupMembership

`func (o *AccountDetails) SetRequiredGroupMembership(v []string)`

SetRequiredGroupMembership gets a reference to the given []string and assigns it to the RequiredGroupMembership field.

### GetSkin

`func (o *AccountDetails) GetSkin() string`

GetSkin returns the Skin field if non-nil, zero value otherwise.

### GetSkinOk

`func (o *AccountDetails) GetSkinOk() (string, bool)`

GetSkinOk returns a tuple with the Skin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkin

`func (o *AccountDetails) HasSkin() bool`

HasSkin returns a boolean if a field has been set.

### SetSkin

`func (o *AccountDetails) SetSkin(v string)`

SetSkin gets a reference to the given string and assigns it to the Skin field.

### GetType

`func (o *AccountDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountDetails) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *AccountDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *AccountDetails) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


