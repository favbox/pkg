package object

import (
	"github.com/bytedance/sonic"
	"io"
	"io/fs"
	"os"
)

func JsonEncode(v any) (string, error) {
	buffer, err := sonic.Marshal(v)

	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func JsonDecode(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func JsonEscape(str string) (string, error) {
	b, err := sonic.Marshal(str)
	if err != nil {
		return "", err
	}
	return string(b[1 : len(b)-1]), err
}

func LoadObjectFromFile(jsonPath string, obj any) (err error) {
	// open json file
	jsonFile, err := os.Open(jsonPath)

	defer jsonFile.Close()
	if err != nil {
		return err
	}

	// parse file to buffer
	byteValue, _ := io.ReadAll(jsonFile)

	// parse buffer to object
	err = sonic.Unmarshal(byteValue, obj)

	return err
}

func SaveObjectToFile(obj any, filePath string, perm fs.FileMode) (err error) {
	buff, err := sonic.MarshalIndent(obj, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, buff, perm)

	return err
}
