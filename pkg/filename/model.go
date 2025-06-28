package filename

type FileBasic struct {
	Name      string
	Extension string
}

func (f *FileBasic) Full() string {
	return f.Name + f.Extension
}

func (f *FileBasic) String() string {
	return f.Full()
}
