package goHelp

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// https://zetcode.com/golang/copyfile/
func backupFile(src string) error {
	fin, err := os.Open(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		err = createFile(src)
		fmt.Fprintln(os.Stderr, "Error:", err)
		fin, err = os.Open(src)
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	defer fin.Close()

	dst := fin.Name() + ".bak"
	fout, err := os.Create(dst)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	return nil
}

func createFile(filepath string) error {
	filePtr, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File Creation Error:", err)
		//log.Fatal(err);
	}
	defer filePtr.Close()
	return nil
}

func checkFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		log.Fatal(err)
		return false, err
	}
	if os.IsNotExist(err) {
		log.Fatal("File path does not exist")
		return false, err
	}
	return true, nil
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func appendToFile(content, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(content))
}

// petergrep
// modified from https://stackoverflow.com/questions/26709971/could-this-be-more-efficient-in-go
// make optional http or dotTLDMap
func grepFile(file string, patterns []byte) (int64 map[int]string) {
	//pattersMap http, dotTLDmap
	patCount := int64(0)
	artifacts := make(map[int]string)
	builder := strings.Builder{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i := 0; i <= len(patterns)-1; i++ {
			if bytes.Contains(scanner.Bytes(), patterns[i]) {
				patCount++
				builder.WriteString("Scan " + file + " contains " + string(patterns[i]))
				artifacts[patCount] = builder.String()
				builder.Reset()
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return patCount, artifacts
}

func encodeGzipStr(s string) (string, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write([]byte(s)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func decodeGzipStr(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buf.Write(decodedBytes)

	gz, err := gzip.NewReader(&buf)
	if err != nil {
		return "", err
	}
	defer gz.Close()

	decompressedBytes, err := io.ReadAll(gz)
	if err != nil {
		return "", err
	}

	return string(decompressedBytes), nil
}

// petergrep
// modified from https://stackoverflow.com/questions/26709971/could-this-be-more-efficient-in-go
func grepFileForFSBits(file string, fsBits) error {
	f, err := os.Open(file)
	if err != nil {
		err.errorHandling(err, 2)
        return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i := 0; i <= 16 -1; i++ {
			if bytes.Contains(scanner.Bytes(), jsonBitID) {
				builder.WriteString("Scanned " + file + " which contains " + string(fsBits) + "at: " string(f.[i]))
				log.Println(builder.String())
				builder.Reset()
			}
		}
	}
	if err := scanner.Err(); err != nil {
		err.errorHandling(err, 2)
        return err
	}
	return patCount, artifacts
}
