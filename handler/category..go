package handler

import (
	"LingDei-Middleware/method"
	"LingDei-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddCategoryHandler 添加分类
//
//	@Summary		添加分类
//	@Description	添加分类
//	@Tags			分类管理
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"分类名称"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/category/add [post]
func AddCategoryHandler(c *fiber.Ctx) error {
	// 获取参数
	var category model.Category
	c.QueryParser(&category)
	category.UUID = uuid.NewString()

	// 创建 Category
	if err := method.AddCategory(category); err != nil {
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

// GetCategoryListHandler 获取分类列表
//
//	@Summary		获取分类列表
//	@Description	获取分类列表
//	@Tags			分类管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.CategoryListResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/category/list [get]
func GetCategoryListHandler(c *fiber.Ctx) error {
	// 获取分类列表
	categories, err := method.GetCategoryList()
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.CategoryListResp{
		Code:         200,
		CategoryList: categories,
	})
}

// GetCategoryHandler 获取分类
//
//	@Summary		获取分类
//	@Description	获取分类
//	@Tags			分类管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"分类 UUID"
//	@Success		200		{object}	model.CategoryResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/category/get [get]
func GetCategoryHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")

	// 获取分类
	category, err := method.GetCategoryByUUID(uuid)
	if err != nil {
		return c.JSON(model.OperationResp{
			Code: 400,
			Msg:  err.Error(),
		})
	}

	return c.JSON(model.CategoryResp{
		Code:     200,
		Category: category,
	})
}

// DeleteCategoryHandler 删除分类
//
//	@Summary		删除分类
//	@Description	删除分类
//	@Tags			分类管理
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"分类 UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/category/delete [delete]
func DeleteCategoryHandler(c *fiber.Ctx) error {
	// 获取参数
	uuid := c.Query("uuid")

	// 删除分类
	if err := method.DeleteCategory(uuid); err != nil {
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
