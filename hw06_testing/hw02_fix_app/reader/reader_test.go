package reader

import (
	"testing"

	"github.com/napryag/hw_otus_basic/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func Test_ReadJSON(t *testing.T) {
	testCases := []struct {
		name     string
		filepath string
		got      []types.Employee
		want     []types.Employee
	}{
		{
			name:     "data1",
			filepath: "data_test1.json",
			want: []types.Employee{
				{UserID: 14, Age: 43, Name: "Sasha", DepartmentID: 2854},
				{UserID: 74, Age: 19, Name: "Oleg", DepartmentID: 7429},
			},
		},
		{
			name:     "data2",
			filepath: "data_test2.json",
			want: []types.Employee{
				{UserID: 255, Age: 25, Name: "Dima", DepartmentID: 122},
				{UserID: 256, Age: 25, Name: "Nikita", DepartmentID: 123},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC.got, _ = ReadJSON(tC.filepath, -1)
			assert.Equal(t, tC.want, tC.got)
		})
	}
}
