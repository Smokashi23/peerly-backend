package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshsoftware/peerly-backend/internal/app/appreciation"
	"github.com/joshsoftware/peerly-backend/internal/pkg/apperrors"
	"github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	logger "github.com/sirupsen/logrus"
)

func createAppreciationHandler(appreciationSvc appreciation.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var appreciation dto.Appreciation
		err := json.NewDecoder(req.Body).Decode(&appreciation)
		if err != nil {
			logger.Error(fmt.Sprintf("Error while decoding request data : %v", err))
			err = apperrors.JSONParsingErrorReq
			dto.ErrorRepsonse(rw, err)
			return
		}

		err = appreciation.ValidateCreateAppreciation()
		if err != nil {
			logger.Error(fmt.Sprintf("Error while validating request data : %v", err))
			dto.ErrorRepsonse(rw, err)
			return
		}

		resp, err := appreciationSvc.CreateAppreciation(req.Context(), appreciation)
		if err != nil {
			logger.Error(fmt.Sprintf("err : %v", err))
			dto.ErrorRepsonse(rw, err)
			return
		}
		dto.SuccessRepsonse(rw, http.StatusCreated, "Appreciation created successfully", resp)
	})
}

func getAppreciationByIDHandler(appreciationSvc appreciation.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		apprID, err := strconv.Atoi(vars["id"])
		if err != nil {
			dto.ErrorRepsonse(rw, err)
			return
		}

		resp, err := appreciationSvc.GetAppreciationById(req.Context(), int32(apprID))
		if err != nil {
			logger.Error(fmt.Sprintf("err : %v", err))
			dto.ErrorRepsonse(rw, err)
			return
		}
		dto.SuccessRepsonse(rw, http.StatusOK, "Appreciation data got successfully", resp)
	})
}

// getAppreciationsHandler handles HTTP requests for appreciations
func getAppreciationsHandler(appreciationSvc appreciation.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var filter dto.AppreciationFilter

		// Extract query parameters or body fields
		filter.Name = req.URL.Query().Get("name")
		filter.SortOrder = req.URL.Query().Get("sort_order")

		// Get pagination parameters
		page, limit := getPaginationParams(req)

		filter.Limit = limit
		filter.Page = page
		filter.Self = getSelfParam(req)
		// Call your appreciationService to fetch appreciations based on filter
		appreciations, err := appreciationSvc.GetAppreciations(req.Context(), filter)
		if err != nil {
			logger.Error(fmt.Sprintf("err : %v", err))
			dto.ErrorRepsonse(rw, err)
			return
		}
		dto.SuccessRepsonse(rw, http.StatusOK, "Appreciations data got successfully ", appreciations)
	})
}

func deleteAppreciationHandler(appreciationSvc appreciation.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		apprId, err := strconv.Atoi(vars["id"])
		if err != nil {
			logger.Error(fmt.Sprintf("Error while decoding request param data : %v", err))
			dto.ErrorRepsonse(rw, apperrors.BadRequest)
			return
		}

		res, err := appreciationSvc.DeleteAppreciation(req.Context(), int32(apprId))
		if err != nil {
			logger.Error(fmt.Sprintf("err : %v", err))
			dto.ErrorRepsonse(rw, err)
			return
		}
		if !res {
			dto.ErrorRepsonse(rw, apperrors.InternalServer)
			return
		}
		dto.SuccessRepsonse(rw, http.StatusOK, "Appreciation invalidate successfully", nil)
	})
}
