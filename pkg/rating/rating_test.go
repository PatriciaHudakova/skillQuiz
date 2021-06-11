package rating

import (
	"database/sql"
	"skillQuiz/pkg/db"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestPrintRatings(t *testing.T) {
	DB := db.NewMockClient()
	testData := []string{"yes", "no", "no", "no", "yes"}

	err := PrintRatings(mockCalculateImmediateRating, mockCalculateAverageRating, DB, testData)
	// Assert no error during function call
	assert.Nil(t, err)

	err = PrintRatings(mockCalculateImmediateRating, mockCalculateAverageRatingErr, DB, testData)
	// Assert error during function call is handled correctly
	assert.NotNil(t, err)
	assert.Equal(t, "something went wrong calculating your average score: test error", err.Error())
}

func TestCalculateImmediateRating(t *testing.T) {
	type args struct {
		params []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly calculate rating using expected input",
			args: args{params: []string{"yes", "yes", "no", "no", "yes"}},
			want: "60",
		},
		{
			name: "should correctly calculate rating using unexpected input",
			args: args{params: []string{"yes", "yes", "", "7927", "';]"}},
			want: "40",
		},
		{
			name: "correct rating should not be limited to 5 parameters",
			args: args{params: []string{"no", "no", "no", "no", "no", "yes", "yes"}},
			want: "29",
		},
		{
			name: "no parameters shouldn't throw NaN as an answer",
			args: args{params: []string{}},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateImmediateRating(tt.args.params); got != tt.want {
				t.Errorf("CalculateImmediateRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateAverageRating(t *testing.T) {
	// Specify desired responses
	mockDB := db.NewMockClient()
	mockDB = &db.MockClient{GetAllRowsFn: func() (*sql.Rows, error) {
		return &sql.Rows{}, nil
	},
		IsEmptyFn: func(rows *sql.Rows) bool {
			return true
		},
		GetOverallAverageFromDBFn: func() (int, error) {
			return 40, nil
		},
		UpdateAverageFn: func(newAverage int) error {
			return nil
		}}

	average, err := CalculateAverageRating(mockDB, "50")

	assert.Nil(t, err)
	assert.NotNil(t, average)
}

func TestNewCalculateAverageRating(t *testing.T) {
	mockDB := db.NewMockClient()
	mockDB = &db.MockClient{GetAllRowsFn: func() (*sql.Rows, error) {
		return &sql.Rows{}, nil
	},
		IsEmptyFn: func(rows *sql.Rows) bool {
			return false
		},
		MakeCurrentRatingTheAverageFn: func(currentRating string) error {
			return nil
		},
	}
	average, err := CalculateAverageRating(mockDB, "50")

	assert.Nil(t, err)
	assert.NotNil(t, average)
}
