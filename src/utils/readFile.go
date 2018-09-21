package utils

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	logname    = "C:/goWorkSpace/src/logService/utils/log.log"
	concurrent = 5                //并发处理数,可以根据物理内存调整
	maxsize    = 10 * 1024 * 1024 //每次读取的大小,可以根据物理内存调整
)

var bufs = make([][]byte, concurrent)

func ReadFile2() {
	var (
		chs         = make(map[int]chan int)
		wait        = new(sync.WaitGroup)
		chct        = make(chan int, concurrent)
		ctx, cancel = context.WithCancel(context.Background())
	)

	File, err := os.Open(logname)
	if err != nil {
		log.Fatalf("Open file error,%s\n", err.Error())
	}

	for i := 0; i < concurrent; i++ {
		chct <- i
		chs[i] = make(chan int)
		bufs[i] = make([]byte, maxsize)
		go resolvectx(ctx, wait, i, chct, chs[i])
	}

	var i, n, l int
	for i = range chct {
		n, err = File.Read(bufs[i])
		if err != nil {
			wait.Wait() //等待数据全部处理完毕,然后返回
			break
		}
		for s := 1; s < n; s++ { //如果行过长,那么效率会变低
			if bufs[i][n-s] == '\n' {
				n = n - s + 1
				File.Seek(int64(l+n), 0)
				break
			}
		}
		l += n
		wait.Add(1)
		chs[i] <- n
	}
	cancel()
	close(chct)
	File.Close()
}

func resolvectx(ctx context.Context, wait *sync.WaitGroup, index int, chct, ch chan int) {
	var (
		err    error
		line   []byte
		length int
		buf    = bufio.NewReader(nil)
	)

	for {
		select {
		case <-ctx.Done():
			return
		case length = <-ch:
			buf.Reset(bytes.NewBuffer(bufs[index][:length]))
			for {
				line, _, err = buf.ReadLine()
				if err != nil {
					break
				}
				_ = line
			}
			chct <- index
			wait.Done()
		}
	}
}

func ListFile(folder string) (fileNames []string, err error) {
	folder = strings.TrimSpace(folder)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			ListFile(folder + "/" + file.Name())
		} else {
			fileNames = append(fileNames, folder+"/"+file.Name())
		}
	}
	return fileNames, nil
}
