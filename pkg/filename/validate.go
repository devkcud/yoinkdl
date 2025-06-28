package filename

import (
	"errors"
	"regexp"
	"runtime"
	"strings"
)

var (
	ErrInvalidName  error = errors.New("filename is invalid")
	ErrEmptyName    error = errors.New("filename is empty")
	ErrReservedName error = errors.New("filename is a reserved keyword")
)

var dotRegex = regexp.MustCompile(`^\.*$`)

var reservedWindowsNames = map[string]struct{}{
	"con": {}, "prn": {}, "aux": {}, "nul": {},
	"com1": {}, "com2": {}, "com3": {}, "com4": {}, "com5": {}, "com6": {}, "com7": {}, "com8": {}, "com9": {},
	"lpt1": {}, "lpt2": {}, "lpt3": {}, "lpt4": {}, "lpt5": {}, "lpt6": {}, "lpt7": {}, "lpt8": {}, "lpt9": {},
}

var compoundExtensions = []string{
	".tar.gz",
	".tar.bz2",
	".tar.xz",
	".tar.zst",
	".tar.lz",
	".tar.lzma",
	".tar.br",
	".tar.Z",
	".cpio.gz",
	".cpio.xz",
	".cpio.bz2",
	".cpio.lzma",
	".cpio.zst",
	".rpm.gz",
	".apk.tar.gz",
}

func isValidFilename(filename string) error {
	if strings.TrimSpace(filename) == "" {
		return ErrEmptyName
	}

	if dotRegex.MatchString(filename) {
		return ErrInvalidName
	}

	if runtime.GOOS == "windows" {
		filename = strings.ToLower(filename)
		_, filename = extractSplitFile(filename)
		if _, exists := reservedWindowsNames[filename]; exists {
			return ErrReservedName
		}
	}

	return nil
}
