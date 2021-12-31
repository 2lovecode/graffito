package lib

import (
	"math"
	"strings"
)

type TfIdf struct {
	docs [][]string
}

func NewTfIdf(docs [][]string) *TfIdf {
	return &TfIdf{
		docs: docs,
	}
}

func (tfidf *TfIdf) Tf(index int, term string) float64 {
	termFrequency := float64(0)
	tf := float64(0)

	if len(tfidf.docs) > index {
		for _, eachTerm := range tfidf.docs[index] {
			if strings.ToLower(eachTerm) == term {
				termFrequency++
			}
		}
		tf = termFrequency / float64(len(tfidf.docs[index]))
	}

	return tf
}

func (tfidf *TfIdf) Df(term string) int {
	n := 0
	if term != "" {
		for _, eachDoc := range tfidf.docs {
			for _, eachTerm := range eachDoc {
				if strings.ToLower(eachTerm) == term {
					n++
					break
				}
			}
		}
	}
	return n
}

func (tfidf *TfIdf) Idf(term string) float64 {
	return math.Log(float64(len(tfidf.docs)) / float64((tfidf.Df(term) + 1)))
}

func (tfidf *TfIdf) TfIdf(index int, term string) float64 {
	return tfidf.Tf(index, term) * tfidf.Idf(term)
}
