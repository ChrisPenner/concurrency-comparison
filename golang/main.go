package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	concurrentMain()
}

func concurrentMain() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	eg := errgroup.Group{}
	for scanner.Scan() {
		path := scanner.Text()
		eg.Go(func() error {
			return readAndPrintFileLines(path)
		})
	}
	err := eg.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func syncMain() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		path := scanner.Text()
		err := readAndPrintFileLines(path)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func readAndPrintFileLines(path string) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	n := bytes.Count(body, []byte("\n"))
	fmt.Printf("%s: %d\n", path, n)
	return nil
}

func complexMain() error {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	eg := errgroup.Group{}
	for scanner.Scan() {
		path := scanner.Text()
		eg.Go(func() error {
			return readAndPrintFileLines(path)
		})
	}
	err := eg.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}
