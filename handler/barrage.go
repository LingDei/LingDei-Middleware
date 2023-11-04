package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"
	"LingDei-Middleware/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddBarrageHandler 添加Barrage
//
//	@Summary		添加Barrage
//	@Description	添加Barrage
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Param			content		query		string	true	"弹幕内容"
//	@Param			second		query		int64	true	"弹幕秒数"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/barrage/add [post]
func AddBarrageHandler(c *fiber.Ctx) error {
	// 获取参数
	var barrage model.Barrage
	c.BodyParser(&barrage)
	barrage.UUID = uuid.NewString()
	barrage.User_UUID = method.GetUserFromToken(c).ID
	barrage.Timestamp = utils.GetTimestamp()

	// 创建 Barrage
	if err := method.AddBarrage(barrage); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "弹幕成功",
	})
}

// GetBarrageListHandler 获取Barrage列表
//
//	@Summary		获取Barrage列表
//	@Description	获取Barrage列表
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.BarrageListResp
//	@Router			/barrage/list [get]
func GetBarrageListHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")

	// 获取 Barrage 列表
	barrages, err := method.GetBarrageList(video_uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.BarrageListResp{
		Code:        200,
		BarrageList: barrages,
	})
}

// GetBarrageCountHandler 获取Barrage数量
//
//	@Summary		获取Barrage数量
//	@Description	获取Barrage数量
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.BarrageCountResp
//	@Router			/barrage/count [get]
func GetBarrageCountHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")

	// 获取 Barrage 数量
	count, err := method.GetBarrageCount(video_uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.BarrageCountResp{
		Code:  200,
		Count: count,
	})
}

// GetBarrageHandler 获取Barrage
//
//	@Summary		获取Barrage
//	@Description	获取Barrage
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"弹幕UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.BarrageResp
//	@Router			/barrage/get [get]
func GetBarrageHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")

	// 获取 Barrage
	barrage, err := method.GetBarrage(uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.BarrageResp{
		Code:    200,
		Barrage: barrage,
	})
}

// DeleteBarrageHandler 删除Barrage
//
//	@Summary		删除Barrage
//	@Description	删除Barrage
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"弹幕UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/barrage/delete [delete]
func DeleteBarrageHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")
	user_uuid := method.GetUserFromToken(c).ID

	if uuid == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "弹幕UUID不能为空",
		})
	}

	// 删除 Barrage
	if err := method.DeleteBarrage(uuid, user_uuid); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "弹幕删除成功",
	})
}
