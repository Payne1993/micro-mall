package handlers

import (
	wordsv1 "micro-mall.dev/app/word/api/words/v1"

	"micro-mall.dev/app/grpc-gateway/internal/registry"
)

func init() {
	registry.Add("word", wordsv1.RegisterWordsHandler)
}
