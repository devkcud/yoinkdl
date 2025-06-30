package filename

import (
	"fmt"
	"mime"
)

type Basic struct {
	Name      string
	Extension string
}

var Default = MustNew("download_file")

// If entryCopy is defined (aka >0), it appends '_{entryCopy}' between Name and Extension
func (b *Basic) Full(entryCopy ...int) string {
	if len(entryCopy) != 0 {
		return fmt.Sprintf("%s_%d%s", b.Name, entryCopy[0], b.Extension)
	}
	return fmt.Sprintf("%s%s", b.Name, b.Extension)
}

func (b *Basic) ToMime() string {
	return mime.TypeByExtension(b.Extension)
}
