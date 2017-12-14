package menu

import (
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
)

func Create(client *core.Client, menus string) error {
	me := new(menu.Menu)
	if err := json.Unmarshal([]byte(menus), me); err != nil {
		return err
	}
	return menu.Create(client, me)
}

func Delete(client *core.Client) error {

	return menu.Delete(client)
}
func DeleteButton(client *core.Client, button_key string) error {
	if this, _, err := menu.Get(client); err != nil {
		return err
	} else {
		if this == nil {
			return errors.New("null buttons.")
		}
		allbutton := this.Buttons[:]
		this.Buttons = make([]menu.Button, 0)
		var key_exist bool
		for _, bt := range allbutton {
			if len(bt.SubButtons) > 0 {
				buttons := bt.SubButtons[:]
				bt.SubButtons = make([]menu.Button, 0)
				for _, bt1 := range buttons {
					if button_key != bt.Key {
						bt.SubButtons = append(bt.SubButtons, bt1)
					} else {
						key_exist = true
					}
				}

			}
			if button_key != bt.Key {
				this.Buttons = append(this.Buttons, bt)
			} else {
				key_exist = true
			}
		}
		if key_exist {
			return nil
		}
		return errors.New("菜单不存在")
	}
}
func Get(client *core.Client) (*menu.Menu, error) {
	menu, _, err := menu.Get(client)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
