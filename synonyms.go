package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string

	response, err := http.Get("https://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: %qの類似語検索に失敗しました: %v", term, err)
	}

	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)

	return syns, nil
}

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした \n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
