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

func (dir directMessage) importance() int {
	if dir.isUrgent {
		return 50
	}
	return dir.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gr groupMessage) importance() int {

	return gr.priorityLevel

}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sys systemAlert) importance() int {

	return 100

}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0

	}
}
