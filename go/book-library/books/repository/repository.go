package repository

import (
	"context"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/where"

	"github.com/pankajworks/Code/go/book_library/books"
)

type BooksRepository struct {
	repo rel.Repository
}

func New(repo rel.Repository) BooksRepository {
	return BooksRepository{
		repo: repo,
	}
}

func (r BooksRepository) FindShelf(ctx context.Context, shelfID int) (books.Shelf, error) {
	var shelf books.Shelf
	err := r.repo.Find(ctx, &shelf, where.Eq("id", shelfID))
	if err != nil {
		return books.Shelf{}, err
	}
	err = r.repo.Preload(ctx, &shelf, "books")
	return shelf, err
}

func (r BooksRepository) InsertBook(ctx context.Context, book books.Book) error {
	return r.repo.Transaction(ctx, func(ctx context.Context) error {
		book.ShelfID = book.BookShelf.ID
		r.repo.MustInsert(ctx, &book, rel.Cascade(false))
		return r.repo.Update(ctx, &book.BookShelf, rel.Inc("amount"), rel.Reload(false))
	})
}

func (r BooksRepository) FindBookByID(ctx context.Context, bookID int) (books.Book, error) {
	var book books.Book
	err := r.repo.Find(ctx, &book, rel.Eq("id", bookID))
	if err != nil {
		return books.Book{}, err
	}
	err = r.repo.Preload(ctx, &book, "book_shelf")
	if err != nil {
		return books.Book{}, err
	}
	return book, nil
}

func (r BooksRepository) UpdateBook(ctx context.Context, book books.Book) error {
	return r.repo.Update(
		ctx,
		&book,
		rel.Set("title", book.Title),
		rel.Set("description", book.Description),
		rel.Set("author", book.Author),
		rel.Set("edition", book.Edition),
		rel.Reload(false),
	)
}

func (r BooksRepository) FindAllBooks(ctx context.Context) ([]books.Book, error) {
	var bookList []books.Book
	err := r.repo.FindAll(ctx, &bookList)
	if err != nil {
		return nil, err
	}
	for i := range bookList {
		err = r.repo.Preload(ctx, &bookList[i], "book_shelf")
		if err != nil {
			return nil, err
		}
	}
	return bookList, err
}

func (r BooksRepository) DeleteBook(ctx context.Context, book books.Book) error {
	return r.repo.Delete(ctx, &book)
}
