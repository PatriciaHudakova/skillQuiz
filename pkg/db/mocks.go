package db

import (
	"database/sql"
)

// MockClient mock.
type MockClient struct {
	GetAllRowsFn                  func() (*sql.Rows, error)
	MakeCurrentRatingTheAverageFn func(currentRating string) error
	GetOverallAverageFromDBFn     func() (int, error)
	UpdateAverageFn               func(newAverage int) error
	IsEmptyFn                     func(rows *sql.Rows) bool
}

// NewMockClient mock.
func NewMockClient() *MockClient {
	return &MockClient{
		GetAllRowsFn: func() (*sql.Rows, error) {
			return &sql.Rows{}, nil
		},
		MakeCurrentRatingTheAverageFn: func(currentRating string) error {
			return nil
		},
		GetOverallAverageFromDBFn: func() (int, error) {
			return 30, nil
		},
		UpdateAverageFn: func(newAverage int) error {
			return nil
		},
		IsEmptyFn: func(rows *sql.Rows) bool {
			return true
		},
	}
}

// GetAllRows mock.
func (m *MockClient) GetAllRows() (*sql.Rows, error) {
	return m.GetAllRowsFn()
}

// MakeCurrentRatingTheAverage mock.
func (m *MockClient) MakeCurrentRatingTheAverage(currentRating string) error {
	return m.MakeCurrentRatingTheAverageFn(currentRating)
}

// GetOverallAverageFromDB mock.
func (m *MockClient) GetOverallAverageFromDB() (int, error) {
	return m.GetOverallAverageFromDBFn()
}

// UpdateAverage mock.
func (m *MockClient) UpdateAverage(newAverage int) error {
	return m.UpdateAverageFn(newAverage)
}

// IsEmpty mock.
func (m *MockClient) IsEmpty(rows *sql.Rows) bool {
	return m.IsEmptyFn(rows)
}
