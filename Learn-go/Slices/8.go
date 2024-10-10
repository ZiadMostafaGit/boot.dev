package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var res []Message
	if filterType == "text" {

		for _, element := range messages {

			if element.Type() == "text" {

				res = append(res, element)
			}

		}
		return res

	}

	if filterType == "media" {

		for _, element := range messages {

			if element.Type() == "media" {

				res = append(res, element)
			}

		}
		return res

	}

	if filterType == "link" {

		for _, element := range messages {

			if element.Type() == "link" {

				res = append(res, element)
			}

		}
		return res

	}

	return nil

}
