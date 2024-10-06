package controllers

import (
	"errors"
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OpponentRecruitingCommentController struct {
	OpponentRecruitingService *services.OpponentRecruitingService
	UserService               *services.UserService
}

// 対戦相手募集のコメントのコントローラーを作成する
func NewOpponentRecruitingCommentController(opponentRecruitingService *services.OpponentRecruitingService, userService *services.UserService) *OpponentRecruitingCommentController {
	return &OpponentRecruitingCommentController{
		OpponentRecruitingService: opponentRecruitingService,
		UserService:               userService,
	}
}

type OpponentRecruitingCommentCreateRequest struct {
	Content string `json:"content"`
}

func (oc *OpponentRecruitingCommentController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request OpponentRecruitingCommentCreateRequest

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		db := c.MustGet("tx").(*gorm.DB)
		// ユーザーを取得する
		user, err := oc.UserService.GetUser(c.MustGet("userId").(string))
		if err != nil || user == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				errors.New("ユーザーを取得できません").Error(),
				nil,
			))
			return
		}

		// 対戦相手募集のIDを取得
		opponentRecruitingId, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))
		// 対戦相手募集のコメントの構造体を作成
		opponentRecruitingComment := &models.OpponentRecruitingComment{
			OpponentRecruitingID: uint(opponentRecruitingId),
			UserID:               &user.ID,
			TeamID:               user.CurrentTeamId,
			Content:              request.Content,
		}

		// リクエストのバリデーションチェック
		if err := opponentRecruitingComment.ValidateCreate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// 対戦相手募集のコメントを作成する
		if err := oc.OpponentRecruitingService.CreateOpponentRecruitingComment(db, opponentRecruitingComment); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// 対戦相手募集のデータを取得する
		opponentRecruiting, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(opponentRecruitingId))
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"コメントを作成しました",
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruiting),
			},
		))
	}
}

type OpponentRecruitingCommentUpdateRequest struct {
	Content string `json:"content"`
}

func (oc *OpponentRecruitingCommentController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &OpponentRecruitingCommentUpdateRequest{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		// 対戦相手募集のコメントの構造体を作成
		commentId, _ := strconv.Atoi(c.Param("comment_id"))
		opponentRecruitingComment := &models.OpponentRecruitingComment{
			Content: request.Content,
		}
		opponentRecruitingComment.ID = uint(commentId)

		// リクエストのバリデーションチェック
		if err := opponentRecruitingComment.ValidateUpdate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		db := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// データを更新する
		if err := oc.OpponentRecruitingService.UpdateOpponentRecruitingComment(db, userId, opponentRecruitingComment); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// 対戦相手募集のIDを取得
		opponentRecruitingId, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))
		// データを取得する
		opponentRecruiting, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(opponentRecruitingId))
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"コメントを更新しました",
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruiting),
			},
		))
	}
}

func (oc *OpponentRecruitingCommentController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, _ := strconv.Atoi(c.Param("comment_id"))
		db := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// データを削除する
		if err := oc.OpponentRecruitingService.DeleteOpponentRecruitingComment(db, userId, uint(commentId)); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// 対戦相手募集のIDを取得
		opponentRecruitingId, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))
		// データを取得する
		opponentRecruiting, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(opponentRecruitingId))
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"コメントを削除しました",
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruiting),
			},
		))
	}
}
