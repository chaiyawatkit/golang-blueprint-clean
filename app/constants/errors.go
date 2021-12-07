package constants

const (
	FailToCreateDB                          = "fail to create database"
	DuplicateSlug                           = "fail duplicated slug"
	IsWeakPassword                          = "fail password is too weak need minimum 8 characters and one more a special character, uppercase, lowercase, number"
	IsShortPassword                         = "too short password"
	InvalidPhoneNumber                      = "phone number invalid"
	InValidEmail                            = "invalid email address"
	InValidUuid                             = "uuid must be a uud v4 format"
	RecordNotFound                          = "fail record not found"
	EmptyParameter                          = "please specify an parameter"
	FailToGetDataFromDB                     = "fail to get data from database"
	AuthenticationFailed                    = "authentication failed"
	FailToUpdateDB                          = "fail to update database"
	FailToDeleteDB                          = "fail to delete database"
	InValidUid                              = "uid invalid"
	FailToGrabPublicKey                     = "fail to get public key"
	FailUnExpectedSigningMethod             = "fail unexpected signing method"
	FailJwtTokenCannotBeClaimed             = "fail jwt token cannot be claimed"
	NotFoundUserWithGivenAccessToken        = "not found user with given access token"
	GivenUserUuidIsUnmatchedWithJwtUserUuid = "given user uuid is unmatched with jwt user uuid"
	FailToLogout                            = "fail to logout"
	OldPasswordNotMatch                     = "old password not match"
	PasswordSameAsBefore                    = "password same as before"
	DuplicateUsername                       = "username already exist"
	ProfileAlreadyExist                     = "profile already exist"
	UsernameAlreadyExist                    = "username taken - please enter another username"
	InvalidUsername                         = "your username must be a-z,A-z,0-9 within 4-20 characters"
	InvalidDisplayName                      = "your display within 2-20 characters"
	InvalidAge                              = "over 10 years old"
	InvalidFormatBirthDay                   = "birthday format yyyy-mm-dd Ex.2010-01-01"
	DuplicateFriendsRelation                = "duplicated friends relation"
	InValidRequestUid                       = "request uid must be the same as uid"
	InvalidChannelType                      = "invalid channel type"
	AccountAlreadyActive                    = "this account is already active"
	NotEmailAccount                         = "this account is not an email account"
	AccountAlreadyActivate                  = "account already activate"
	InvalidActivateCode                     = "invalid activate code"
	CodeExpire                              = "code expire"
	YourAccountWasDeleted                   = "your account was deleted"
)

func MissingWithKey(key string) string {
	return "required field is missing: field '" + key + "' is required."
}

var (
	HumanErrorCode = map[string]interface{}{
		"default":             99,
		InvalidUsername:       13,
		PasswordSameAsBefore:  24,
		IsWeakPassword:        11,
		InvalidPhoneNumber:    1,
		InvalidChannelType:    28,
		InvalidAge:            16,
		InvalidDisplayName:    26,
		InvalidFormatBirthDay: 27,
		YourAccountWasDeleted: 36,
	}

	HumanSuccessCode = map[string]interface{}{
		"default":           100,
		"usernameAvailable": 18,
	}
)
