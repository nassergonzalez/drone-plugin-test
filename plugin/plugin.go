// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"github.com/peterhellberg/swapi"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// TODO replace or remove
	Param1 string `envconfig:"PLUGIN_PARAM1"`
	Param2 string `envconfig:"PLUGIN_PARAM2"`
}

// Exec executes the plugin.
func Exec(_ context.Context, args Args) error {
	c:= swapi.DefaultClient
	planet, err := c.Planet(args.Build.Number)
	if err != nil {
		return err
	}
	fmt.Println(planet.Name)
	return nil
}
