package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["gocms/controllers:GxChannelController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxChannelController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxChannelController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxChannelController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxChannelController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxChannelController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxChannelController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxChannelController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxChannelController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxChannelController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxInfoController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxInfoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxInfoController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxInfoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxInfoController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxInfoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxInfoController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxInfoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxInfoController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxInfoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxLinkController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxLinkController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxLinkController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxLinkController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxLinkController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxLinkController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxLinkController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxLinkController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxLinkController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxLinkController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxMasterController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxMasterController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxMasterController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxMasterController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxMasterController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxMasterController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxMasterController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxMasterController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxMasterController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxMasterController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSlideController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSlideController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSlideController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSlideController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSlideController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSlideController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSlideController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSlideController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSlideController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSlideController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxSpecialController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxUserController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxUserController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxUserController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxUserController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gocms/controllers:GxUserController"] = append(beego.GlobalControllerRouter["gocms/controllers:GxUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
