package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddFollowHandler 关注用户
//
//	@Summary		关注用户
//	@Description	关注用户
//	@Tags			关注管理
//	@Accept			json
//	@Produce		json
//	@Param			follow_uuid	query		string	true	"被关注用户UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/follow/add [post]
func AddFollowHandler(c *fiber.Ctx) error {
	// 获取参数
	var follow model.Follow
	c.BodyParser(&follow)
	follow.UUID = uuid.NewString()
	follow.User_UUID = method.GetUserFromToken(c).ID

	// 创建 Follow
	if err := method.AddFollow(follow); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "关注成功",
	})
}

// DeleteFollowHandler 取消关注用户
//
//	@Summary		取消关注用户
//	@Description	取消关注用户
//	@Tags			关注管理
//	@Accept			json
//	@Produce		json
//	@Param			follow_uuid	query		string	true	"被关注用户UUID"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/follow/delete [delete]
func DeleteFollowHandler(c *fiber.Ctx) error {
	// 获取参数
	var follow model.Follow
	c.QueryParser(&follow)
	follow.User_UUID = method.GetUserFromToken(c).ID

	// 删除 Follow
	if err := method.DeleteFollow(follow.Follow_UUID, follow.User_UUID); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "取消关注成功",
	})
}

// GetFollowListHandler 获取我的关注列表
//
//	@Summary		获取我的关注列表
//	@Description	获取我的关注列表
//	@Tags			关注管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.FollowListResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/follow/list [get]
func GetFollowListHandler(c *fiber.Ctx) error {
	// 获取参数
	user_uuid := method.GetUserFromToken(c).ID

	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取关注列表
	followList, err := method.GetFollowListByUserUUID(user_uuid, page, page_size)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.FollowListResp{
		Code:       200,
		FollowList: followList,
	})
}

// GetFollowCountHandler 获取已关注up主的数量
//
//	@Summary		获取已关注up主的数量
//	@Description	获取已关注up主的数量
//	@Tags			关注管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.FollowCountResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/follow/count [get]
func GetFollowCountHandler(c *fiber.Ctx) error {
	// 获取参数
	user_uuid := method.GetUserFromToken(c).ID

	// 获取粉丝数量
	count, err := method.GetFollowCount(user_uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.FollowCountResp{
		Code:  200,
		Count: count,
	})
}

// CheckFollowStatusHandler 获取关注状态
//
//	@Summary		获取关注状态
//	@Description	获取关注状态
//	@Tags			关注管理
//	@Accept			json
//	@Produce		json
//	@Param			follow_uuid	query		string	true	"被关注用户UUID"
//	@Success		200			{object}	model.FollowStatusResp
//	@Failure		400			{object}	model.FollowStatusResp
//	@Security		ApiKeyAuth
//	@Router			/follow/check [get]
func CheckFollowStatusHandler(c *fiber.Ctx) error {
	// 获取参数
	var follow model.Follow
	c.QueryParser(&follow)
	follow.User_UUID = method.GetUserFromToken(c).ID

	// 获取关注状态
	status, err := method.GetFollowStatus(follow.Follow_UUID, follow.User_UUID)
	if err != nil {
		return c.JSON(model.FollowStatusResp{
			Code:   400,
			Status: false,
		})
	}

	return c.JSON(model.FollowStatusResp{
		Code:   200,
		Status: status,
	})
}

// GetFansHandler 获取我的粉丝列表
//
//	@Summary		获取我的粉丝列表
//	@Description	获取我的粉丝列表
//	@Tags			粉丝管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.FollowListResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/fan/list [get]
func GetFansHandler(c *fiber.Ctx) error {
	// 获取参数
	user_uuid := method.GetUserFromToken(c).ID

	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取粉丝列表
	fansList, err := method.GetFanListByFollowUUID(user_uuid, page, page_size)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.FanListResp{
		Code:    200,
		FanList: fansList,
	})
}

// GetFanCountHandler 获取粉丝数量
//
//	@Summary		获取粉丝数量
//	@Description	获取粉丝数量
//	@Tags			粉丝管理
//	@Accept			json
//	@Produce		json
//	@Param			follow_uuid	query		string	false	"用户UUID"
//	@Success		200			{object}	model.FollowCountResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/fan/count [get]
func GetFanCountHandler(c *fiber.Ctx) error {
	// 获取参数
	follow_uuid := c.Query("follow_uuid")
	if follow_uuid == "" {
		follow_uuid = method.GetUserFromToken(c).ID
	}

	// 获取粉丝数量
	count, err := method.GetFanCount(follow_uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.FollowCountResp{
		Code:  200,
		Count: count,
	})
}
