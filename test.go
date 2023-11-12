package hostsfile

import "os"

// TestFakeReadFile is a fake method to test the package.
func TestFakeReadFile(method func(name string) ([]byte, error)) {
	readFile = method
}

// TestFakeWriteFile is a fake method to test the package.
func TestFakeWriteFile(method func(name string, data []byte, perm os.FileMode) error) {
	writeFile = method
}

// TestFakeStatFile is a fake method to test the package.
func TestFakeStatFile(method func(name string) (os.FileInfo, error)) {
	statFile = method
}
