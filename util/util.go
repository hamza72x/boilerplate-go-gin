package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func EnsureEnvs(envs ...string) error {
	missingEnvs := ""
	for _, env := range envs {
		if os.Getenv(env) == "" {
			if missingEnvs == "" {
				missingEnvs += env
				continue
			}
			missingEnvs += ", " + env
		}
	}
	if missingEnvs != "" {
		return fmt.Errorf("missing envs: %s", missingEnvs)
	}
	return nil
}

func RemoveFileIfExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(path)
}

// MustBytesReaderFromMap marshals the given map to json and returns it as a byte slice.
func MustBytesReaderFromMap(m gin.H) io.Reader {
	b, err := json.Marshal(m)
	if err != nil {
		panic("could not marshal map to json: " + err.Error())
	}
	return bytes.NewReader(b)
}

// InArray checks if the given needle is in the given haystack.
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}
