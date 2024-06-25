package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joshsoftware/peerly-backend/internal/pkg/apperrors"
	"github.com/joshsoftware/peerly-backend/internal/pkg/config"
	"github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	"github.com/joshsoftware/peerly-backend/internal/repository"
	logger "github.com/sirupsen/logrus"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	ValidatePeerly(ctx context.Context, authToken string) (data dto.ValidateResp, err error)
	GetIntranetUserData(ctx context.Context, req dto.GetIntranetUserDataReq) (data dto.IntranetApiResp, err error)
	LoginUser(ctx context.Context, u dto.IntranetApiResp) (dto.LoginUserResp, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (cs *service) ValidatePeerly(ctx context.Context, authToken string) (data dto.ValidateResp, err error) {
	client := &http.Client{}
	validationReq, err := http.NewRequest("GET", "http://localhost:33001/intranet/validate", nil)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	validationReq.Header.Add("Authorization", authToken)
	validationReq.Header.Add("PeerlyCode", "peerly")
	resp, err := client.Do(validationReq)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = apperrors.IntranetValidationFailed
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = apperrors.JSONParsingErrorResp
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		err = apperrors.JSONParsingErrorResp
		return
	}

	return
}

func (cs *service) GetIntranetUserData(ctx context.Context, req dto.GetIntranetUserDataReq) (data dto.IntranetApiResp, err error) {

	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:33001/intranet/getuser/%d", req.UserId)
	intranetReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}

	intranetReq.Header.Add("Authorization", req.Token)
	resp, err := client.Do(intranetReq)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = apperrors.JSONParsingErrorResp
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		err = apperrors.JSONParsingErrorResp
		return
	}

	return
}

func (cs *service) LoginUser(ctx context.Context, u dto.IntranetApiResp) (dto.LoginUserResp, error) {
	var resp dto.LoginUserResp
	resp.NewUserCreated = false
	user, err := cs.userRepo.GetUserByEmail(ctx, u.Email)
	if err == apperrors.InternalServerError {
		return resp, err
	}

	if err == apperrors.UserNotFound {
		//get organization id
		orgId := 1
		//get grade id
		gradeId, err := cs.userRepo.GetGradeByName(ctx, u.Grade)
		if err != nil {
			return resp, err
		}
		// gradeId := 1
		//reward_quota_balance from organization
		reward_quota_balance := 10

		//get role by name
		roleId, err := cs.userRepo.GetRoleByName(ctx, "user")
		if err != nil {
			err = apperrors.InternalServerError
			return resp, err
		}

		var userData dto.RegisterUser
		userData.User = u
		userData.OrgId = orgId
		userData.GradeId = gradeId
		userData.RewardQuotaBalance = reward_quota_balance
		userData.RoleId = roleId

		//register user
		user, err = cs.userRepo.CreateNewUser(ctx, userData)
		if err != nil {
			err = apperrors.InternalServerError
			return resp, err
		}

		resp.NewUserCreated = true
	}

	//login user

	expirationTime := time.Now().Add(time.Hour * 5)

	claims := &dto.Claims{
		Id:     user.Id,
		RoleId: user.RoleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// var jwtKey = []byte("secret_key")
	var jwtKey = config.JWTKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error generating authtoken")
		err = apperrors.InternalServerError
		return resp, err
	}

	resp.User = user
	resp.AuthToken = tokenString

	return resp, nil

}
