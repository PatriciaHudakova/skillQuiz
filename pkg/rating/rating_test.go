package rating

import (
	"database/sql"
	"skillQuiz/pkg"
	"skillQuiz/pkg/db"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func createQuestionSet() []pkg.Question {
	return []pkg.Question{{
		Text:   "test Question 1",
		Answer: "yes",
		Value:  1,
	}, {
		Text:   "test Question 2",
		Answer: "no",
		Value:  0,
	}, {
		Text:   "test Question 3",
		Answer: "yes",
		Value:  1,
	}, {
		Text:   "test Question 4",
		Answer: "yes",
		Value:  1,
	}, {
		Text:   "test Question 5",
		Answer: "no",
		Value:  0,
	},
	}
}

func createQuestionSetInvalid() []pkg.Question {
	return []pkg.Question{{
		Text:   "test Question 1",
		Answer: "invalid response",
		Value:  0,
	}, {
		Text:   "test Question 2",
		Answer: "random string",
		Value:  0,
	}, {
		Text:   "test Question 3",
		Answer: "no",
		Value:  0,
	}, {
		Text:   "test Question 4",
		Answer: "yes",
		Value:  1,
	}, {
		Text:   "test Question 5",
		Answer: "another random string",
		Value:  0,
	},
	}
}

func TestPrintRatings(t *testing.T) {
	DB := db.NewMockClient()
	testData := createQuestionSet()

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
		params []pkg.Question
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly calculate rating using expected input",
			args: args{params: createQuestionSet()},
			want: "60",
		},
		{
			name: "should correctly calculate rating using unexpected input",
			args: args{params: createQuestionSetInvalid()},
			want: "20",
		},
		{
			name: "no parameters shouldn't throw NaN as an answer",
			args: args{params: []pkg.Question{}},
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
	type args struct {
		db            db.IDatabase
		currentRating string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should use current run's average as new",
			args: args{
				db: &db.MockClient{GetAllRowsFn: func() (*sql.Rows, error) {
					return &sql.Rows{}, nil
				},
					IsEmptyFn: func(rows *sql.Rows) bool {
						return true
					},
					MakeCurrentRatingTheAverageFn: func(currentRating string) error {
						return nil
					},
				},
				currentRating: "50",
			},
			wantErr: false,
		},
		{
			name: "should recognise existing average and use it to persist new",
			args: args{
				db: &db.MockClient{GetAllRowsFn: func() (*sql.Rows, error) {
					return &sql.Rows{}, nil
				},
					IsEmptyFn: func(rows *sql.Rows) bool {
						return false
					},
					GetOverallAverageFromDBFn: func() (int, error) {
						return 40, nil
					},
					UpdateAverageFn: func(newAverage int) error {
						return nil
					}},
				currentRating: "50",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateAverageRating(tt.args.db, tt.args.currentRating)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAverageRating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
