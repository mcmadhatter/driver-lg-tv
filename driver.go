package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/mcmadhatter/go-lg-tv"
	"github.com/ninjasphere/go-ninja/api"
	"github.com/ninjasphere/go-ninja/logger"
	"github.com/ninjasphere/go-ninja/model"
	"github.com/ninjasphere/go-ninja/support"
)

var mydefaultpin = 123456 /* chnage this to the pin that gets shown on you screen the first time the driver is run */

var info = ninja.LoadModuleInfo("./package.json")
var log = logger.GetLogger(info.Name)

type Driver struct {
	support.DriverSupport
	config  Config
	devices map[string]*Device
}

type Config struct {
	TVs map[string]*TVConfig
}

func (c *Config) get(id string) *TVConfig {
	for _, tv := range c.TVs {
		if tv.ID == id {

			return tv
		}
	}
	return nil
}

type TVConfig struct {
	ID    string
	Name  string
	IP    net.IP
	Found bool
	Pin   string
}

func NewDriver() (*Driver, error) {

	driver := &Driver{
		devices: make(map[string]*Device),
	}

	err := driver.Init(info)
	if err != nil {
		log.Fatalf("Failed to initialize driver: %s", err)
	}

	err = driver.Export(driver)
	if err != nil {
		log.Fatalf("Failed to export driver: %s", err)
	}

	return driver, nil
}

func (d *Driver) deleteTV(id string) error {
	delete(d.config.TVs, id)

	err := d.SendEvent("config", &d.config)

	// TODO: Can't unexport devices at the moment, so we should restart the driver...
	go func() {
		time.Sleep(time.Second * 2)
		os.Exit(0)
	}()

	return err
}

func (d *Driver) saveTV(tvcfg TVConfig) error {

	existing := d.config.get(tvcfg.ID)

	if existing != nil {
		existing.Pin = tvcfg.Pin
		existing.Name = tvcfg.Name
		existing.IP = tvcfg.IP
		existing.ID = tvcfg.Name + tvcfg.Pin
	} else {
		tvcfg.ID = tvcfg.Name + tvcfg.Pin
		d.config.TVs[tvcfg.ID] = &tvcfg

		go d.createTVDevice(&tvcfg)
	}

	tv := lgtv.TV{}
	tv.Id = tvcfg.ID
	tv.Ip = tvcfg.IP
	tv.Name = tvcfg.Name
	tv.Pin = tvcfg.Pin
	fmt.Print("Save Config - ID:%s IP:%s\n", tv.Id, tvcfg.IP.String())
	tv.PairWithPin()

	return d.SendEvent("config", d.config)
}

func (d *Driver) Start(config *Config) error {

	fmt.Println("Driver Starting with config %+v", config)

	d.Conn.MustExportService(&configService{d}, "$driver/"+info.ID+"/configure", &model.ServiceAnnouncement{
		Schema: "/protocol/configuration",
	})

	if config.TVs == nil {
		config.TVs = make(map[string]*TVConfig)

		var tvcfg TVConfig

		tv := lgtv.TV{}
		tv.GetTVToShowPin()

		/* once you have the pin number for you lg tv uncomment the lines below, recompile, re-upload and reboot ) */

		tvcfg.Name = tv.Name
		tvcfg.IP = tv.Ip
		tvcfg.Found = true
		tvcfg.Pin = strconv.Itoa(mydefaultpin)

		tvcfg.ID = tv.Name + tvcfg.Pin

		config.TVs[tvcfg.ID] = &tvcfg

		fmt.Println("Added First TV config with Id " + tvcfg.ID)

	}

	d.config = *config

	for _, cfg := range config.TVs {

		fmt.Println("Creating device with TV name " + cfg.Name + " and id " + cfg.ID)

		d.createTVDevice(cfg)

	}

	return nil
}

func (d *Driver) requestPinShow() net.IP {
	tv := lgtv.TV{}
	tv.GetTVToShowPin()
	return tv.Ip

}

func (d *Driver) createTVDevice(cfg *TVConfig) {

	device, err := newDevice(d, d.Conn, cfg)

	if err != nil {
		log.Fatalf("Failed to create new LG TV device host:%s id:%s name:%s : %s", cfg.IP, cfg.ID, cfg.Name, err)
	}

	d.devices[cfg.ID] = device
}
