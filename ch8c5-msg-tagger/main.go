package main

import (
	"strings"
)

const (
	promoTag  = "Promo"
	urgentTag = "Urgent"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for idx, msg := range messages {
		messages[idx].tags = tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lwContent := strings.ToLower(msg.content)
	if strings.Contains(lwContent, "urgent") {
		tags = append(tags, urgentTag)
	}
	if strings.Contains(lwContent, "sale") {
		tags = append(tags, promoTag)
	}
	return tags
}
