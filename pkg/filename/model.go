package filename

import "fmt"

type FileBasic struct {
	Name      string
	Extension string
}

// If entryCopy is defined (aka >0), it appends '_{entryCopy}' between Name and Extension
func (f *FileBasic) Full(entryCopy ...int) string {
	if len(entryCopy) != 0 {
		return fmt.Sprintf("%s_%d%s", f.Name, entryCopy[0], f.Extension)
	}
	return fmt.Sprintf("%s%s", f.Name, f.Extension)
}

func NewDefault() *FileBasic {
	return &FileBasic{Name: "download_file", Extension: ".goondl"}
}
