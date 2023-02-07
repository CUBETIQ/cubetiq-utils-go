package file

import (
	"errors"
	"fmt"
	"math"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

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
	destFile := ""
	if strings.HasSuffix(destPath, "/") {
		destFile = destPath + name + "." + extension
	} else {
		destFile = destPath + "/" + name + "." + extension
	}

	// check destination of file not exists, it will create a directory
	if _, err := os.Stat(destPath); errors.Is(err, os.ErrNotExist) {
		// 0777 => public-read-write
		os.MkdirAll(destPath, 0777) // create directory
	}

	// save file
	formatFilePath := fmt.Sprint(destFile)

	c.SaveFile(f, formatFilePath)

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

func FormatFileSize(val float64) (size float64, unit, format, normalized string) {
	var round float64
	suffixes := [5]string{"B", "KB", "MB", "GB", "TB"}
	base := math.Log(val) / math.Log(1024)
	value := math.Pow(1024, base-math.Floor(base))
	pow := math.Pow(10, float64(2))
	digit := pow * value
	_, div := math.Modf(digit)
	if div >= .5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	size = round / pow
	unit = suffixes[int(math.Floor(base))]
	format = strconv.FormatFloat(size, 'f', -1, 64)
	normalized = format + " " + unit
	return
}
