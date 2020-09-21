// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	books "github.com/pankajworks/Code/go/book_library/books"

	mock "github.com/stretchr/testify/mock"
)

// Finder is an autogenerated mock type for the Finder type
type Finder struct {
	mock.Mock
}

// FindBookByID provides a mock function with given fields: ctx, bookID
func (_m *Finder) FindBookByID(ctx context.Context, bookID int) (books.Book, error) {
	ret := _m.Called(ctx, bookID)

	var r0 books.Book
	if rf, ok := ret.Get(0).(func(context.Context, int) books.Book); ok {
		r0 = rf(ctx, bookID)
	} else {
		r0 = ret.Get(0).(books.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
