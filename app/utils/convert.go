package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-blueprint-clean/app/constants"
)

func FromContextToGinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(constants.GinContextKey)
	if ginContext == nil {
		return nil, errors.New("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("wrong type is in GinContextKey")
	}

	return gc, nil
}
