package crawler

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"strings"
)

func getTitle(c *colly.Collector, title *string) {
	c.OnHTML(".word-head .title", func(e *colly.HTMLElement) {
		e.DOM.Contents().EachWithBreak(func(i int, s *goquery.Selection) bool {
			if s.Nodes[0].Type == html.TextNode {
				text := strings.TrimSpace(s.Text())
				if text != "" {
					*title = text
					return false // break the loop
				}
			}
			return true // continue the loop
		})
	})
}

func getPhonetic(c *colly.Collector, phonetic *string) {
	var count int
	c.OnHTML(".per-phone", func(e *colly.HTMLElement) {
		count++
		if count == 2 {
			*phonetic = e.Text
		}
	})
}

type Definition struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Description  string `json:"description"`
}

func getDefinition(c *colly.Collector, definition *[]Definition) {
	c.OnHTML(".basic .word-exp", func(e *colly.HTMLElement) {
		partOfSpeech := e.ChildText(".pos")
		description := e.ChildText(".trans")
		*definition = append(*definition, Definition{
			PartOfSpeech: partOfSpeech,
			Description:  description,
		})
	})
}

type Phrase struct {
	EnglishPhrase      string `json:"englishPhrase"`
	ChineseTranslation string `json:"chineseTranslation"`
}

func getPhrase(c *colly.Collector, phrase *[]Phrase) {
	var count int
	c.OnHTML(".webPhrase .mcols-layout", func(e *colly.HTMLElement) {
		if count < 2 {
			englishPhrase := e.ChildText(".point")
			chineseTranslation := e.ChildText(".sen-phrase")
			*phrase = append(*phrase, Phrase{
				EnglishPhrase:      englishPhrase,
				ChineseTranslation: chineseTranslation,
			})
			count++
		}
	})
}

type Example struct {
	Sentence string `json:"sentence"`
}

func getExample(c *colly.Collector, example *[]Example) {
	var count int
	c.OnHTML(".catalogue_sentence .dict-book .mcols-layout", func(e *colly.HTMLElement) {
		if count < 2 {
			sentence := e.ChildText(".word-exp .sen-eng")
			*example = append(*example, Example{
				Sentence: sentence,
			})
			count++
		}
	})
}

type Translate struct {
	Title      string
	Phonetic   string
	Definition string
	Phrase     string
	Example    string
}

func GetTranslate(url string) Translate {
	c := colly.NewCollector()
	title := ""
	phonetic := ""
	var definition []Definition
	var phrase []Phrase
	var example []Example
	getTitle(c, &title)
	getPhonetic(c, &phonetic)
	getDefinition(c, &definition)
	getPhrase(c, &phrase)
	getExample(c, &example)
	c.Visit(url)
	c.Wait()
	jsonDefinition, _ := json.Marshal(definition)
	jsonPhrase, _ := json.Marshal(phrase)
	jsonExample, _ := json.Marshal(example)
	// 构建一个结构体返回
	translate := Translate{
		Title:      title,
		Phonetic:   phonetic,
		Definition: string(jsonDefinition),
		Phrase:     string(jsonPhrase),
		Example:    string(jsonExample),
	}
	return translate
}
