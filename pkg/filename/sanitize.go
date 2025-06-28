package filename

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	nameRegex            = regexp.MustCompile(`[^a-zA-Z0-9._-]`)
	multiUnderscoreRegex = regexp.MustCompile(`_+`)
)

func New(filename string) (*Basic, error) {
	if strings.TrimSpace(filename) == "" {
		return nil, ErrEmptyName
	}

	filename = cleanupFilename(filename)

	if err := isValidFilename(filename); err != nil {
		return nil, err
	}

	base, extension := extractSplitFile(filename)

	return &Basic{
		Name:      base,
		Extension: extension,
	}, nil
}

// Panics if the filename is invalid.
func MustNew(filename string) *Basic {
	fb, err := New(filename)
	if err != nil {
		panic(err)
	}
	return fb
}

func cleanupFilename(filename string) string {
	filename = filepath.Base(filename)
	filename = nameRegex.ReplaceAllString(filename, "_")
	filename = multiUnderscoreRegex.ReplaceAllString(filename, "_")
	filename = strings.Trim(filename, "_")

	return filename
}
