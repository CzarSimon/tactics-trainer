package authentication_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/client/rpc"
	"github.com/CzarSimon/httputil/crypto"
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
	svc := setupAuthService()
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

	user, found, err := svc.UserRepo.FindByUsername(ctx, body.Username)
	assert.NoError(err)
	assert.True(found)
	assert.Equal(body.Username, user.Username)
	assert.NotEqual(body.Password, user.Credentials.Password)
}

func setupAuthService() *service.AuthenticationService {
	db := testutil.InMemoryDB(true, "../../../resources/db/sqlite")
	userRepo := repository.NewUserRepository(db)

	k1, _ := models.ParseKeyEncryptionKey("0:2557BC625F244FAFFDC817A7DDB6A20E09B4C1200C4A2B3ABAA49BA968476452")
	k2, _ := models.ParseKeyEncryptionKey("1:BCF39E65064252E1C4043A576ED9EA9AC84BCAFECEF9588D18479DC0F8A2AEA9")
	keys := []models.KeyEncryptionKey{k1, k2}

	kekRepo := repository.NewKeyEncryptionKeyRepository(keys, db)
	cipher := &service.Cipher{
		KEKRepo: kekRepo,
	}

	return &service.AuthenticationService{
		UserRepo: userRepo,
		Cipher:   cipher,
		Hasher:   crypto.DefaultScryptHasher(),
	}
}

func setupRouter(svc *service.AuthenticationService) http.Handler {
	r := httputil.NewRouter("iam-server", func() error {
		return nil
	})
	authentication.AttachController(svc, r)
	return r
}
