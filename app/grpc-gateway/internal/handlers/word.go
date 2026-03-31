package handlers

import (
	wordsv1 "proxima/app/word/api/words/v1"

	"proxima/app/grpc-gateway/internal/registry"
)

func init() {
	registry.Add("word", wordsv1.RegisterWordsHandler)
}
