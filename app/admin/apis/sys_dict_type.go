package apis

import (
	"go-admin/app/admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysDictType struct {
	api.Api
}

// GetSysDictTypeList 字典类型列表数据
// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [get]
// @Security Bearer
func (e SysDictType) GetSysDictTypeList(c *gin.Context) {
	s := service.SysDictType{}
	d := &dto.SysDictTypeSearch{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysDictType, 0)
	var count int64

	err = s.GetPage(d, &list, &count)
	if err != nil {
		e.Error(http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetSysDictType 字典类型通过字典id获取
// @Summary 字典类型通过字典id获取
// @Description 获取JSON
// @Tags 字典类型
// @Param dictId path int true "字典类型编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [get]
// @Security Bearer
func (e SysDictType) GetSysDictType(c *gin.Context) {
	s := service.SysDictType{}
	d := &dto.SysDictTypeById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}
	var object models.SysDictType
	err = s.Get(d, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(object, "查看成功")
}

//InsertSysDictType 字典类型创建
// @Summary 添加字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeControl true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/type [post]
// @Security Bearer
func (e SysDictType) InsertSysDictType(c *gin.Context) {
	s := service.SysDictType{}
	d := &dto.SysDictTypeControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	d.SetCreateBy(user.GetUserId(c))

	err = s.Insert(d)
	if err != nil {
		e.Logger.Error(err)
		e.Error(http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(d.GetId(), "创建成功")
}

// UpdateSysDictType
// @Summary 修改字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeControl true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/type/{dictId} [put]
// @Security Bearer
func (e SysDictType) UpdateSysDictType(c *gin.Context) {
	s := service.SysDictType{}
	d := &dto.SysDictTypeControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	d.SetUpdateBy(user.GetUserId(c))

	err = s.Update(d)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(d.GetId(), "更新成功")
}

// DeleteSysDictType
// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictId path int true "dictId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/type/{dictId} [delete]
func (e SysDictType) DeleteSysDictType(c *gin.Context) {
	s := service.SysDictType{}
	d := new(dto.SysDictTypeById)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}
	// 设置编辑人
	d.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(d)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(d.GetId(), "删除成功")
}

// GetSysDictTypeAll
// @Summary 字典类型全部数据 代码生成使用接口
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type-option-select [get]
// @Security Bearer
func (e SysDictType) GetSysDictTypeAll(c *gin.Context) {
	s := service.SysDictType{}
	d := &dto.SysDictTypeSearch{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysDictType, 0)

	err = s.GetAll(d, &list)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}