package config

import (
	"errors"
	"flag"
)

// Flags defines a structure used to hold and pass parsed command line flags
type Flags struct {
	Fetch   bool   // download to computer
	Upload  bool   // upload to host
	Write   bool   // overwrite page
	Anon    bool   // operate without login
	Host    string // site url to use
	Title   string // page title to work with
	Pageid  string // page id to work with
	Infile  string // file to read content from
	Outfile string // file to write content to
}

// assembleMutual takes a dirty Flags struct and returns a new one which has
// basic mutual fields validated such as the Host and the Title or ID
func assembleMutual(flags *Flags) (*Flags, error) {
	clean := Flags{false, false, false, false, "", "", "", "", ""}

	// Required: Host is required
	if flags.Host != "" {
		clean.Host = flags.Host
	} else {
		return nil, errors.New("[mutual][required] Host field cannot be nil")
	}

	// Required: Title or Pageid is required
	if flags.Title != "" {
		clean.Title = flags.Title
	} else if flags.Pageid != "" {
		clean.Pageid = flags.Pageid
	} else {
		return nil, errors.New("[mutual][required] Title or Page ID field cannot be nil")
	}

	return &clean, nil
}

// assembleFetch takes a dirty Flags struct and returns a validated clean
// struct suitable for the fetch action
func assembleFetch(flags *Flags) (*Flags, error) {
	clean, err := assembleMutual(flags)
	if err != nil {
		return clean, err
	}

	// set operation
	clean.Fetch = true

	// Required: Outfile is required
	if flags.Outfile != "" {
		clean.Outfile = flags.Outfile
	} else {
		return nil, errors.New("[fetch][required] Outfile field cannot be nil")
	}

	return clean, nil
}

// assembleUpload takes a dirty Flags struct and returns a validated clean
// struct suitable for the upload action
func assembleUpload(flags *Flags) (*Flags, error) {
	clean, err := assembleMutual(flags)
	if err != nil {
		return clean, err
	}

	// set operation
	clean.Upload = true

	// Required: Outfile is required
	if flags.Infile != "" {
		clean.Infile = flags.Infile
	} else {
		return nil, errors.New("[upload][required] Infile field cannot be nil")
	}

	// Assign optional fields
	clean.Anon = flags.Anon
	clean.Write = flags.Write

	return clean, nil
}

// validateOperations returns whether or not a valid operation combination has been
// configured and returns a new clean Flags struct.
func validateOperations(flags *Flags) (*Flags, error) {
	if flags.Fetch && !flags.Upload {
		return assembleFetch(flags)
	} else if !flags.Fetch && flags.Upload {
		return assembleUpload(flags)
	}

	return nil, errors.New("[validate] Fetch or Upload field cannot both be true/false")
}

// New creates a new instance of Flags populated with the parsed command line
// flags supplied by the user
func New() (*Flags, error) {
	flags := Flags{
		Fetch:   false,
		Upload:  false,
		Write:   false,
		Anon:    false,
		Host:    "",
		Title:   "",
		Pageid:  "",
		Infile:  "",
		Outfile: "",
	}

	// add flag pointers to the parser
	flag.BoolVar(&flags.Fetch, "f", false, "Download to computer")
	flag.BoolVar(&flags.Upload, "u", false, "Upload to host")
	flag.BoolVar(&flags.Write, "O", false, "Overwrite page content")
	flag.BoolVar(&flags.Anon, "A", false, "Operate without logging in")
	flag.StringVar(&flags.Host, "s", "", "Host url")
	flag.StringVar(&flags.Title, "t", "", "Page title to work with")
	flag.StringVar(&flags.Pageid, "I", "", "Page ID to work with")
	flag.StringVar(&flags.Infile, "i", "", "File to upload content from")
	flag.StringVar(&flags.Outfile, "o", "", "File to download content to")

	// parse the flags
	flag.Parse()

	// validate the flags
	clean, err := validateOperations(&flags)

	return clean, err
}
