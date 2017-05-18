package controllers

import (
	"dao-service/resource-dao-service/models"
	"dao-service/resource-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Resource
type ResourceController struct {
	beego.Controller
}

// @Title Create
// @Description create resource
// @Param	body		body 	models.Resource	true		"body for resource content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *ResourceController) Create() {
	var err error
	var resource model.Resource
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &resource)
	if err == nil {
		var svc service.ResourceService
		var result []byte
		err = svc.Create(&resource)
		if err == nil {
			result, err = json.Marshal(&resource)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetByUserId
// @Description get user's resource
// @Param	userid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:userId [get]
func (this *ResourceController) GetByUserId() {
	var err error
	var response models.Response

	var userId int64
	userId, err = this.GetInt64(":userId")
	beego.Debug("GetByUserId", userId)
	if userId > 0 && err == nil {
		var svc service.ResourceService
		var resource *model.Resource
		var result []byte
		resource, err = svc.GetByUserId(userId)
		if err == nil {
			result, err = json.Marshal(resource)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title Update
// @Description update the resource
// @Param	body		body 	models.Resource	true		"body for resource content"
// @Success 200 {object} models.Response
// @router / [put]
func (this *ResourceController) Update() {
	var err error
	var resource model.Resource
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &resource)
	if err == nil {
		var svc service.ResourceService
		var result []byte
		err = svc.Update(&resource)
		if err == nil {
			result, err = json.Marshal(&resource)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title DeleteById
// @Description delete the resource by id
// @Param	id		path 	int64	true		"The int you want to delete"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router /id/:id [delete]
func (this *ResourceController) DeleteById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("DeleteById", id)
	if id > 0 && err == nil {
		var svc service.ResourceService
		err = svc.DeleteById(id)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "resource id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
