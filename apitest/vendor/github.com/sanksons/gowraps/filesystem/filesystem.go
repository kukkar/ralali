package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// File not found
var ErrFileNotFound error = errors.New("File not found")

//Permission Denied
var ErrPermissionDenied error = errors.New("Permission Denied")

func CheckIfFileExists(filepath string) (bool, error) {
	_, err := os.Lstat(filepath)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetFile(filepath string) ([]byte, error) {

	f, err := os.Open(filepath)
	if err != nil {
		// file not exists
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		// permission issue.
		if os.IsPermission(err) {
			return nil, ErrPermissionDenied
		}
		return nil, err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	return data, nil
}

func DeleteFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil //donot send error
		}
		if os.IsPermission(err) {
			return fmt.Errorf("Donot have permission to delete")
		}
		return err
	}
	return nil
}

func CreateFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//
// Creates the file incase it does not exists.
//
func AppendToFile(filePath string, data []byte) error {

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}

//
// Create Directory Tree.
//
func CreateDirTree(dirPath string, perm int) error {
	return os.MkdirAll(dirPath, os.FileMode(perm))
}
