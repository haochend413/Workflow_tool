package defs

type NoteSendMsg = *Note
type TopicSendMsg = *Topic
type CurrentViewMsg = string

type InitMsg struct{}
type SwitchContextMsg struct{}
type DeleteNoteMsg struct{}

/*
Daily Task
*/
type TaskSucMsg struct{}

type DeleteTaskMsg struct{}
