package main

import "testing"

type doc struct {
	id          string
	title       string
	description string
}

func (d doc) Key() string {
	return d.id
}
func Test_sortSpecificKey(t *testing.T) {
	type args struct {
		elems []doc
		key   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"empty elemet -> success",
			args{
				[]doc{},
				"foo",
			},
		},
		{
			"1 elemet -> success",
			args{
				[]doc{{"id1", "title1", "description1"}},
				"id1",
			},
		},
		{
			"3 elemet, key is first -> success",
			args{
				[]doc{
					{"id1", "title1", "description1"},
					{"id2", "title2", "description2"},
					{"id3", "title3", "description3"},
				},
				"id1",
			},
		},
		{
			"3 elemet, key is last -> success",
			args{
				[]doc{
					{"id1", "title1", "description1"},
					{"id2", "title2", "description2"},
					{"id3", "title3", "description3"},
				},
				"id3",
			},
		},
		{
			"3 elemet, key is last -> success",
			args{
				[]doc{
					{"id1", "title1", "description1"},
					{"id2", "title2", "description2"},
					{"id3", "title3", "description3"},
				},
				"id3",
			},
		},
		{
			"key is not exist -> success",
			args{
				[]doc{
					{"id1", "title1", "description1"},
					{"id2", "title2", "description2"},
					{"id3", "title3", "description3"},
				},
				"id5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortSpecificKey(tt.args.elems, tt.args.key)
		})
	}
}
