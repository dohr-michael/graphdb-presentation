package importer

import (
	"net/http"
	"fmt"
	"strings"
	"io"
	"strconv"
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
	"encoding/csv"
	"log"
	"os"
	"io/ioutil"
)

var languagesMap = map[int]string{
	1:  "ja",
	2:  "ja",
	3:  "ko",
	4:  "zh",
	5:  "fr",
	6:  "de",
	7:  "es",
	8:  "it",
	9:  "en",
	10: "cs",
	11: "ja",
	12: "zh",
}

type row map[string]string

func parseId(base string) pokeapi.Id {
	return pokeapi.Id(parseInt(base))
}

func parseInt(base string) int {
	value, err := strconv.ParseInt(parseString(base), 10, 32)
	if err != nil {
		return 0
	}
	return int(value)
}

func parseString(base string) string {
	return strings.TrimSpace(base)
}

type loaderParams struct {
	BaseUrl              string
	TranslationUrl       string
	IdentifierField      string
	IdField              string
	TranslationIdField   string
	TranslationLangField string
	TranslationField     string
	Factory              func(id int, identifier string) pokeapi.WithIdentity
}

func translationLoader(params loaderParams, process func(item pokeapi.WithIdentity, row map[string]string)) {
	byId := make(map[pokeapi.Id]pokeapi.WithIdentity)
	downloadAndRead(params.BaseUrl, func(row row) {
		item := params.Factory(parseInt(row[params.IdField]), parseString(row[params.IdentifierField]))
		process(item, row)
		byId[item.GetId()] = item

	})
	downloadAndRead(params.TranslationUrl, func(row row) {
		genId := parseInt(row[params.TranslationIdField])
		langId := parseInt(row[params.TranslationLangField])
		t := parseString(row[params.TranslationField])
		if byId[pokeapi.Id(genId)] != nil && languagesMap[langId] != "" {
			byId[pokeapi.Id(genId)].SetName(parseString(languagesMap[langId]), t)
		}
	})
}

func downloadAndRead(url string, onRow func(row row)) {
	onError := func(err error) {
		log.Println(err)
	}
	idx := strings.LastIndex(url, "/")
	filename := ""
	if idx > -1 {
		filename = url[idx+1:]
	}

	if _, err := os.Stat("target/" + filename); os.IsNotExist(err) {
		fmt.Println("Download file", url)
		httpRes, err := http.Get(url)
		if err != nil {
			onError(err)
			return
		}
		if httpRes.StatusCode >= 400 {
			fmt.Printf("failed to download %d (%s)\n", httpRes.StatusCode, url)
			onError(err)
			return
		}
		defer httpRes.Body.Close()
		bytes, err := ioutil.ReadAll(httpRes.Body)
		if err != nil {
			onError(err)
			return
		}

		f, err := os.Create("target/" + filename)
		if err != nil {
			onError(err)
			return
		}
		_, err = f.Write(bytes)
		if err != nil {
			onError(err)
			return
		}
		f.Close()

	} else {
		fmt.Println("Read cached file ", "target/"+filename)
	}

	cachedFile, err := os.Open("target/" + filename)
	if err != nil {
		onError(err)
		return
	}
	defer cachedFile.Close()

	csvReader := csv.NewReader(cachedFile)

	header, err := csvReader.Read()
	if err != nil {
		onError(err)
		return
	}

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			onError(err)
			break
		}
		r := make(row)
		for idx, c := range line {
			r[header[idx]] = c
		}
		onRow(r)
	}
	return
}
