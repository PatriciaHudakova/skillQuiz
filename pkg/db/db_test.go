package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "should be able to initialise db and connect to it",
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Assert the received db is not nil
			assert.NotNil(t, db)

			// Close the db connection during cleanup
			t.Cleanup(func() {
				db.Close()
			})
		})
	}
}
