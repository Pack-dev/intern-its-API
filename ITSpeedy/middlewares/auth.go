package middlewares

import (
	"errors"
	"fmt"
	"strings"
	db "templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Headerdata Header data
type Headerdata struct {
	Result struct {
		ID          string `bson:"_id,omitempty" json:"id"`
		Username    string `bson:"username" json:"username"`
		Password    string `bson:"password" json:"password"`
		Name        string `bson:"name" json:"name"`
		LastName    string `bson:"last_name" json:"last_name"`
		PhoneNumber string `bson:"phone_number" json:"phone_number"`
		FullName    string `bson:"full_name" json:"full_name"`
		ImgUrl      string `bson:"img_url" json:"img_url"`
		Active      bool   `bson:"active" json:"active"`
	} `json:"result"`
	jwt.StandardClaims
}
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtServices struct {
	secretKey string
	issure    string
}

type authCustomClaims struct {
	ID       string `bson:"_id" json:"id"`
	Prefix   string `bson:"prefix" json:"prefix"`
	Username string `bson:"username" json:"username"`
	Name     string `bson:"name" json:"name"`
	LastName string `bson:"last_name" json:"last_name"`
	FullName string `bson:"full_name" json:"full_name"`
	jwt.StandardClaims
}
type payload struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// AuthRequired to check the authentication key in HTTP Header
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "" && c.GetString("user_id") != "" {
			c.Next()
		} else {
			authorizationHeader := c.GetHeader("Authorization")
			if authorizationHeader == "" {
				c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
				return
			}
			AuthHeader(c)
		}
		// //service
		// service, _ := b64.StdEncoding.DecodeString(c.Params.ByName("service"))
		// c.Param("service") := string(service)
		c.Next()
	}
}

// AuthHeader auth header
func AuthHeader(c *gin.Context) {
	bearToken, errToken := ExtractToken(c)
	if errToken != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"code": 401,
			"msg":  fmt.Sprint(errToken),
		})
		return
	}
	tokenString := bearToken[1]
	var claims = &authCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("BILL0078"), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"code": 401,
			"msg":  "AuthenHeader Fail !!",
		})
		return
	}
	if claims.Username == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"code": 401,
			"msg":  "AuthenHeader Faild !!",
		})
		return
	}

	c.Set("username", claims.Username)
	c.Set("name", claims.Name)
	c.Set("user_id", claims.ID)
}

// ExtractToken Extract Token
func ExtractToken(c *gin.Context) ([]string, error) {
	var token []string
	bearToken := c.Request.Header.Get("Authorization")
	if strings.Index(bearToken, "Bearer") == -1 {
		return token, errors.New("ERROR BEARER TOKEN")
	}
	token = strings.Split(bearToken, "Bearer ")
	return token, nil
}

func GenToken(c *gin.Context, body model.UserModelS) (string, error) {
	mySigningKey := []byte("BILL0078")

	// Create the Claims
	claims := &authCustomClaims{
		body.ID,
		body.Prefix,
		body.Username,
		body.Name,
		body.LastName,
		body.FullName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "system",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	c.Set("token", ss)
	return ss, nil
}

func AuthToken(resource *db.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("token")
		if authorizationHeader == "" {
			return
		}
		filter := bson.M{
			"api_token": authorizationHeader,
		}
		user := model.UserModelS{}
		if err := repo.GetOneStatement(resource, "user", filter, nil, &user); err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("username", user.Username)
		c.Set("name", user.Name)
		c.Set("user_id", user.ID)
		c.Next()
	}
}
