package main

import (
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli"
)

func deployToNoProvider(ctx *cli.Context) error {
	// Create node directory
	name := ctx.String("name")
	tags := ctx.String("tags")
	if err := mkdir(name, tags); err != nil {
		return err
	}
	nodePath := nodePath(name)

	// Generate config and ssh key for the node
	config, err := GetConfigOrGenerateNew(ctx, nodePath)
	if err != nil {
		return err
	}
	key, err := NewSshKeyPair(nodePath)
	if err != nil {
		return err
	}

	if _, err = os.Create(path.Join(nodePath, "main.tf")); err != nil {
		return err
	}

	if err := runTerraform(nodePath); err != nil {
		return err
	}

	file, err := os.Create(path.Join(nodePath, "multiAddress.out"))
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintf(file, "/ip4/0.0.0.0/tcp/18514/republic/%s", config.Address)

	return outputURL(nodePath, name, key.Marshal())
}
