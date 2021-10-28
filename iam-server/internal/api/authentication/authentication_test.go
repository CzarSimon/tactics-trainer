package authentication_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/client/rpc"
	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/api/authentication"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	assert := assert.New(t)
	svc, jwtVerifier := setupAuthService()
	router := setupRouter(svc)
	ctx := context.Background()

	body := models.AuthenticationRequest{
		Username: "valid-username",
		Password: "valid-password",
	}

	_, found, err := svc.UserRepo.FindByUsername(ctx, body.Username)
	assert.NoError(err)
	assert.False(found)

	req, _ := rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/signup", body)
	res := testutil.PerformRequest(router, req)
	assert.Equal(http.StatusOK, res.Code)
	var resBody models.AuthenticationResponse
	err = rpc.DecodeJSON(res.Result(), &resBody)
	assert.NoError(err)
	assert.Len(resBody.User.ID, 36)
	assert.Equal(body.Username, resBody.User.Username)
	assert.Equal(models.UserRole, resBody.User.Role)
	assert.Empty(resBody.User.Credentials)

	assert.NotEmpty(resBody.Token)
	jwtUser, err := jwtVerifier.Verify(resBody.Token)
	assert.NoError(err)
	assert.True(jwtUser.HasRole(models.UserRole))
	assert.Equal(resBody.User.ID, jwtUser.ID)

	user, found, err := svc.UserRepo.FindByUsername(ctx, body.Username)
	assert.NoError(err)
	assert.True(found)
	assert.Equal(body.Username, user.Username)
	assert.NotEqual(body.Password, user.Credentials.Password)

	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/signup", body)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusConflict, res.Code)

	body.Password = "toshort"
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/signup", body)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusBadRequest, res.Code)

	body.Username = ""
	body.Password = "valid-password"
	req, _ = rpc.NewClient(time.Second).CreateRequest(http.MethodPost, "/v1/signup", body)
	res = testutil.PerformRequest(router, req)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func setupAuthService() (*service.AuthenticationService, jwt.Verifier) {
	db := testutil.InMemoryDB(true, "../../../resources/db/sqlite")
	userRepo := repository.NewUserRepository(db)

	k1, _ := models.ParseKeyEncryptionKey("0:2557BC625F244FAFFDC817A7DDB6A20E09B4C1200C4A2B3ABAA49BA968476452")
	k2, _ := models.ParseKeyEncryptionKey("1:BCF39E65064252E1C4043A576ED9EA9AC84BCAFECEF9588D18479DC0F8A2AEA9")
	keys := []models.KeyEncryptionKey{k1, k2}

	kekRepo := repository.NewKeyEncryptionKeyRepository(keys, db)
	cipher := &service.Cipher{
		KEKRepo: kekRepo,
	}

	jwtCreds := jwt.Credentials{
		Issuer: "tactics-trainer/iam-service",
		Secret: "614f10d529be2dd62b3554160c5247d5",
	}

	return &service.AuthenticationService{
		UserRepo:      userRepo,
		Cipher:        cipher,
		Hasher:        crypto.DefaultScryptHasher(),
		Issuer:        jwt.NewIssuer(jwtCreds),
		TokenLifetime: time.Hour * 24 * 7,
	}, jwt.NewVerifier(jwtCreds, time.Minute)
}

func setupRouter(svc *service.AuthenticationService) http.Handler {
	r := httputil.NewRouter("iam-server", func() error {
		return nil
	})
	authentication.AttachController(svc, r)
	return r
}
