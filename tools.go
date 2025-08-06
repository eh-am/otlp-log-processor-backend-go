//go:build tools
// +build tools

// Package tools is used to describe various tools we use.
// https://marcofranssen.nl/manage-go-tools-via-go-modules/
package tools

import (
	_ "github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
