package service

import (
	"fmt"
	"model"

	"dao-service/resource-dao-service/dao"

	"github.com/astaxie/beego"
)

type ResourceService struct {
}

func (this *ResourceService) Create(resource *model.Resource) error {
	var err error
	var resourceDao = dao.NewResourceDao()

	err = resourceDao.Create(resource)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create resource failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ResourceService) GetByUserId(userId int64) (*model.Resource, error) {
	var err error
	var resourceDao = dao.NewResourceDao()
	var resource *model.Resource

	// get resource
	resource, err = resourceDao.GetByUserId(userId)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "resource is not existed!")
		return nil, err
	}

	return resource, err
}

func (this *ResourceService) Update(resource *model.Resource) error {
	var err error
	var resourceDao = dao.NewResourceDao()

	err = resourceDao.Update(resource)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update resource failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ResourceService) DeleteById(id int64) error {
	beego.Debug("delete by id")
	var err error
	var resourceDao = dao.NewResourceDao()

	// delete resource
	err = resourceDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete resource failed!", "reason:"+err.Error())
		return err
	}

	return err
}
