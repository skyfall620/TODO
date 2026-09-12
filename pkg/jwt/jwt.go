package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTdata struct {
	UserId int64
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTdata) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": data.UserId,
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, err
}

func (j *JWT) Parse(token string) (bool, *JWTdata) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil
	}

	userId := int64(claims["userId"].(float64))

	return t.Valid, &JWTdata{
		UserId: userId,
	}
}
