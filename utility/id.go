package utility

import (
	"github.com/google/uuid"
)

func NewID(module ...string) string {
	var title string 
	for i, s := range module {
		if i == 0 {
			title = s
		} else {
			title = title + "-" + s		
		}
	}

	id := uuid.NewString()

	if len(module) == 0 {
		return id
	} else {
		return title + "@" + id
	}
}