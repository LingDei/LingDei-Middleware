package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddCollectHandler 收藏视频
//
//	@Summary		收藏视频
//	@Description	收藏视频
//	@Tags			收藏管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/collect/add [post]
func AddCollectHandler(c *fiber.Ctx) error {
	// 获取参数
	var collect model.Collect
	c.QueryParser(&collect)
	collect.UUID = uuid.NewString()
	collect.User_UUID = method.GetUserFromToken(c).ID

	// 创建 Collect
	if err := method.AddCollect(collect); err != nil {
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

// GetCollectListHandler 获取某位用户自己的收藏列表
//
//	@Summary		获取某位用户自己的收藏列表
//	@Description	获取某位用户自己的收藏列表
//	@Tags			收藏管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.CollectListResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/collect/list [get]
func GetCollectListHandler(c *fiber.Ctx) error {
	// 获取参数
	var collect model.Collect
	c.QueryParser(&collect)
	collect.User_UUID = method.GetUserFromToken(c).ID

	// 获取收藏列表
	collectList, err := method.GetCollectListByUserUUID(collect.User_UUID)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.CollectListResp{
		Code:        200,
		CollectList: collectList,
	})
}

// DeleteCollectHandler 删除收藏
//
//	@Summary		删除收藏
//	@Description	删除收藏
//	@Tags			收藏管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"收藏UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/collect/delete [delete]
func DeleteCollectHandler(c *fiber.Ctx) error {
	// 获取参数
	var collect model.Collect
	c.QueryParser(&collect)
	collect.User_UUID = method.GetUserFromToken(c).ID

	// 删除Collect
	if err := method.DeleteCollect(collect.UUID); err != nil {
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

// CheckCollectExistHandler 检查是否收藏过
//
//	@Summary		检查是否收藏过
//	@Description	检查是否收藏过
//	@Tags			收藏管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.CollectStatusResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/collect/check [get]
func CheckCollectExistHandler(c *fiber.Ctx) error {
	// 获取参数
	var collect model.Collect
	c.QueryParser(&collect)
	collect.User_UUID = method.GetUserFromToken(c).ID

	// 检查是否收藏过
	if flag, err := method.CheckCollectExist(collect.Video_UUID, collect.User_UUID); flag {
		return c.JSON(model.CollectStatusResp{
			Code:   200,
			Status: true,
		})
	} else if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.CollectStatusResp{
		Code:   200,
		Status: false,
	})
}
