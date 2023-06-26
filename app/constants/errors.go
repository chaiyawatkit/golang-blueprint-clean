package constants

const (
	FailToAuthenticate          = "fail to authenticate"
	FailUnExpectedSigningMethod = "fail unexpected signing method"
	Unauthorized                = "unauthorized"
	FailNotFound                = "fail record not found"
	InValidSegment              = "invalid Segment"
	SegmentType                 = "invalid type of segment in list [pb, general, wisdom]"
	EmptyParameter              = "please specify an parameter"
	FailToGet                   = "fail to get"
)

func MissingWithKey(key string) string {
	return "required field is missing: field '" + key + "' is required."
}

var (
	HumanErrorCode = map[string]interface{}{
		"default":                             99,
		Unauthorized:                          401,
		"fail jwt token cannot be claimed":    401,
		"x-finplus-auth is missing":           401,
		FailUnExpectedSigningMethod:           401,
		"unauthorized, api gateway is denied": 401,
		FailToAuthenticate:                    401,
		"illegal base64 data at input byte 100; see https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to retrieve a valid ID token": 401,
		"EOF; see https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to retrieve a valid ID token":                                   401,
		"email-sign-in-authentication failed": 401,
		"ID token must be a non-empty string": 401,
		FailNotFound:                          20,
		SegmentType:                           21,
		InValidSegment:                        22,
	}

	HumanSuccessCode = map[string]interface{}{
		"default":           100,
		"usernameAvailable": 18,
	}
)
