package contract

import (
	"KeChuang-Middleware/config"
	"KeChuang-Middleware/method"
	"KeChuang-Middleware/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddContractAdminHandler godoc
//
//	@Summary		添加合同
//	@Description	添加合同
//	@Tags			合同管理（管理员）
//	@Accept			json
//	@Produce		json
//	@Param			name		query		string	true	"合同名称"
//	@Param			order_uuid	query		string	true	"关联订单UUID"
//	@Param			file		formData	file	true	"文件"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/admin/add [post]
func AddContractAdminHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 获取合同
	contract := model.Contract{}
	if err := c.QueryParser(&contract); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 随机生成合同UUID
	contract.UUID = uuid.NewString()

	// 获取file
	file_mul, err := c.FormFile("file")
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 打开文件
	file, err := file_mul.Open()
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	defer file.Close()

	if err := method.AddContract(contract, file); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "success",
	})
}

// GetContractListAdminHandler godoc
//
//	@Summary		获取合同列表
//	@Description	获取合同列表
//	@Tags			合同管理（管理员）
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.ContractListResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/admin/getList [get]
func GetContractListAdminHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	contracts, err := method.GetContractList()
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.ContractListResp{
		Code:      200,
		Contracts: contracts,
	})
}

// GetContractAdminHandler godoc
//
//	@Summary		获取合同及下载地址
//	@Description	获取合同及下载地址
//	@Tags			合同管理（管理员）
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"合同UUID"
//	@Success		200		{object}	model.ContractResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/admin/get [get]
func GetContractAdminHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 获取合同UUID
	uuid := c.Query("uuid")
	if uuid == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "合同UUID不能为空"})
	}

	contract, base64, err := method.GetContractWithBase64(uuid)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.ContractResp{
		Code:     200,
		URL:      "https://contract-lib.sanful.cn/" + uuid + ".pdf",
		Contract: contract,
		Base64:   base64,
	})
}

// DeleteContractAdminHandler godoc
//
//	@Summary		删除合同
//	@Description	删除合同
//	@Tags			合同管理（管理员）
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"合同UUID"
//	@Success		200		{object}	model.OperationResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/admin/delete [delete]
func DeleteContractAdminHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 获取合同UUID
	uuid := c.Query("uuid")
	if uuid == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "合同UUID不能为空"})
	}

	if err := method.DeleteContract(uuid); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "success",
	})
}

// UpdateContractAdminHandler godoc
//
//	@Summary		更新合同
//	@Description	更新合同
//	@Tags			合同管理（管理员）
//	@Accept			json
//	@Produce		json
//	@Param			uuid		query		string	true	"合同UUID"
//	@Param			name		query		string	false	"合同名称"
//	@Param			order_uuid	query		string	true	"关联订单UUID"
//	@Param			file		formData	file	false	"文件"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/admin/update [post]
func UpdateContractAdminHandler(c *fiber.Ctx) error {
	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	contract := model.Contract{}
	if err := c.QueryParser(&contract); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 检查FormFile是否存在

	// 获取file
	file_mul, err := c.FormFile("file")
	if err != nil && err.Error() != "request has no multipart/form-data Content-Type" {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	if err := method.UpdateContract(contract, file_mul); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.OperationResp{
		Code: 200,
		Msg:  "success",
	})
}
