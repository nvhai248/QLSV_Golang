package middleware

import (
	"errors"
	"fmt"
	"strings"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/component/tokenprovider/jwt"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong auth header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := studentstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.DetailStudent(c.Request.Context(), payload.UserId)

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		// save user in context
		c.Set(common.CurrentStudent, user)
		c.Next()
	}
}
