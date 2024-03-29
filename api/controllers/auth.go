package controllers

import (
	"OfflineSearchEngine/configs"
	"OfflineSearchEngine/internals/db"
	"OfflineSearchEngine/internals/utils"
	"OfflineSearchEngine/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTHandler struct {
	configs *configs.Configs
	db      *db.DatabaseHandler
	userIds map[string]uint
}

func NewJWTHandler(configs *configs.Configs, db *db.DatabaseHandler) *JWTHandler {
	return &JWTHandler{configs: configs, db: db, userIds: make(map[string]uint)}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (handler *JWTHandler) SignInHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbUser, err := handler.db.ReadUserByUsername(user.Username)
	if err != nil || !utils.ComparePasswords(dbUser.Password, []byte(user.Password)) {
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(handler.configs.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	handler.userIds[tokenString] = dbUser.ID
	c.JSON(http.StatusOK, jwtOutput)
}

func (handler *JWTHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(handler.configs.JWTSecret), nil
			})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthoized"})
			return
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthoized"})
		}
		if id, ok := handler.userIds[tokenValue]; ok {
			c.Set("userId", int(id))
		}
		c.Next()
	}
}

func (handler *JWTHandler) RefreshHandler(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenValue, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(handler.configs.JWTSecret), nil
		})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if tkn == nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) >
		30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not expired yet"})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		claims)
	tokenString, err := token.SignedString(handler.configs.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}
