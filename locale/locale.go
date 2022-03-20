package locale

type Message int

const (
	MessageWelcome Message = iota
	MessageKick
	MessageJoin
	MessageWrongAnswerLettersOnly
	MessageWrongAnswer
	MessageNonText
)
