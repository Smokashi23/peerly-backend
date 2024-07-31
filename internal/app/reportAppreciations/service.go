package reportappreciations

import (
	"context"
	"fmt"
	"time"

	"github.com/joshsoftware/peerly-backend/internal/pkg/apperrors"
	"github.com/joshsoftware/peerly-backend/internal/pkg/constants"
	"github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	"github.com/joshsoftware/peerly-backend/internal/repository"
	logger "github.com/sirupsen/logrus"
)

type service struct {
	reportAppreciationRepo repository.ReportAppreciationStorer
	userRepo               repository.UserStorer
}

type Service interface {
	ReportAppreciation(ctx context.Context, reqData dto.ReportAppreciationReq) (resp dto.ReportAppricaitionResp, err error)
	ListReportedAppreciations(ctx context.Context) ([]dto.ReportedAppreciation, error)
	DeleteAppreciation(ctx context.Context, reqData dto.ModerationReq) (err error)
}

func NewService(reportAppreciationRepo repository.ReportAppreciationStorer, userRepo repository.UserStorer) Service {
	return &service{
		reportAppreciationRepo: reportAppreciationRepo,
		userRepo:               userRepo,
	}
}

func (rs *service) ReportAppreciation(ctx context.Context, reqData dto.ReportAppreciationReq) (resp dto.ReportAppricaitionResp, err error) {

	reporterId := ctx.Value(constants.UserId)
	fmt.Printf("reporterId: %T", reporterId)
	data, ok := reporterId.(int64)
	if !ok {
		logger.Error("Error in typecasting reporter id")
		err = apperrors.InternalServerError
		return
	}
	reqData.ReportedBy = data

	doesAppreciationExist, err := rs.reportAppreciationRepo.CheckAppreciation(ctx, reqData)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	if !doesAppreciationExist {
		err = apperrors.InvalidId
		return
	}

	isDupliate, err := rs.reportAppreciationRepo.CheckDuplicateReport(ctx, reqData)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	if isDupliate {
		err = apperrors.RepeatedReport
		return
	}

	usersData, err := rs.reportAppreciationRepo.GetSenderAndReceiver(ctx, reqData)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}
	if usersData.Sender == reqData.ReportedBy || usersData.Receiver == reqData.ReportedBy {
		err = apperrors.CannotReportOwnAppreciation
		return
	}

	resp, err = rs.reportAppreciationRepo.ReportAppreciation(ctx, reqData)
	if err != nil {
		err = apperrors.InternalServerError
		return
	}

	return
}

func (rs *service) ListReportedAppreciations(ctx context.Context) ([]dto.ReportedAppreciation, error) {

	var appreciationList []dto.ReportedAppreciation

	appreciations, err := rs.reportAppreciationRepo.ListReportedAppreciations(ctx)
	if err != nil {
		logger.Error(err.Error())
		err = apperrors.InternalServerError
		return appreciationList, err
	}

	for _, appreciation := range appreciations {

		senderDataReq := dto.GetUserByIdReq{
			UserId:          appreciation.Sender,
			QuaterTimeStamp: GetQuarterStartUnixTime(),
		}

		sender, err := rs.userRepo.GetUserById(ctx, senderDataReq)
		if err != nil {
			return appreciationList, err
		}

		receiverDataReq := dto.GetUserByIdReq{
			UserId:          appreciation.Receiver,
			QuaterTimeStamp: GetQuarterStartUnixTime(),
		}

		receiver, err := rs.userRepo.GetUserById(ctx, receiverDataReq)
		if err != nil {
			return appreciationList, err
		}

		svcApp := mapDbAppreciationsToSvcAppreciations(appreciation, sender, receiver)

		appreciationList = append(appreciationList, svcApp)

	}
	return appreciationList, err
}

func (rs *service) DeleteAppreciation(ctx context.Context, reqData dto.ModerationReq) (err error) {
	moderatorId := ctx.Value(constants.UserId)
	fmt.Printf("moderatorId: %T", moderatorId)
	data, ok := moderatorId.(int64)
	if !ok {
		logger.Error("Error in typecasting moderator id")
		err = apperrors.InternalServerError
		return
	}
	reqData.ModeratedBy = data
	doesExist, appreciationId, err := rs.reportAppreciationRepo.CheckResolution(ctx, reqData.ResolutionId)
	if err != nil {
		logger.Error(err.Error())
		err = apperrors.InternalServerError
		return
	}
	if !doesExist {
		logger.Errorf("No such resolution exists")
		err = apperrors.InvalidId
		return
	}
	reqData.AppreciationId = appreciationId
	err = rs.reportAppreciationRepo.DeleteAppreciation(ctx, reqData)
	if err != nil {
		logger.Error(err.Error())
		err = apperrors.InternalServerError
		return
	}
	return
}

func GetQuarterStartUnixTime() int64 {
	// Example function to get the Unix timestamp of the start of the quarter
	now := time.Now()
	quarterStart := time.Date(now.Year(), (now.Month()-1)/3*3+1, 1, 0, 0, 0, 0, time.UTC)
	return quarterStart.Unix() * 1000 // convert to milliseconds
}

func mapDbAppreciationsToSvcAppreciations(dbApp repository.ListReportedAppreciations, sender dto.GetUserByIdResp, receiver dto.GetUserByIdResp) (svcApp dto.ReportedAppreciation) {
	svcApp.Id = dbApp.Id
	svcApp.Appreciation_id = dbApp.Appreciation_id
	svcApp.AppreciationDesc = dbApp.AppreciationDesc
	svcApp.TotalRewardPoints = dbApp.TotalRewardPoints
	svcApp.Quarter = dbApp.TotalRewardPoints
	svcApp.CoreValueName = dbApp.CoreValueName
	svcApp.CoreValueDesc = dbApp.CoreValueDesc
	svcApp.SenderFirstName = sender.FirstName
	svcApp.SenderLastName = sender.LastName
	svcApp.SenderImgUrl = sender.ProfileImgUrl
	svcApp.SenderDesignation = sender.Designation
	svcApp.ReceiverFirstName = receiver.FirstName
	svcApp.ReceiverLastName = receiver.LastName
	svcApp.ReceiverImgUrl = receiver.ProfileImgUrl
	svcApp.ReceiverDesignation = receiver.Designation
	svcApp.ReportingComment = dbApp.ReportingComment
	svcApp.ReportedBy = dbApp.ReportedBy
	svcApp.ReportedAt = dbApp.ReportedAt
	return
}
