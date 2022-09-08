package file

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	stringutil "github.com/cubetiq/cubetiq-utils-go/string"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveMultipartToFile(c *fiber.Ctx, f *multipart.FileHeader, path string, filename string) *multipart.FileHeader {
	// check if filename is empty, it will put uuid to filename
	name := ""
	if stringutil.IsEmpty(filename) {
		name = uuid.New().String()
	} else {
		name = filename
	}

	// check if path is empty, it will put ./ in the root directory of project
	destPath := ""
	if stringutil.IsEmpty(path) {
		destPath = "./"
	} else {
		destPath = path
	}

	// get source filename to split a filename and take only extension
	sourceFileName := f.Filename
	extension := stringutil.GetPartOfLast(sourceFileName, ".")

	// sum up for destination of file
	destFile := destPath + "\\" + name + "." + extension

	// check destination of file not exists, it will create a directory
	if _, err := os.Stat(destPath); errors.Is(err, os.ErrNotExist) {
		// 0777 => public-read-write
		os.MkdirAll(destPath, 0777)
	}

	// save file
	c.SaveFile(f, fmt.Sprint(destFile))

	return f
}

// Remove a single file
func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
}

// Remove an entire directory
func RemoveDirectory(directory string) {
	err := os.RemoveAll(directory)
	if err != nil {
		fmt.Println(err)
	}
}
