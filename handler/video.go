package handler

import (
	"LingDei-Middleware/config"
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddDisciplineHandler 添加视频
//
//	@Summary		添加视频
//	@Description	添加视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			name			query		string	true	"视频名称"
//	@Param			category_uuid	query		string	true	"视频分类 UUID"
//	@Param			file			formData	file	true	"视频文件"
//	@Success		200				{object}	model.OperationResp
//	@Failure		400				{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/add [post]
func AddVideoHandler(c *fiber.Ctx) error {
	// 获取参数
	var video model.Video
	c.QueryParser(&video)
	video.UUID = uuid.NewString()
	video.Author_UUID = method.GetUserFromToken(c).ID

	// 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	// 检查文件类型
	if file.Header["Content-Type"][0] != "video/mp4" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "请上传 mp4 格式的视频文件",
		})
	}

	video.URL = "https://bucket.lingdei.doyi.online/videos/" + video.UUID + ".mp4"
	video.Thumbnail_URL = "https://bucket.lingdei.doyi.online/frames/" + video.UUID + ".jpg"
	video.Views = 0

	// 创建 Video
	if err := method.AddVideo(video, file); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "ok",
	})
}

// GetVideoListHandler 获取视频列表
//
//	@Summary		获取视频列表
//	@Description	获取视频列表
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.VideoListResp
//	@Failure		400	{object}	model.OperationResp
//	@Router			/video/list [get]
func GetVideoListHandler(c *fiber.Ctx) error {
	// 获取视频列表
	videos, err := method.GetVideoList()
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.VideoListResp{
		Code:       200,
		Video_List: videos,
	})
}

// GetVideoHandler 获取视频
//
//	@Summary		获取视频
//	@Description	获取视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"视频 UUID"
//	@Success		200		{object}	model.VideoResp
//	@Failure		400		{object}	model.OperationResp
//	@Router			/video/get [get]
func GetVideoHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")

	// 获取视频
	video, err := method.GetVideo(uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.VideoResp{
		Code:  200,
		Video: video,
	})
}

// DeleteVideoHandler 删除视频
//
//	@Summary		删除视频 🦸
//	@Description	删除视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"视频 UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/delete [delete]
func DeleteVideoHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 获取参数
	uuid := c.Query("uuid")

	// 删除视频
	if err := method.DeleteVideo(uuid); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "ok",
	})
}

// UpdateVideoHandler 更新视频
//
//	@Summary		更新视频 🦸
//	@Description	更新视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid			query		string	true	"视频 UUID"
//	@Param			name			query		string	false	"视频名称"
//	@Param			category_uuid	query		string	false	"视频分类 UUID"
//	@Success		200				{object}	model.OperationResp
//	@Failure		400				{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/update [post]
func UpdateVideoHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 获取参数
	var video model.Video
	c.QueryParser(&video)

	// 更新视频
	if err := method.UpdateVideo(video); err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "ok",
	})
}
