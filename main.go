package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const chanelSize = 10
const K = 5

func main() {
	te := InitTaskExecutor(K, chanelSize, task)


	//read stdin row by row and pushing strings to task executor
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		te.AppendData(url)
	}
	//wait until task executor finishes jobs
	te.Close()
}


func task(url string) error {
	//Fetch text from url
	text, err := getTextByUrl(url)
	if err != nil {
		return err
	}

	//get count of substrings "Go" in fetched result
	count := strings.Count(text, "Go")
	fmt.Printf("Count for %s: %d\n", url, count)
	return nil
}

//Fetch text from url
func getTextByUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	stringBody := string(body)
	return stringBody, nil
}
