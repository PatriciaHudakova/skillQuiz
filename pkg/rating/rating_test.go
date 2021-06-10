package rating

import (
	"database/sql"
	"fmt"
	"skillQuiz/pkg/db"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
)

var testConn *sql.DB

const (
	driverName     = "sqlite3"
	dataSourceName = "identifier.sqlite"
)

func setup(t *testing.T) {
	testConn, _ = sql.Open(driverName, dataSourceName)
	query, err := testConn.Prepare("DROP TABLE averages;")
	query.Exec()
	query2, err := testConn.Prepare("CREATE TABLE averages (uuid INTEGER, overallAverage INTEGER);")
	_, err = query2.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func teardown() {
	_ = testConn.Close()
}

func TestPrintRatings(t *testing.T) { //TODO: fix go mocks
	setup(t)

	//database := &db.Database{Conn: testConn}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRating := NewMockRating(mockCtrl)
	mockDatabase := db.NewMockIDatabase(mockCtrl)

	testData := []string{"yes", "no", "no", "yes", "no"}

	mockRating.EXPECT().CalculateImmediateRating(gomock.Eq(testData)).Return("40")
	mockRating.EXPECT().CalculateAverageRating(mockDatabase, "40")

	PrintRatings(mockDatabase, testData)

	teardown()
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
