package models

type (
	Book struct {
		BookId        string `json:"book_id"`
		Title         string `json:"title"`
		Author        string `json:"suthor"`
		PublisherYear int    `json:"published_year"`
	}

	CreateBookRequest struct {
		Title         string `json:"title"`
		Author        string `json:"suthor"`
		PublisherYear int    `json:"published_year"`
	}
	UpdateBookRequest struct {
		Title         string `json:"title"`
		Author        string `json:"suthor"`
		PublisherYear int    `json:"published_year"`
	}
	GetAllBooksRequest struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	GetBookByIdRequest struct {
		BookId string `json:"book_id"`
	}
	GetBooksByAuthorRequest struct {
		Author string `json:"author"`
	}
	GetBooksByNameRequest struct {
		BookName string `json:"book_name"`
	}
	GetSeveralResponse struct {
		Books []*Book `json:"books"`
	}
	SearchBooksRequest struct {
		Search string `json:"search"`
	}
)

/*
	CreateBook()
	UpdateBook()
	GetBookById()
	GetAllBooks()
	GetBooksByAuthor()
	GetBooksByName()
	SearchBooks()
*/
