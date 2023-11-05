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
//	@Param			page		query		int		false	"页数"
//	@Param			page_size	query		int		false	"每页数量"
//	@Success		200			{object}	model.BarrageListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/barrage/list [get]
func GetBarrageListHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")

	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取 Barrage 列表
	barrages, err := method.GetBarrageList(video_uuid, page, page_size)
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

// GetRecentBarrageListHandler 获取最近的Barrage列表
//
//	@Summary		获取最近的Barrage列表
//	@Description	获取最近的Barrage列表
//	@Tags			弹幕管理
//	@Accept			json
//	@Produce		json
//	@Param			video_uuid	query		string	true	"视频UUID"
//	@Param			second		query		int64	true	"弹幕秒数"
//	@Param			limit		query		int64	false	"几秒内的"
//	@Success		200			{object}	model.BarrageListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/barrage/recent_list [get]
func GetRecentBarrageListHandler(c *fiber.Ctx) error {
	// 获取参数
	video_uuid := c.Query("video_uuid")
	second := c.QueryInt("second")
	limit := c.QueryInt("limit", 60)

	// 获取 Barrage 列表
	barrages, err := method.GetRecentBarrageList(video_uuid, int64(second), int64(limit))
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
//	@Success		200			{object}	model.BarrageCountResp
//	@Failure		400			{object}	model.OperationResp
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
//	@Success		200		{object}	model.BarrageResp
//	@Failure		400		{object}	model.OperationResp
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
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
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
