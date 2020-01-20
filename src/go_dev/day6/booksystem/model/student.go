package model

import "errors"

var (
	ErrNotFoundBook = errors.New("not found book")
	ErrBackOver     = errors.New("back over")
)

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	books []*BorrowItem
}

type BorrowItem struct {
	Book *Book
	Num  int
}

func CreateStudent(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		Id:    id,
		Sex:   sex,
		books: nil,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.books = append(s.books, b)
}

func (s *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].Book.Name == b.Book.Name {
			if b.Num == s.books[i].Num {
				front := s.books[0:i]
				end := s.books[i+1:]
				front = append(front, end...)
				s.books = front
				return
			}
		} else if b.Num < s.books[i].Num {
			s.books[i].Num -= 1
			return
		} else {
			return ErrBackOver
		}
	}
	err = ErrNotFoundBook
	return err
}

func (s *Student) GetBookList() []*BorrowItem {
	return s.books
}
