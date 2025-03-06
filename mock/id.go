package mock

import "log"

var words []string = []string{
	"apple", "table", "chair", "book", "pen",
	"car", "fish", "ball", "tree", "bird",
	"shoe", "desk", "lamp", "phone", "shirt",
	"door", "wall", "plant", "grass", "clock",
}

func Id(n int) []string {
	if n > len(words) {
		log.Fatal("not enough ids")
	}

	return words[:n]
}
