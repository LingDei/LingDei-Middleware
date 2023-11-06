package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddLikeHandler 给视频点赞
//
//	@Summary		给视频点赞
//	@Description	给视频点赞
//	@Tags			点赞管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/like/add [post]
func AddLikeHandler(c *fiber.Ctx) error {
	// 获取参数
	var like model.Like
	c.BodyParser(&like)
	like.UUID = uuid.NewString()
	like.User_UUID = method.GetUserFromToken(c).ID

	// 创建 Like
	if err := method.AddLike(like); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "添加成功",
	})
}

// GetLikeListHandler 获取某位用户自己的点赞列表
//
//	@Summary		获取某位用户自己的点赞列表
//	@Description	获取某位用户自己的点赞列表
//	@Tags			点赞管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.LikeListResp
//	@Failure		400			{object}	model.LikeListResp
//	@Security		ApiKeyAuth
//	@Router			/like/list [get]
func GetLikeListHandler(c *fiber.Ctx) error {
	// 获取参数
	var like model.Like
	like.User_UUID = method.GetUserFromToken(c).ID

	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取Like列表
	likes, err := method.GetLikeListByUserUUID(like.User_UUID, page, page_size)
	if err != nil {
		return c.JSON(model.LikeListResp{
			Code:     400,
			LikeList: nil,
		})
	}

	// 获取Like数量
	count, err := method.GetLikeCountByUserUUID(like.User_UUID)
	if err != nil {
		return c.JSON(model.LikeListResp{
			Code:     400,
			LikeList: nil,
		})
	}

	return c.JSON(model.LikeListResp{
		Code:     200,
		LikeList: likes,
		Total:    count,
	})
}

// GetLikeCountHandler 获取某个视频的点赞数量
//
//	@Summary		获取某个视频的点赞数量
//	@Description	获取某个视频的点赞数量
//	@Tags			点赞管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.LikeCountResp
//	@Failure		400			{object}	model.LikeCountResp
//	@Router			/like/count [get]
func GetLikeCountHandler(c *fiber.Ctx) error {
	// 获取参数
	var like model.Like
	c.QueryParser(&like)

	// 获取Like数量
	count, err := method.GetLikeCount(like.Video_UUID)
	if err != nil {
		return c.JSON(model.LikeCountResp{
			Code:  400,
			Count: 0,
		})
	}

	return c.JSON(model.LikeCountResp{
		Code:  200,
		Count: count,
	})
}

// DeleteLikeHandler 取消点赞
//
//	@Summary		取消点赞
//	@Description	取消点赞
//	@Tags			点赞管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/like/delete [delete]
func DeleteLikeHandler(c *fiber.Ctx) error {
	// 获取参数
	var like model.Like
	c.QueryParser(&like)
	like.User_UUID = method.GetUserFromToken(c).ID

	fmt.Println(like)

	// 删除Like
	if err := method.DeleteLike(like.Video_UUID, like.User_UUID); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "删除成功",
	})
}

// CheckLikeExistHandler 检查是否点赞过
//
//	@Summary		检查是否点赞过
//	@Description	检查是否点赞过
//	@Tags			点赞管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.LikeStatusResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/like/check [get]
func CheckLikeExistHandler(c *fiber.Ctx) error {
	// 获取参数
	var like model.Like
	c.QueryParser(&like)
	like.User_UUID = method.GetUserFromToken(c).ID

	// 检查是否点赞过
	if flag, err := method.CheckLikeExist(like.Video_UUID, like.User_UUID); flag {
		return c.JSON(model.LikeStatusResp{
			Code:   200,
			Status: true,
		})
	} else if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.LikeStatusResp{
		Code:   200,
		Status: false,
	})
}
