package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

var cfg = beego.AppConfig

type ResourceDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Resource
}

func NewResourceDao() *ResourceDao {
	d := new(ResourceDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *ResourceDao) Create(resource *model.Resource) error {
	num, err := this.m_Orm.Insert(resource)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *ResourceDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no resource to delete")
		return err
	}

	return err
}

// update
func (this *ResourceDao) Update(resource *model.Resource) error {
	num, err := this.m_Orm.Update(resource)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no resource to update")
	}

	return err
}

// find
func (this *ResourceDao) GetByUserId(userId int64) (*model.Resource, error) {
	var resource model.Resource

	err := this.m_QuerySeter.Filter("USER_ID", userId).One(&resource)

	if err != nil {
		return nil, err
	}

	return &resource, err
}

func (this *ResourceDao) GetById(Id int64) (*model.Resource, error) {
	var resource model.Resource

	err := this.m_QuerySeter.Filter("ID", Id).One(&resource)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &resource, err
}
