package helpers

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args[T any, S any] struct {
		arr       []T
		action    func(S, T, int) S
		initValue S
	}

	type TestItem[T any, S any] struct {
		name string
		args args[T, S]
		want S
	}
	tests := []TestItem[any, any]{
		{
			name: "numberTest",
			args: args[any, any]{
				arr:       []any{1, 2, 3},
				action:    func(i1, i2 any, _ int) any { return i1.(int) + i2.(int) },
				initValue: 0,
			},
			want: 6,
		},
		{
			name: "stringTest",
			args: args[any, any]{
				arr:       []any{"1", "2", "3", "4", "5", "6"},
				action:    func(i1, i2 any, _ int) any { return i1.(string) + i2.(string) },
				initValue: "",
			},
			want: "123456",
		},
		{
			name: "mapTest",
			args: args[any, any]{
				arr: []any{"1", "2", "3"},
				action: func(i1, i2 any, i3 int) any {
					i1.(map[int]string)[i3] = i2.(string)
					return i1
				},
				initValue: map[int]string{},
			},
			want: map[int]string{0: "1", 1: "2", 2: "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.arr, tt.args.action, tt.args.initValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, S any] struct {
		arr    []T
		action func(T, int) S
	}

	type TestItem[T any, S any] struct {
		name string
		args args[T, S]
		want S
	}
	tests := []TestItem[any, any]{
		{
			name: "powerTest",
			args: args[any, any]{
				arr: []any{1, 2, 3, 4, 5, 6},
				action: func(i1 any, _ int) any {
					_i1 := i1.(int)
					return _i1 * _i1
				},
			},
			want: []any{1, 4, 9, 16, 25, 36},
		},
		{
			name: "mapTest",
			args: args[any, any]{
				arr: []any{1, 2, 3},
				action: func(i1 any, index int) any {
					_i1 := i1.(int)
					return map[string]int{
						"index": index,
						"value": _i1,
					}
				},
			},

			want: []any{
				map[string]int{"index": 0, "value": 1},
				map[string]int{"index": 1, "value": 2},
				map[string]int{"index": 2, "value": 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.arr, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
