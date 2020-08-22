package general

import (
	"app/app/v1/apis/param"
	"app/app/v1/entities"
	"app/app/v1/injection/repositories"
	"app/app/v1/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type AuthService struct {
	Repositories *repositories.RepositoryInjection
}

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
		"uid": result.User.Id,
		"exp": exp,
	})
	token, err = tokenGenerated.SignedString([]byte(keyJwt))
	return
}

func NewInstanceAuthService(repositories *repositories.RepositoryInjection) services.AuthServices {
	return &AuthService{
		Repositories: repositories,
	}
}
