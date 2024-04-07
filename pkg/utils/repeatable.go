package repeatable

import (
	"os"
	"time"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

var (
	PublicFilePath = "./../../public"
)

func CrateDir() error {
	return os.MkdirAll(PublicFilePath, os.ModePerm)
}

//func TokenClaims(token, secretKey string) (jwt.MapClaims, error) {
//
//	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(secretKey), nil
//	})
//
//	if err != nil {
//		fmt.Println("err", err)
//		return nil, appresult.ErrMissingParam
//	}
//
//	claims, ok := decoded.Claims.(jwt.MapClaims)
//
//	if !ok {
//		// TODO tokenin omrini test etmeli
//		return nil, appresult.ErrInternalServer
//	}
//
//	return claims, nil
//}
