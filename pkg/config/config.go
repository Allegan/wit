package config

import (
	"flag"
)

// Config holds application per-run configuration
type Config struct {
	Fetch     bool
	Upload    bool
	Overwrite bool
	Anonymous bool
	Site      string
	Title     string
	Input     string
	Output    string
}

// New returns a new instance of Config
func New() *Config {
	config := Config{
		Fetch:     false,
		Upload:    false,
		Overwrite: false,
		Anonymous: false,
		Site:      "",
		Title:     "",
		Input:     "",
		Output:    ""}

	flag.BoolVar(&config.Fetch, "f", false, "Download page from wiki")
	flag.BoolVar(&config.Upload, "u", false, "Upload page to wiki")
	flag.BoolVar(&config.Overwrite, "O", false, "Overwrite page content")
	flag.BoolVar(&config.Anonymous, "A", false, "Perform uploads anonymously")
	flag.StringVar(&config.Site, "s", "", "Wiki url")
	flag.StringVar(&config.Title, "t", "", "Page title to fetch or upload to")
	flag.StringVar(&config.Input, "i", "", "File to read for upload")
	flag.StringVar(&config.Output, "o", "", "File to output fetch to")

	flag.Parse()

	return &config
}
