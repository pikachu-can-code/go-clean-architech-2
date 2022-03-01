package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/tokenprovider"
)

type jwtProvider struct {
	secret string
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (c *myClaims) Valid() error {
	var leeway int64 = 10000000
	c.StandardClaims.IssuedAt -= leeway
	valid := c.StandardClaims.Valid()
	c.StandardClaims.IssuedAt += leeway
	return valid
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int64) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Phi Khanh đẹp trai thanh lịch vô địch vũ trụ!",
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		AuthId:  int64(data.UserID),
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now().Unix(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}
