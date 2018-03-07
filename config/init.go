package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type (
	Conf struct {
		FileName string
		Ctl      *ini.File
	}
)

func (c *Conf) Instance() {
	f, err := os.Open(c.FileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.Ctl, err = ini.Load(f)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func (c *Conf) GetSection(sec string) *ini.Section {
	c.checkInstance()

	section, err := c.Ctl.GetSection(sec)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return section

}

func (c *Conf) checkInstance() {
	if c.Ctl == nil {
		log.Fatal("need to init first")
	}
}
