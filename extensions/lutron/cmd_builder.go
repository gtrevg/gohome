package lutron

import (
	"fmt"
	"io"
	"time"

	lutronExt "github.com/go-home-iot/lutron"
	"github.com/markdaws/gohome"
	"github.com/markdaws/gohome/attr"
	"github.com/markdaws/gohome/cmd"
)

type cmdBuilder struct {
	System *gohome.System
	device lutronExt.Device
}

func (b *cmdBuilder) Build(c cmd.Command) (*cmd.Func, error) {

	switch command := c.(type) {
	case *cmd.FeatureSetAttrs:
		f, ok := b.System.Features[command.FeatureID]
		if !ok {
			return nil, fmt.Errorf("unknown feature ID: %s", command.FeatureID)
		}

		d, ok := b.System.Devices[f.DeviceID]
		if !ok {
			return nil, fmt.Errorf("unknown device ID: %s", f.DeviceID)
		}

		var level float32 = -1
		for _, attribute := range command.Attrs {
			attribute := attribute
			switch attribute.Type {
			case attr.ATOnOff:
				if attribute.Value.(int32) == attr.OnOffOff {
					level = 0
				} else {
					level = 100
				}
			case attr.ATBrightness, attr.ATOffset:
				level = attribute.Value.(float32)
			}
		}

		if level == -1 {
			return nil, fmt.Errorf("unsupported attribute")
		}

		return &cmd.Func{
			Func: func() error {
				return getWriterAndExec(d, func(d lutronExt.Device, w io.Writer) error {
					return d.SetLevel(level, f.Address, w)
				})
			},
		}, nil

	default:
		return nil, fmt.Errorf("unsupported command type")
	}
	return nil, nil
}

func getWriterAndExec(d *gohome.Device, f func(lutronExt.Device, io.Writer) error) error {
	var hub *gohome.Device = d
	if d.Hub != nil {
		hub = d.Hub
	}

	conn, err := hub.Connections.Get(time.Second*5, true)
	if err != nil {
		return fmt.Errorf("error connecting, pool returned err: %s", err)
	}

	lDev, err := lutronExt.DeviceFromModelNumber(hub.ModelNumber)
	if err != nil {
		return err
	}

	err = f(lDev, conn)
	hub.Connections.Release(conn, err)
	if err != nil {
		return fmt.Errorf("Failed to send command %s\n", err)
	}
	return nil
}
