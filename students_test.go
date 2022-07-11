package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople_Len(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{
			name: "test1",
			p:    []Person{{"A", "Asurname", time.Now()}},
			want: 1,
		},
		{
			name: "test2",
			p:    []Person{{"B", "Bsurname", time.Now()}, {"John", "Johnson", time.Now().Add(-time.Hour * 1)}},
			want: 2,
		},
		{
			name: "test3",
			p:    []Person{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("People.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{
			name: "test1",
			p:    []Person{{"B", "Bsurname", time.Now()}, {"John", "Johnson", time.Now().Add(-time.Hour * 1)}},
			args: args{0, 1},
			want: true,
		},

		{
			name: "test2",
			p:    []Person{{"A", "B", time.Now()}, {"B", "A", time.Now()}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "test3",
			p:    []Person{{"A", "B", time.Now()}, {"B", "A", time.Now()}},
			args: args{1, 0},
			want: false,
		},

		{
			name: "test4",
			p:    []Person{{"A", "A", time.Now()}, {"A", "B", time.Now()}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "test5",
			p:    []Person{{"A", "A", time.Now()}, {"A", "B", time.Now()}},
			args: args{1, 0},
			want: false,
		},
		{
			name: "test6",
			p:    []Person{{"A", "A", time.Now()}, {"A", "A", time.Now()}},
			args: args{0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("People.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}