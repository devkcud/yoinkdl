package filename_test

import (
	"runtime"
	"testing"

	"github.com/devkcud/yoinkdl/pkg/filename"
)

type validFileInfo struct {
	input, wantName, wantExt string
}

type invalidFileInfo struct {
	input   string
	wantErr error
}

func TestNew_ValidFilenames(t *testing.T) {
	tests := []validFileInfo{
		{"file.txt", "file", ".txt"},
		{"archive.tar.gz", "archive", ".tar.gz"},
		{"README", "README", ""},
		{".bashrc", ".bashrc", ""},
		{".config.json", ".config", ".json"},
	}

	for _, test := range tests {
		fb, err := filename.New(test.input)
		if err != nil {
			t.Errorf("New(%q) returned unexpected error: %v", test.input, err)
		}
		if fb.Name != test.wantName || fb.Extension != test.wantExt {
			t.Errorf("New(%q) = Name %q, Ext %q; want Name %q, Ext %q", test.input, fb.Name, fb.Extension, test.wantName, test.wantExt)
		}
	}
}

func TestNew_InvalidFilenames(t *testing.T) {
	tests := []invalidFileInfo{
		{"", filename.ErrEmptyName},
		{"   ", filename.ErrEmptyName},
		{".", filename.ErrInvalidName},
		{"..", filename.ErrInvalidName},
		{"...", filename.ErrInvalidName},
	}

	for _, test := range tests {
		fb, err := filename.New(test.input)
		if fb != nil {
			t.Errorf("New(%q) = %v; want nil FileBasic", test.input, fb)
		}
		if err != test.wantErr {
			t.Errorf("New(%q) error = %v; want %v", test.input, err, test.wantErr)
		}
	}
}

func TestFileBasic_Full(t *testing.T) {
	f := filename.MustNew("example.md")
	want := "example.md"
	if got := f.Full(); got != want {
		t.Errorf("Full() = %q; want %q", got, want)
	}
}

func TestFileBasic_FullCopy(t *testing.T) {
	f := filename.MustNew("example.md")
	want := "example_1.md"
	if got := f.Full(1); got != want {
		t.Errorf("Full() = %q; want %q", got, want)
	}
}

func TestReservedWindowsNames(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows reserved names test on non-Windows OS")
	}
	names := []string{"CON", "aux.txt", "nul"}
	for _, name := range names {
		_, err := filename.New(name)
		if err != filename.ErrReservedName {
			t.Errorf("New(%q) error = %v; want ErrReservedName", name, err)
		}
	}
}
