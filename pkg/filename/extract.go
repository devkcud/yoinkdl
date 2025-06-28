package filename

import (
	"path/filepath"
	"strings"
)

func extractSplitFile(filename string) (base, extension string) {
	extension = extractExtension(filename)
	return strings.TrimSuffix(filename, extension), strings.ToLower(extension)
}

func extractExtension(filename string) string {
	for _, ext := range compoundExtensions {
		if strings.HasSuffix(filename, ext) {
			return ext
		}
	}

	// Skip hidden files like: .bashrc, .zshrc, .waverc, etc
	if strings.Count(filename, ".") == 1 && strings.HasPrefix(filename, ".") {
		return ""
	}

	return filepath.Ext(filename)
}
