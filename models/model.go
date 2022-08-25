package models

type (
	User struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Books []Book `json:"books"`
	}

	Book struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	DeleteBook struct {
		Id string `json:"id"`
	}

	Successful struct {
		Success string `json:"success"`
		Err     error  `json:"err"`
	}
)
