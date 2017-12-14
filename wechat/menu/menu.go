package menu

import (
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/oauth2"
	"log"
	"util/wechat/core"
)

const (
	// 下面6个类型(包括view类型)的按钮是在公众平台官网发布的菜单按钮类型
	ButtonTypeText  = "text"
	ButtonTypeImage = "img"
	ButtonTypePhoto = "photo"
	ButtonTypeVideo = "video"
	ButtonTypeVoice = "voice"

	// 上面5个类型的按钮不能通过API设置

	ButtonTypeView  = "view"  // 跳转URL
	ButtonTypeClick = "click" // 点击推事件

	// 下面的按钮类型仅支持微信 iPhone5.4.1 以上版本, 和 Android5.4 以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	ButtonTypeScanCodePush    = "scancode_push"      // 扫码推事件
	ButtonTypeScanCodeWaitMsg = "scancode_waitmsg"   // 扫码带提示
	ButtonTypePicSysPhoto     = "pic_sysphoto"       // 系统拍照发图
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album" // 拍照或者相册发图
	ButtonTypePicWeixin       = "pic_weixin"         // 微信相册发图
	ButtonTypeLocationSelect  = "location_select"    // 发送位置

	// 下面的按钮类型专门给第三方平台旗下未微信认证(具体而言, 是资质认证未通过)的订阅号准备的事件类型,
	// 它们是没有事件推送的, 能力相对受限, 其他类型的公众号不必使用.
	ButtonTypeMediaId     = "media_id"     // 下发消息
	ButtonTypeViewLimited = "view_limited" // 跳转图文消息URL

	temp = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx6ddf008341937de1&redirect_uri=http%3A%2F%2Fjyk.sunzhaoning.com&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect"
)

type Menu struct {
	Buttons   []Button   `json:"button,omitempty"`
	MatchRule *MatchRule `json:"matchrule,omitempty"`
	MenuId    int64      `json:"menuid,omitempty"` // 有个性化菜单时查询接口返回值包含这个字段
}

type MatchRule struct {
	GroupId            *int64 `json:"group_id,omitempty"`
	Sex                *int   `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType *int   `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

type Button struct {
	Type       string   `json:"type,omitempty"`       // 非必须; 菜单的响应动作类型
	Name       string   `json:"name,omitempty"`       // 必须;  菜单标题
	Key        string   `json:"key,omitempty"`        // 非必须; 菜单KEY值, 用于消息接口推送
	URL        string   `json:"url,omitempty"`        // 非必须; 网页链接, 用户点击菜单可打开链接
	MediaId    string   `json:"media_id,omitempty"`   // 非必须; 调用新增永久素材接口返回的合法media_id
	SubButtons []Button `json:"sub_button,omitempty"` // 非必须; 二级菜单数组
}

func InitMenu(appid, secret string) error {
	log.Println(appid, secret)

	mine := oauth2.AuthCodeURL(appid,
		"http://jyk.sunzhaoning.com",
		"snsapi_userinfo",
		"mine")
	transfercard := oauth2.AuthCodeURL(appid,
		"http://jyk.sunzhaoning.com/attorncard",
		"snsapi_userinfo",
		"transfercard")
	quickPay := oauth2.AuthCodeURL(appid,
		"http://jyk.sunzhaoning.com/quick_payment",
		"snsapi_userinfo",
		"quick_pay")
	selfPay := "http://jyk.sunzhaoning.com/oilcard/selfhelpcharge"

	client := core.NewClient(appid, secret)
	menu.Delete(client.Clt)
	if me, _, err := menu.Get(client.Clt); err == nil && len(me.Buttons) > 0 {
		log.Println("menus:=", me)
		//menu.Delete(client.Clt)
		return err
	} else {
		log.Println("menu: ", me)
		log.Println("开始初始化菜单...")
	}
	allmenu := new(menu.Menu)
	allmenu.MenuId = 1
	//一级菜单： 充值油卡
	button := new(menu.Button)
	button.Key = "pay"
	button.MediaId = "121"
	button.Name = "充油卡"
	button.URL = "http://www.baidu.com"
	button.Type = "view"
	// 二级菜单： 支付菜单  // 自助充值页面URL： http://m.sinopecsales.com/websso/loginAction_form.action
	li_button := new(menu.Button)
	li_button.Key = "quick_pay"
	li_button.Name = "极速充值"
	li_button.URL = quickPay
	li_button.Type = "view"
	button.SubButtons = append(button.SubButtons, *li_button)
	li_button.Key = "self_pay"
	li_button.Name = "自助充值"
	li_button.URL = selfPay
	li_button.Type = "view"
	button.SubButtons = append(button.SubButtons, *li_button)
	// 一级菜单： 转让油卡
	allmenu.Buttons = append(allmenu.Buttons, *button)
	button = new(menu.Button)
	button.Key = "transfercard"
	button.MediaId = "122"
	button.Name = "转让油卡"
	button.URL = transfercard
	button.Type = "view"
	allmenu.Buttons = append(allmenu.Buttons, *button)
	//一级菜单： 我
	button = new(menu.Button)
	button.Key = "mine"
	button.MediaId = "123"
	button.Name = "我"
	button.URL = mine
	button.Type = "view"
	allmenu.Buttons = append(allmenu.Buttons, *button)

	err := menu.Create(client.Clt, allmenu)
	if err != nil {
		log.Println("create_error:", err)
	} else {
		log.Println("create_menu_success")
	}
	return err
}
