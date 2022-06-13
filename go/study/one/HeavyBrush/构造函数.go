package HeavyBrush

type File struct {
	fd      int
	name    string
	dirinfo []byte
	nepipe  int
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

func NewFile1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0}
	return &f
}

func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name, nil, 0}
}

func NewFile3(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd: fd, name: name}
}
