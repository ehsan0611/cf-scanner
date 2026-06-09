package main

import (
	"log"
	"os"
	"sync"
)

type FileMutex struct {
	file   *os.File
	locker sync.Mutex
}

func (self *FileMutex) Write(data string) error {
	self.locker.Lock()
	defer self.locker.Unlock()

	_, err := self.file.WriteString(data)
	return err
}

func (self *FileMutex) Close() error {
	self.locker.Lock()
	defer self.locker.Unlock()

	return self.file.Close()
}

func resultFile(csv bool) *os.File {
	if csv {
		will_be_created := false
		_, exist := os.Stat("result.csv")
		if exist != nil {
			will_be_created = true
		}
		csv_file, err := os.OpenFile("result.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln(err)
		}
		if will_be_created {
			csv_file.Write([]byte("ip:port,ping,latency,jitter,download\n"))
		}
		return csv_file
	} else {
		file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln(err)
		}
		return file
	}
}
