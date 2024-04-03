package fileutil

import (
	"log"
	"testing"
)

func TestDir(t *testing.T) {
	log.Println("MkDirAll", MkDirAll("test"))
	log.Println("IsExist", IsExist("test"))
	log.Println("GetAbsDirPath test", GetAbsDirPath("test"))
	log.Println("GetAbsDirPath ../test", GetAbsDirPath("../test"))
	log.Println("DelDir", DelDir("test"))
}
