package reqres

type OutputTask1 struct {
	Number []int `json:"number"`
	Target int   `json:"target"`
	Output []int `json:"output"`
}

type OutputTask2 struct {
	Number []int   `json:"number"`
	Output [][]int `json:"output"`
}

type OutputTask3 struct {
	String string   `json:"string"`
	Words  []string `json:"words"`
	Output []int    `json:"output"`
}
