// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	books "github.com/pankajworks/Code/go/book_library/books"

	mock "github.com/stretchr/testify/mock"
)

// Deleter is an autogenerated mock type for the Deleter type
type Deleter struct {
	mock.Mock
}

// DeleteBook provides a mock function with given fields: ctx, book
func (_m *Deleter) DeleteBook(ctx context.Context, book books.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, books.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
