package model

type Author struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Books *[]Book `json:"books"`
}

var stephenKingBooks = []Book{
	Book{ID: "1", Name: "Outsider"},
	Book{ID: "2", Name: "On Writing"},
}

var Authors = []Author{
	Author{ID: "1", Name: "Stephen King", Books: &stephenKingBooks},
}
