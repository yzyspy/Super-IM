package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义自定义声明结构体，可以包含标准声明和自定义字段
type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 密钥，用于签名和验证 Token，请务必保密且足够复杂
var secretKey = []byte("your-secret-key")

// GenerateJWT 生成 JWT Token
func GenerateJWT(userID, username string) (string, error) {
	// 设置 Token 过期时间，例如 24 小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建自定义声明
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			Issuer:    "your-app-name",       // 签发者
			IssuedAt:  time.Now().Unix(),     // 签发时间
		},
	}

	// 使用 HS256 算法创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并获取字符串形式的 Token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析和验证 JWT Token
func ParseJWT(tokenString string) (*CustomClaims, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 Token 是否有效并提取声明
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// 示例使用
func main() {
	// 生成 Token
	userID := "12345"
	username := "john_doe"
	token, err := GenerateJWT(userID, username)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}
	fmt.Printf("Generated Token: %s\n", token)

	// 解析和验证 Token
	claims, err := ParseJWT(token)
	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return
	}
	fmt.Printf("UserID: %s, Username: %s, ExpiresAt: %d\n", claims.UserID, claims.Username, claims.ExpiresAt)
}
