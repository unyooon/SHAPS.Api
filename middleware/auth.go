package middleware

import (
	"context"
	"errors"
	"log"
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
		j := strings.Fields(c.GetHeader("Authorization"))[1]
		if len(j) == 0 {
			log.Print("Error: not found authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// jwk取得
		set, err := jwk.Fetch(context.Background(), s.B2CJwkUri)
		if err != nil {
			log.Printf("Error: failed to parse JWK: %s", err)
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
			log.Printf("Error: failed to parse payload: %s", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// 期限チェック
		if token.Expiration().Unix() < time.Now().Unix() {
			log.Print("Error: expired token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}
		// 発行者チェック
		if strings.Split(token.Issuer(), "/")[3] != s.B2CTenantId {
			log.Print("Error: invalid grant")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// 対象者チェック
		if token.Audience()[0] != s.B2CClientId {
			log.Print("Error: invalid grant")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		c.Set("userId", token.Subject())
	}
}
