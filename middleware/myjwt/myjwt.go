package myjwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/foxkillerli/IELTS-assist/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func NeedJwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
func NoNeedJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "" {
			j := NewJWT()
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			if err == nil {
				c.Set("claims", claims)
			}
		}
	}
}

// 生成令牌
func GenerateToken(c *gin.Context, user models.User) string {
	j := JWT{
		[]byte("ielts"),
	}
	claims := CustomClaims{
		strconv.FormatUint(uint64(user.BaseModel.ID), 10),
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 604800), // 过期时间 一小时
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return ""
	}
	return token
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("token not active yet")
	TokenMalformed   error  = errors.New("that's not even a token")
	TokenInvalid     error  = errors.New("couldn't handle this token")
	SignKey          string = "eStarGo"
)

type CustomClaims struct {
	ID string `json:"user_id"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func FromTokenGetStudentId(token string) (studentId int, error string) {
	// 从token中获取用户id
	jwt := NewJWT()
	if claims, err := jwt.ParseToken(token); err != nil {
		return 0, ""
	} else {
		if studentId, err := strconv.Atoi(claims.ID); err != nil {
			return 0, err.Error()
		} else {
			return studentId, ""
		}
	}
}
