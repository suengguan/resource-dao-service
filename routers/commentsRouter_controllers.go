package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"] = append(beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"] = append(beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "GetByUserId",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"] = append(beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"] = append(beego.GlobalControllerRouter["dao-service/resource-dao-service/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "DeleteById",
			Router: `/id/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
