package bahn

import (
	"bytes"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io"
	"regexp"
	"strings"
)

func HafasMessagesFromBytes(source []byte) ([]HafasMessage, error) {
	return HafasMessagesFromReader(bytes.NewReader(source))
}

var hafasMessageHighSelector = cascadia.MustCompile(".himMessagesHigh > div")
var hafasMessageMiddleSelector = cascadia.MustCompile(".himMessagesMiddle > div")
var hafasMessageLowSelector = cascadia.MustCompile(".himMessagesLow > div")

var hafasMessageValiditySelector = cascadia.MustCompile("span.bold")
var hafasMessageContentSelector = cascadia.MustCompile("span:not(.bold)")

var hafasMessageValidityRegex = regexp.MustCompile("^(?P<From>.+?)(?:\\p{Z}-\\p{Z}\\n(?P<To>.+?))?(?::\\p{Z}\\n(?P<Subject>.+?)\\.?)?$")
var hafasMessageIdRegex = regexp.MustCompile("^HIM_Text__(?P<Id>\\d+)$")

func HafasMessagesFromReader(source io.Reader) ([]HafasMessage, error) {
	var err error
	var messages []HafasMessage

	var document *html.Node
	if document, err = html.Parse(source); err != nil {
		return messages, err
	}

	parseMessage := func(node *html.Node, priority HafasMessagePriority) {
		validityNode := hafasMessageValiditySelector.MatchFirst(node)
		contentNode := hafasMessageContentSelector.MatchFirst(node)

		validity := strings.TrimSpace(parseText(validityNode))
		content := strings.TrimSpace(parseText(contentNode))

		var id string
		for _, attr := range node.Attr {
			if attr.Namespace == "" && attr.Key == "id" {
				id = attr.Val
			}
		}

		parsedId := parseRegexGroups(hafasMessageIdRegex, strings.TrimSpace(id))
		parsedValidity := parseRegexGroups(hafasMessageValidityRegex, strings.TrimSpace(validity))

		messages = append(messages, HafasMessage{
			Priority: priority,
			Id:       parsedId["Id"],
			From:     parsedValidity["From"],
			To:       parsedValidity["To"],
			Subject:  parsedValidity["Subject"],
			Content:  content,
		})
	}

	for _, node := range hafasMessageHighSelector.MatchAll(document) {
		parseMessage(node, HafasMessagePriorityHigh)
	}
	for _, node := range hafasMessageMiddleSelector.MatchAll(document) {
		parseMessage(node, HafasMessagePriorityMiddle)
	}
	for _, node := range hafasMessageLowSelector.MatchAll(document) {
		parseMessage(node, HafasMessagePriorityLow)
	}
	return messages, nil
}

func parseText(node *html.Node) string {
	var result string
	parseTextInternal(node, &result)
	return result
}

func parseTextInternal(node *html.Node, out *string) {
	if node.Type == html.TextNode {
		*out += node.Data
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			parseTextInternal(c, out)
		}
	}
}

func parseRegexGroups(regex *regexp.Regexp, url string) map[string]string {
	match := regex.FindStringSubmatch(url)

	paramsMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}
