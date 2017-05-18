// @APIVersion 1.0.0
// @Title resource-dao-service API
// @Description resource-dao-service only serve the RESOURCE_T
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dao-service/resource-dao-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/resource",
			beego.NSInclude(
				&controllers.ResourceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
