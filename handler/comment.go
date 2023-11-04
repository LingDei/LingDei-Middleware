package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"
	"LingDei-Middleware/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddCommentHandler 添加Comment
//
//	@Summary		添加Comment
//	@Description	添加Comment
//	@Tags			评论管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Param			content		query		string	true	"评论内容"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/comment/add [post]
func AddCommentHandler(c *fiber.Ctx) error {
	// 获取参数
	var comment model.Comment
	c.BodyParser(&comment)
	comment.UUID = uuid.NewString()
	comment.User_UUID = method.GetUserFromToken(c).ID
	comment.Timestamp = utils.GetTimestamp()

	// 创建 Comment
	if err := method.AddComment(comment); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "评论成功",
	})
}

// GetCommentListHandler 获取Comment列表
//
//	@Summary		获取Comment列表
//	@Description	获取Comment列表
//	@Tags			评论管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	CommentListResp
//	@Failure		400			{object}	CommentListResp
//	@Router			/comment/list [get]
func GetCommentListHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")

	// 获取 Comment 列表
	comments, err := method.GetCommentList(video_uuid)
	if err != nil {
		return c.JSON(model.CommentListResp{
			Code:        400,
			CommentList: nil,
		})
	}

	return c.JSON(model.CommentListResp{
		Code:        200,
		CommentList: comments,
	})
}

// GetCommentCountHandler 获取Comment数量
//
//	@Summary		获取Comment数量
//	@Description	获取Comment数量
//	@Tags			评论管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	CommentCountResp
//	@Failure		400			{object}	CommentCountResp
//	@Router			/comment/count [get]
func GetCommentCountHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")

	// 获取 Comment 数量
	count, err := method.GetCommentCount(video_uuid)
	if err != nil {
		return c.JSON(model.CommentCountResp{
			Code:  400,
			Count: 0,
		})
	}

	return c.JSON(model.CommentCountResp{
		Code:  200,
		Count: count,
	})
}

// DeleteCommentHandler 删除Comment
//
//	@Summary		删除Comment
//	@Description	删除Comment
//	@Tags			评论管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"评论UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/comment/delete [delete]
func DeleteCommentHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")
	user_uuid := method.GetUserFromToken(c).ID

	if uuid == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "评论UUID不能为空",
		})
	}

	// 删除 Comment
	if err := method.DeleteComment(uuid, user_uuid); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "删除评论成功",
	})
}
