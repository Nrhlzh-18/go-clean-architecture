package usecases

type BookInteractor struct {
	BookRepository domain.BookRepository
}


func NewBookInteractor(repository domain.)