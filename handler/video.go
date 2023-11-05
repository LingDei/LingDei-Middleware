package handler

import (
	"LingDei-Middleware/config"
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"
	"LingDei-Middleware/utils"

	"github.com/gofiber/fiber/v2"
)

// AddDisciplineHandler 添加视频
//
//	@Summary		添加视频
//	@Description	添加视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid			query		string	true	"视频 UUID"
//	@Param			name			query		string	true	"视频名称"
//	@Param			category_uuid	query		string	true	"视频分类 UUID"
//	@Success		200				{object}	model.OperationResp
//	@Failure		400				{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/add [post]
func AddVideoHandler(c *fiber.Ctx) error {
	// 获取参数
	var video model.Video
	c.BodyParser(&video)
	video.Author_UUID = method.GetUserFromToken(c).ID

	video.URL = "https://bucket.lingdei.doyi.online/videos/" + video.UUID + ".mp4"
	video.Thumbnail_URL = "https://bucket.lingdei.doyi.online/frames/" + video.UUID + ".jpg"
	video.Timestamp = utils.GetTimestamp()
	video.Views = 0

	// 创建 Video
	if err := method.AddVideo(video); err != nil {
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

// GetUploadTokenHandler 获取上传凭证
//
//	@Summary		获取上传凭证
//	@Description	获取上传凭证
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UploadTokenResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/upload_token [get]
func GetUploadTokenHandler(c *fiber.Ctx) error {
	// 获取上传凭证
	video_uuid, upload_token, err := utils.GetUploadToken()
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.UploadTokenResp{
		Code:         200,
		Video_UUID:   video_uuid,
		Upload_Token: upload_token,
	})
}

// GetVideoListHandler 获取视频列表
//
//	@Summary		获取视频列表
//	@Description	获取视频列表
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			user_uuid		query		string	false	"用户 UUID"
//	@Param			category_uuid	query		string	false	"视频分类 UUID"
//	@Param			page			query		int		false	"页数"
//	@Param			page_size		query		int		false	"每页数量"
//	@Success		200				{object}	model.VideoListResp
//	@Failure		400				{object}	model.OperationResp
//	@Router			/video/list [get]
func GetVideoListHandler(c *fiber.Ctx) error {
	user_uuid := c.Query("user_uuid")
	category_uuid := c.Query("category_uuid")

	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取视频列表
	videos, err := method.GetVideoList(category_uuid, user_uuid, page, page_size)
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

// GetRecommendVideoList 获取推荐的视频列表
//
//	@Summary		获取推荐的视频列表
//	@Description	获取推荐的视频列表
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.VideoListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/video/recommend_list [get]
func GetRecommendVideoListHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取视频列表
	videos, err := method.GetRecommendVideoList(page, page_size)
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

// GetMyVideoListHandler 获取我的视频列表
//
//	@Summary		获取我的视频列表
//	@Description	获取我的视频列表
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.VideoListResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/my_list [get]
func GetMyVideoListHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取视频列表
	videos, err := method.GetVideoList("", method.GetUserFromToken(c).ID, page, page_size)
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
	c.BodyParser(&video)

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

// AddVideoViewsCountHandler 添加视频播放量
//
//	@Summary		添加视频播放量
//	@Description	添加视频播放量
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"视频 UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Router			/video/views/add [post]
func AddVideoViewsCountHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.FormValue("uuid")

	if uuid == "" {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  "视频 UUID 不能为空",
		})
	}

	// 添加视频播放量
	if err := method.AddVideoViewsCount(uuid); err != nil {
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

// GetMyFollowVideoListHandler 获取我关注的用户的视频
//
//	@Summary		获取我关注的用户的视频
//	@Description	获取我关注的用户的视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"页数"
//	@Param			page_size	query		int	false	"每页数量"
//	@Success		200			{object}	model.VideoListResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/video/follow_list [get]
func GetMyFollowVideoListHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 获取视频列表
	videos, err := method.GetFollowVideos(method.GetUserFromToken(c).ID, page, page_size)
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

// SearchVideoHandler   搜索视频
//
//	@Summary		搜索视频
//	@Description	搜索视频
//	@Tags			视频管理
//	@Accept			json
//	@Produce		json
//	@Param			keyword		query		string	fasle	"关键词"
//	@Param			page		query		int		false	"页数"
//	@Param			page_size	query		int		false	"每页数量"
//	@Success		200			{object}	model.VideoListResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/video/search [get]
func SearchVideoHandler(c *fiber.Ctx) error {
	// 获取参数
	keyword := c.Query("keyword")
	page := c.QueryInt("page", 1)
	page_size := c.QueryInt("page_size", 9)

	// 搜索视频
	videos, err := method.SearchVideo(keyword, page, page_size)
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
