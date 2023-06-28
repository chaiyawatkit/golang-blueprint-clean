package http

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang-blueprint-clean/app/constants"
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/errors"
	"golang-blueprint-clean/app/layers/usecases/users"
	"golang-blueprint-clean/app/utils"
	"net/http"
	"strconv"
)

type AuthMiddleware interface {
	Authentication(c *gin.Context)
}

type authMiddleware struct {
	UsersUseCase users.UseCase
}

func (middleware *authMiddleware) Authentication(c *gin.Context) {
	jwtAccessToken := c.GetHeader(constants.XFinPlusAuth)

	if jwtAccessToken == "" {
		err := errors.Unauthorized{Message: constants.MissingFinPlusAuth}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, http.StatusUnauthorized, err, humanMsg)
		return
	}
	newJwtAccessToken, err, errMsg := middleware.UsersUseCase.Authentication(c.Request.Context(), jwtAccessToken)
	if (err != nil && errMsg != nil) || newJwtAccessToken == nil {
		err = errors.Unauthorized{Message: constants.FailToAuthenticate}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, http.StatusUnauthorized, err, humanMsg)
		return
	}

	publicKey, err := utils.GetPublicKey()
	if err != nil {
		err = errors.Unauthorized{Message: err.Error()}
		humanMsg := utils.GetHumanErrorCode(errMsg.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(*newJwtAccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		err = errors.Unauthorized{Message: err.Error()}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	if jwtToken == nil {
		err = errors.Unauthorized{Message: "invalid JWT token"}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	jwtUserUuid := claims["uuid"]
	jwtUserUuidStr := fmt.Sprintf("%v", jwtUserUuid)
	if jwtUserUuidStr == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT uuid '%s' not found", jwtUserUuidStr)}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	jwtStatus := fmt.Sprintf("%v", claims["status"])
	jwtStatusValue := fmt.Sprintf("%v", jwtStatus)
	if jwtStatusValue == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT status '%v' not found", jwtStatus)}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}
	jwtSegment := fmt.Sprintf("%v", claims["segment"])
	jwtSegmentValue := fmt.Sprintf("%v", jwtSegment)
	if jwtStatusValue == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT segment '%v' not found", jwtSegmentValue)}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	jwtIssueAt := fmt.Sprintf("%v", claims["iat"])
	jwtIssueAtValue, err := strconv.ParseFloat(jwtIssueAt, 64)
	if err != nil {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT iat '%v' not found", jwtIssueAt)}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	jwtGuid := fmt.Sprintf("%v", claims["guid"])
	if claims["guid"] == nil || jwtGuid == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT guid '%v' not found", jwtGuid)}
		humanMsg := utils.GetHumanErrorCode(err.Error())
		utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
		return
	}

	//values, err := utils.DecryptAES(jwtGuid, env.EncryptKey)
	//
	//uid, err := strconv.ParseUint(*values, 10, 32)
	//if err != nil {
	//
	//	humanMsg := utils.GetHumanErrorCode(err.Error())
	//	utils.JSONErrorCodeResponse(c, 400, err, humanMsg)
	//	return
	//}

	jwtData := entities.JwtData{
		Uuid:    jwtUserUuidStr,
		Status:  jwtStatusValue,
		IssueAt: jwtIssueAtValue,
		Segment: jwtSegmentValue,
	}
	c.Set(constants.JWTDataKey, jwtData)
	c.Header(constants.XFinPlusAuth, *newJwtAccessToken)
	c.Next()
}

func InitAuthMiddleware(usersUseCase users.UseCase) AuthMiddleware {
	return &authMiddleware{
		UsersUseCase: usersUseCase,
	}
}
