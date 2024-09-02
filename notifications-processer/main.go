package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	impScore := dm.priorityLevel
	if dm.isUrgent {
		impScore = 50
	}
	return impScore
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch not := n.(type) {
	case directMessage:
		return not.senderUsername, not.importance()
	case groupMessage:
		return not.groupName, not.importance()
	case systemAlert:
		return not.alertCode, not.importance()
	default:
		return "", 0
	}
}
