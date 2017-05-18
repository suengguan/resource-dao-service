package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	_ "dao-service/resource-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	resource_base_url = "http://localhost:8080/v1/resource"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(localhost:3306)/PME?charset=utf8")
}

func Test_Resource_Create(t *testing.T) {
	// create user
	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM USER_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		t.Log("get user failed!", err)
		return
	}

	if num == 0 {
		// create user
		_, err := o.Raw("insert into USER_T(ID) values(1)").Exec()
		if err != nil {
			t.Log("insert user failed!", err)
			return
		}
		t.Log("create user success!")
	} else if num == 1 {
		// user is existed, nothing todo
		t.Log("user is already exited")
	} else {
		// error
		t.Log("get user failed!", err, num)
		return
	}

	// create resource
	var resource model.Resource
	var user model.User
	user.Id = 1
	resource.Id = 0
	resource.AlgorithmResource = "algorithm list"
	resource.CpuTotalResource = 10
	resource.CpuUsageResource = 5
	resource.CpuUnit = "core"
	resource.MemoryTotalResource = 100
	resource.MemoryUsageResource = 50
	resource.MemoryUnit = "Gi"
	resource.User = &user
	resource.QuotaNamespace = "user"
	resource.QuotaName = "user"

	// post create resource
	requestData, err := json.Marshal(&resource)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err := http.Post(resource_base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Resource_GetByUserId(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get resource by user id
	res, err = http.Get(resource_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Resource_Update(t *testing.T) {
	var resource model.Resource
	var user model.User
	user.Id = 1
	resource.Id = 0
	resource.AlgorithmResource = "algorithm list update"
	resource.CpuTotalResource = 100
	resource.CpuUsageResource = 50
	resource.CpuUnit = "core"
	resource.MemoryTotalResource = 1000
	resource.MemoryUsageResource = 500
	resource.MemoryUnit = "Gi"
	resource.User = &user
	resource.QuotaNamespace = "user-update"
	resource.QuotaName = "user-update"

	// put update project
	requestData, err := json.Marshal(&resource)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", resource_base_url, strings.NewReader(string(requestData)))

	res, err := client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Resource_DeleteById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get by user id
	res, err = http.Get(resource_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var resource model.Resource
	json.Unmarshal(([]byte)(response.Result), &resource)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(response.Result) <= 0 {
		t.Log("error : ", "there is no resource to delete!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", resource_base_url+"/id/"+strconv.FormatInt(resource.Id, 10), nil)

	res, err = client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
