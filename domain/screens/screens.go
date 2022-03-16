package screens

import "github.com/novabankapp/ussdcommon/localisation"

type USSDOpType int

const (
	Request  USSDOpType = 0x1
	Notify   USSDOpType = 0x2
	Response USSDOpType = 0x3
	Release  USSDOpType = 0x4
)

type ServiceType int

const (
	WelcomeScreen ServiceType = iota
	HomeScreen
	EndApplication
)

type Screen struct {
	opType           USSDOpType
	ServiceType      ServiceType
	ScreenId         int
	Message          string
	ResponseRequired bool
	repeatNumber     int
	AllowAutoAnswers bool
	IsARepeat        bool
	ExpectedVariable *ExpectedVariable
}

func (s Screen) IncrementRepeat() *Screen {
	s.repeatNumber++
	return &s
}
func (s Screen) RepeatNumber() int {
	return s.repeatNumber
}
func (s Screen) OptType() USSDOpType {
	return s.opType
}
func (s Screen) setOptType(optType USSDOpType) {
	s.opType = optType
}
func (s Screen) GetExitScreen(language localisation.LanguageEnum) *Screen {
	//get message according to language
	message := ""
	return &Screen{
		ScreenId:         -1,
		Message:          message,
		ResponseRequired: false,
		ExpectedVariable: nil,
		IsARepeat:        false,
		ServiceType:      EndApplication,
	}
}

func NewScreen(
	allowAutoAnswers bool,
	serviceType ServiceType,
	screenId int,
	message string,
	variableName *ExpectedVariable,
	isARepeat bool,
	responseRequired bool) *Screen {
	return &Screen{
		AllowAutoAnswers: allowAutoAnswers,
		ScreenId:         screenId,
		Message:          message,
		ServiceType:      serviceType,
		ResponseRequired: responseRequired,
		ExpectedVariable: variableName,
		IsARepeat:        isARepeat,
	}

}
