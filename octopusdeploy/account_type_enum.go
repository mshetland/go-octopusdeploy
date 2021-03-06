// Code generated by go-enum
// DO NOT EDIT!

package octopusdeploy

import (
	"fmt"
	"strings"
)

const (
	// None is a AccountType of type None
	None AccountType = iota
	// UsernamePassword is a AccountType of type UsernamePassword
	UsernamePassword
	// SshKeyPair is a AccountType of type SshKeyPair
	SshKeyPair
	// AzureSubscription is a AccountType of type AzureSubscription
	AzureSubscription
	// AzureServicePrincipal is a AccountType of type AzureServicePrincipal
	AzureServicePrincipal
	// AmazonWebServicesAccount is a AccountType of type AmazonWebServicesAccount
	AmazonWebServicesAccount
	// AmazonWebServicesRoleAccount is a AccountType of type AmazonWebServicesRoleAccount
	AmazonWebServicesRoleAccount
)

const _AccountTypeName = "NoneUsernamePasswordSshKeyPairAzureSubscriptionAzureServicePrincipalAmazonWebServicesAccountAmazonWebServicesRoleAccount"

var _AccountTypeNames = []string{
	_AccountTypeName[0:4],
	_AccountTypeName[4:20],
	_AccountTypeName[20:30],
	_AccountTypeName[30:47],
	_AccountTypeName[47:68],
	_AccountTypeName[68:92],
	_AccountTypeName[92:120],
}

// AccountTypeNames returns a list of possible string values of AccountType.
func AccountTypeNames() []string {
	tmp := make([]string, len(_AccountTypeNames))
	copy(tmp, _AccountTypeNames)
	return tmp
}

var _AccountTypeMap = map[AccountType]string{
	0: _AccountTypeName[0:4],
	1: _AccountTypeName[4:20],
	2: _AccountTypeName[20:30],
	3: _AccountTypeName[30:47],
	4: _AccountTypeName[47:68],
	5: _AccountTypeName[68:92],
	6: _AccountTypeName[92:120],
}

// String implements the Stringer interface.
func (x AccountType) String() string {
	if str, ok := _AccountTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("AccountType(%d)", x)
}

var _AccountTypeValue = map[string]AccountType{
	_AccountTypeName[0:4]:                     0,
	strings.ToLower(_AccountTypeName[0:4]):    0,
	_AccountTypeName[4:20]:                    1,
	strings.ToLower(_AccountTypeName[4:20]):   1,
	_AccountTypeName[20:30]:                   2,
	strings.ToLower(_AccountTypeName[20:30]):  2,
	_AccountTypeName[30:47]:                   3,
	strings.ToLower(_AccountTypeName[30:47]):  3,
	_AccountTypeName[47:68]:                   4,
	strings.ToLower(_AccountTypeName[47:68]):  4,
	_AccountTypeName[68:92]:                   5,
	strings.ToLower(_AccountTypeName[68:92]):  5,
	_AccountTypeName[92:120]:                  6,
	strings.ToLower(_AccountTypeName[92:120]): 6,
}

// ParseAccountType attempts to convert a string to a AccountType
func ParseAccountType(name string) (AccountType, error) {
	if x, ok := _AccountTypeValue[name]; ok {
		return x, nil
	}
	return AccountType(0), fmt.Errorf("%s is not a valid AccountType, try [%s]", name, strings.Join(_AccountTypeNames, ", "))
}

// MarshalText implements the text marshaller method
func (x *AccountType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *AccountType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseAccountType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
