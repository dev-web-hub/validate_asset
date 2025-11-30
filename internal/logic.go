package internal

type Input struct {
    Source string `json:"source"`
    Target string `json:"target"`
}

func Run(input Input) error {
    return nil
}
