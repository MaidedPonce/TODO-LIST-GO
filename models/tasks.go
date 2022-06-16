package tasks

type Task struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
type List struct {
	Tasks []Task `json:"tasks"`
}
