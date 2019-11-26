# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNonExpired** | Pointer to **bool** |  | [optional] 
**AccountNonLocked** | Pointer to **bool** |  | [optional] 
**AllowedAccounts** | Pointer to **[]string** |  | [optional] 
**Authorities** | Pointer to [**[]GrantedAuthority**](GrantedAuthority.md) |  | [optional] 
**CredentialsNonExpired** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### GetAccountNonExpired

`func (o *User) GetAccountNonExpired() bool`

GetAccountNonExpired returns the AccountNonExpired field if non-nil, zero value otherwise.

### GetAccountNonExpiredOk

`func (o *User) GetAccountNonExpiredOk() (bool, bool)`

GetAccountNonExpiredOk returns a tuple with the AccountNonExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountNonExpired

`func (o *User) HasAccountNonExpired() bool`

HasAccountNonExpired returns a boolean if a field has been set.

### SetAccountNonExpired

`func (o *User) SetAccountNonExpired(v bool)`

SetAccountNonExpired gets a reference to the given bool and assigns it to the AccountNonExpired field.

### GetAccountNonLocked

`func (o *User) GetAccountNonLocked() bool`

GetAccountNonLocked returns the AccountNonLocked field if non-nil, zero value otherwise.

### GetAccountNonLockedOk

`func (o *User) GetAccountNonLockedOk() (bool, bool)`

GetAccountNonLockedOk returns a tuple with the AccountNonLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountNonLocked

`func (o *User) HasAccountNonLocked() bool`

HasAccountNonLocked returns a boolean if a field has been set.

### SetAccountNonLocked

`func (o *User) SetAccountNonLocked(v bool)`

SetAccountNonLocked gets a reference to the given bool and assigns it to the AccountNonLocked field.

### GetAllowedAccounts

`func (o *User) GetAllowedAccounts() []string`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *User) GetAllowedAccountsOk() ([]string, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedAccounts

`func (o *User) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### SetAllowedAccounts

`func (o *User) SetAllowedAccounts(v []string)`

SetAllowedAccounts gets a reference to the given []string and assigns it to the AllowedAccounts field.

### GetAuthorities

`func (o *User) GetAuthorities() []GrantedAuthority`

GetAuthorities returns the Authorities field if non-nil, zero value otherwise.

### GetAuthoritiesOk

`func (o *User) GetAuthoritiesOk() ([]GrantedAuthority, bool)`

GetAuthoritiesOk returns a tuple with the Authorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthorities

`func (o *User) HasAuthorities() bool`

HasAuthorities returns a boolean if a field has been set.

### SetAuthorities

`func (o *User) SetAuthorities(v []GrantedAuthority)`

SetAuthorities gets a reference to the given []GrantedAuthority and assigns it to the Authorities field.

### GetCredentialsNonExpired

`func (o *User) GetCredentialsNonExpired() bool`

GetCredentialsNonExpired returns the CredentialsNonExpired field if non-nil, zero value otherwise.

### GetCredentialsNonExpiredOk

`func (o *User) GetCredentialsNonExpiredOk() (bool, bool)`

GetCredentialsNonExpiredOk returns a tuple with the CredentialsNonExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCredentialsNonExpired

`func (o *User) HasCredentialsNonExpired() bool`

HasCredentialsNonExpired returns a boolean if a field has been set.

### SetCredentialsNonExpired

`func (o *User) SetCredentialsNonExpired(v bool)`

SetCredentialsNonExpired gets a reference to the given bool and assigns it to the CredentialsNonExpired field.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName gets a reference to the given string and assigns it to the FirstName field.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName gets a reference to the given string and assigns it to the LastName field.

### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() ([]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles gets a reference to the given []string and assigns it to the Roles field.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername gets a reference to the given string and assigns it to the Username field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


