package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("checktool v1.1")
	fmt.Println("--------")
	Listname := "all.list"
	if _, err := os.Stat(Listname); err != nil {
		log.Fatal("Cannot find all.list")
	}
	file, err := os.Open("all.list")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var failures []string
	var fileCount,passCount,failCount int
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}
		//TODO : bugs 因为采用空格分割，所以文件名有空格就会失败
		words := strings.Fields(scanner.Text())
		if len(words) < 1{
			break
		}
		fileCount++
		if err := CompareFileMD5(words[1],words[0]);err != nil{
			failures = append(failures, words[1])
			failCount++
			fmt.Print("[FAILED] ")
		}else{
			passCount++
			fmt.Print("[PASS] ")
		}
		fmt.Println(words[1])
	}

	if len(failures) > 0{
		//写失败列表
		file2, err := os.Create("failed.list")
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()

		w2 := bufio.NewWriter(file2)
		for _, failure := range failures {
			//fmt.Println(path.Base(failure))  //文件名
			w2.WriteString(fmt.Sprintf("%s\n", failure))
		}
		defer w2.Flush()
	}

	fmt.Println("--------")
	fmt.Println("Finished.")
	fmt.Println("ALL:",fileCount,"/","PASS:",passCount,"/","FAIL:",failCount)
	fmt.Println("Exit after 5 seconds...")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

func CompareFileMD5(filepath, filemd5 string) error {
	if _, err := os.Stat(filepath); err != nil {
		return errors.New("File not found")
	}
	// path/to/whatever exists
	f, err := os.Open(filepath)
	if err != nil {
		return errors.New("Failed to open file")
	}
	defer f.Close()
	h := md5.New()
	io.Copy(h, f)
	//return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
	bytemd5,_ := hex.DecodeString(filemd5)
	if bytes.Compare(h.Sum(nil),bytemd5) != 0{
		return errors.New("Compare not OK")
	}
	return nil
}