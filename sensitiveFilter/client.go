package sensitiveFilter

import (
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) FindSensitiveWordsBySlices(sentence string, sensitiveWordsSlices []string) []string {
	filter := stringx.NewTrie(sensitiveWordsSlices)
	return c.filterKeywords(sentence, filter)
}
func (c *Client) ValidateBySlices(sentence string, sensitiveWordsSlcies []string) bool {
	filter := stringx.NewTrie(sensitiveWordsSlcies)
	return c.validate(sentence, filter)
}

func (c *Client) FindSensitiveWordsByStringSplit(sentence string, sensitiveStr string, split string) []string {
	return c.FindSensitiveWordsBySlices(sentence, strings.Split(sensitiveStr, split))
}
func (c *Client) ValidateByStringSplit(sentence string, sensitiveStr string, split string) bool {
	return c.ValidateBySlices(sentence, strings.Split(sensitiveStr, split))
}

func (c *Client) validate(sentence string, filter stringx.Trie) bool {
	return len(filter.FindKeywords(sentence)) < 1
}

func (c *Client) filterKeywords(sentence string, filter stringx.Trie) []string {
	return filter.FindKeywords(sentence)
}
