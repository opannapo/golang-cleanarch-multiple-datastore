package services

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

//AuthService general AuthService
type AuthService struct {
	Repositories *repositories.RepositoryInjection
}

//ValidateCredential service to validate credential then return jwt token
func (instance *AuthService) ValidateCredential(param *param.AuthParam) (result entities.Credential, token string, err error) {
	key := param.Key
	signature := param.Signature

	result, err = instance.Repositories.MysqlCredentialRepo.GetByKeySignature(key, signature)

	if err != nil {
		return
	}

	keyJwt := viper.GetString("jwtKey")
	//exp := time.Now().Unix() + (60 * (60 * 1)) //1 jam
	exp := time.Now().Unix() + (60 * (60 * 2)) //2 jam
	tokenGenerated := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": result.User.ID,
		"exp": exp,
	})
	token, err = tokenGenerated.SignedString([]byte(keyJwt))
	return
}

//NewInstanceAuthService new instance of AuthService
func NewInstanceAuthService(repositories *repositories.RepositoryInjection) AuthServices {
	return &AuthService{
		Repositories: repositories,
	}
}
