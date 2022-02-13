package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"shaps.api/domain/setting"
)

func ValidateToken(s setting.Setting) gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt取得
		bearer := c.GetHeader("Authorization")
		if len(bearer) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}
		j := strings.Fields(bearer)[1]

		// jwk取得
		set, err := jwk.Fetch(context.Background(), s.B2CJwkUri)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		// 署名検証
		token, err := jwt.Parse(
			[]byte(j),
			jwt.WithKeySet(set),
			jwt.InferAlgorithmFromKey(true),
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// 期限チェック
		if token.Expiration().Unix() < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}
		// 発行者チェック
		if strings.Split(token.Issuer(), "/")[3] != s.B2CTenantId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// 対象者チェック
		if token.Audience()[0] != s.B2CClientId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		c.Set("userId", token.Subject())
	}
}
