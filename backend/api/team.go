package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/token"
	"github.com/sqlc-dev/pqtype"
)

type createTeamRequest struct {
	Name         string                `json:"name" binding:"required"`
	Currency     string                `json:"currency" binding:"required"`
	Team_members pqtype.NullRawMessage `json:"team_members" binding:"required"`
}

func (server *Server) createTeam(ctx *gin.Context) {
	var req createTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTeamParams{
		TeamName:    req.Name,
		Currency:    req.Currency,
		TeamMembers: pqtype.NullRawMessage{RawMessage: json.RawMessage(authPayload.ID.String()), Valid: true},
	}

	team, err := server.store.CreateTeam(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, team)

}

// type listTeamRequest struct {
// 	PageID int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func(server *Server) listTeams(ctx *gin.Context){
// 		var req listTeamRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest,errorResponse(err))
// 		return
// 	}

// 	arg := db.ListUsersParams{
// 		Limit: req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}

// 	users, err := server.store.ListUsers(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
// 		return
// 	}

// 	var userResponses []userResponse
// for _, user := range users {
//     userResponse := newUserResponse(user)
//     userResponses = append(userResponses, userResponse)
// }

// 	ctx.JSON(http.StatusOK, userResponses)
// }
