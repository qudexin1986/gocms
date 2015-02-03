// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gocms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/gx_channel",
			beego.NSInclude(
				&controllers.GxChannelController{},
			),
		),

		beego.NSNamespace("/gx_info",
			beego.NSInclude(
				&controllers.GxInfoController{},
			),
		),

		beego.NSNamespace("/gx_link",
			beego.NSInclude(
				&controllers.GxLinkController{},
			),
		),

		beego.NSNamespace("/gx_master",
			beego.NSInclude(
				&controllers.GxMasterController{},
			),
		),

		beego.NSNamespace("/gx_slide",
			beego.NSInclude(
				&controllers.GxSlideController{},
			),
		),

		beego.NSNamespace("/gx_special",
			beego.NSInclude(
				&controllers.GxSpecialController{},
			),
		),

		beego.NSNamespace("/gx_user",
			beego.NSInclude(
				&controllers.GxUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
