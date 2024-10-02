package main

import "embed"

// Embed a single file
//
//go:embed ui/dist/*
var f embed.FS
