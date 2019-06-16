package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type InfoJson struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	ImgUrl     string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

var InfoUrl = "https://xkcd.com/%d/info.0.json"

// execute once
// cache comic json result in local dictionary
func cacheJson() {
	comicIndex := 405
	failNul := 0
	cacheDir := "./xkcd"
	// if dir not exist, create
	_, err := os.Stat(cacheDir)
	if err != nil {
		err = os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			log.Fatal("create cache dir %s fail\n", cacheDir)
		}
		log.Printf("mkdir %s success\n", cacheDir)
	} else {
		fmt.Printf("cache dictionary already exists, stop make dictionary")
	}
	for {
		resp, err := http.Get(fmt.Sprintf(InfoUrl, comicIndex))
		jsonPath := cacheDir + "/" + fmt.Sprintf("%d.json", comicIndex)
		comicIndex++
		if err != nil || resp.StatusCode != http.StatusOK {
			failNul++
			fmt.Printf("fail cache comic with index=%d\n", comicIndex)
			if failNul >= 10 {
				break
			} else {
				continue
			}
		}

		var info InfoJson
		if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
			log.Printf("Parse Json Fail with index=%d", comicIndex)
			continue
		}

		defer resp.Body.Close()
		content, err := json.MarshalIndent(info, " ", "    ")
		if err != nil {
			log.Printf("Mashal Json Fail with index=%d", comicIndex)
			continue
		}
		err = ioutil.WriteFile(jsonPath, content, os.ModePerm)
		if err != nil {
			fmt.Printf("Fail write content to json with index=%d\n", comicIndex)
		}
	}
}

func searchByCache() {
	cacheDir := "xkcd"
	dir, err := ioutil.ReadDir(cacheDir)
	if err != nil {
		log.Fatalf("Read dir fail, err=%v", err)
	}
	var comicMap = make(map[string]string)
	for _, fi := range dir {
		if strings.HasSuffix(fi.Name(), ".json") {
			filePath := cacheDir + string(os.PathSeparator) + fi.Name()
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("read %s fail\n", filePath)
				continue
			}
			var infoJosn InfoJson
			err = json.Unmarshal(content, &infoJosn)
			if err != nil {
				fmt.Printf("Parse %s fail, err=%v\n", filePath, err)
				continue
			}
			// unmarshal succ
			comicName := infoJosn.SafeTitle
			comicUrl := infoJosn.ImgUrl
			comicMap[comicName] = comicUrl
		}
	}
	fmt.Println("===Input Comic Name, And You Will Get Url Of This Comic;If Comic Not Exists in Our Cache,  You Will Receive Sorry===")
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanLines)
	for input.Scan() {
		key := input.Text()
		value, ok := comicMap[key]
		if !ok {
			fmt.Println("Sorry")
		} else {
			fmt.Println(value)
		}
	}
}
