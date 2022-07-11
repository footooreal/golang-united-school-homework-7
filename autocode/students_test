package coverage

import (
	"os"
	"reflect"
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

func TestPeople_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
	}{
		{
			name: "test1",
			p:    []Person{{"A", "A", time.Now()}, {"A", "A", time.Now()}},
			args: args{0, 1},
		},
		{
			name: "test2",
			p:    []Person{{"A", "A", time.Now()}, {"B", "B ", time.Now()}},
			args: args{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x1 := tt.p[0]
			x2 := tt.p[1]
			tt.p.Swap(tt.args.i, tt.args.j)
			y1 := tt.p[0]
			y2 := tt.p[1]

			if x1 == y1 && x2 == y2 {
				t.Errorf("bad..")
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: "both strings equal and num",
			args: args{"123\n123"},
			want: &Matrix{
				rows: 2,
				cols: 1,
				data: []int{123, 123},
			},
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{"1\n"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "atoi err",
			args:    args{"x"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "atoi err",
			args:    args{"123 123 123\n123"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want [][]int
	}{
		{
			name: "test1",
			m: Matrix{
				rows: 2,
				cols: 1,
				data: []int{123, 123},
			},
			want: [][]int{{123}, {123}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cols(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want [][]int
	}{
		{
			name: "test1",
			m: Matrix{
				rows: 2,
				cols: 1,
				data: []int{123, 123},
			},
			want: [][]int{{123, 123}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Cols(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type args struct {
		row   int
		col   int
		value int
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want bool
	}{
		{
			name: "test1",
			m: &Matrix{
				rows: 2,
				cols: 1,
				data: []int{123, 123},
			},
			args: args{1, 1, 1},
			want: false,
		},
		{
			name: "test2",
			m: &Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2},
			},
			args: args{0, 0, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Set(tt.args.row, tt.args.col, tt.args.value); got != tt.want {
				t.Errorf("Matrix.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
