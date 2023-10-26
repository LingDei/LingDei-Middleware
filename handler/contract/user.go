package contract

import (
	"KeChuang-Middleware/method"
	"KeChuang-Middleware/model"

	"github.com/gofiber/fiber/v2"
)

// GetContractHandler godoc
//
//	@Summary		获取合同及下载地址
//	@Description	获取合同及下载地址
//	@Tags			合同管理（学生/教师）
//	@Accept			json
//	@Produce		json
//	@Param			uuid	query		string	true	"合同UUID"
//	@Success		200		{object}	model.ContractResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/contract/get [get]
func GetContractHandler(c *fiber.Ctx) error {
	// 获取合同UUID
	uuid := c.Query("uuid")
	if uuid == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "合同UUID不能为空"})
	}

	// 查询用户对某个合同的权限
	if err := method.CheckContractUserUUID(method.GetUserFromToken(c).ID, uuid); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
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
