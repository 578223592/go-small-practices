package main

import (
	"reflect"
	"testing"
)

func TestAnalyser_parseQuery(t *testing.T) {
	type fields struct {
		index map[ /*term*/ string][]local
	}
	type args struct {
		keyWords string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]byte
	}{
		{
			name: "test1",
			fields: fields{
				index: nil,
			},
			args: args{
				keyWords: "do AND \"you can\" -tomorrow",
			},
			want: [][]byte{[]byte("do"), []byte("AND"), []byte("\""), []byte("you"), []byte("can"), []byte("\""),
				[]byte("-"), []byte("tomorrow")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Analyser{
				index: tt.fields.index,
			}
			if got := a.parseQuery(tt.args.keyWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
