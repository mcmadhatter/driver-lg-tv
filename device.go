package main

import (
	"github.com/mcmadhatter/go-lg-tv"
	"github.com/ninjasphere/go-ninja/api"
	"github.com/ninjasphere/go-ninja/config"
	"github.com/ninjasphere/go-ninja/devices"
	"github.com/ninjasphere/go-ninja/model"
)

type Device struct {
	devices.MediaPlayerDevice
	tv *lgtv.TV
}

func newDevice(driver ninja.Driver, conn *ninja.Connection, cfg *TVConfig) (*Device, error) {

	player, err := devices.CreateMediaPlayerDevice(driver, &model.Device{
		NaturalID:     cfg.ID,
		NaturalIDType: "samsung-tv",
		Name:          &cfg.Name,
		Signatures: &map[string]string{
			"ninja:manufacturer": "LG",
			"ninja:productName":  "Smart TV",
			"ninja:thingType":    "mediaplayer",
			"ip:mac":             cfg.ID,
		},
	}, conn)

	if err != nil {
		return nil, err
	}

	tv := lgtv.TV{
		Pin:             cfg.Pin,
		Name:            cfg.Name,
		ApplicationID:   config.MustString("userId"),
		ApplicationName: "Ninja Sphere",
		Id:              cfg.ID,
		Ip:              cfg.IP,
	}

	tv.PairWithPin()

	// Volume Channel
	player.ApplyVolumeUp = func() error {
		tv.SendCommandCode(lgtv.TV_CMD_VOLUME_UP)
		return nil
	}

	player.ApplyVolumeDown = func() error {
		tv.SendCommandCode(lgtv.TV_CMD_VOLUME_DOWN)
		return nil
	}

	player.ApplyToggleMuted = func() error {
		tv.SendCommandCode(lgtv.TV_CMD_MUTE_TOGGLE)
		return nil
	}

	if err := player.EnableVolumeChannel(false); err != nil {
		player.Log().Fatalf("Failed to enable volume channel: %s", err)
	}

	// Media Control Channel
	player.ApplyPlayPause = func(play bool) error {
		if play {
			tv.SendCommandCode(lgtv.TV_CMD_PLAY)
			return nil
		}

		tv.SendCommandCode(lgtv.TV_CMD_PAUSE)
		return nil
	}

	player.ApplyStop = func() error {
		tv.SendCommandCode(lgtv.TV_CMD_STOP)
		return nil
	}

	if err := player.EnableControlChannel([]string{}); err != nil {
		player.Log().Fatalf("Failed to enable control channel: %s", err)
	}

	return &Device{*player, &tv}, nil
}
