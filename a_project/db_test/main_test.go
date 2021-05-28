package main

import "testing"

func TestCamel2Case(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ceshi",
			args: args{
				name: "creationTime",
			},
			want: "creation_time",
		},
		{
			name: "ceshi2",
			args: args{
				name: "name",
			},
			want: "name",
		},
		{
			name: "cesh3",
			args: args{
				name: "databaseName",
			},
			want: "database_name",
		},
		{
			name: "cesh4",
			args: args{
				name: "sysDatabaseID",
			},
			want: "sys_database_id",
		},
		{
			name: "ceshi5",
			args: args{
				name: "upstreamLibrary",
			},
			want: "upstream_library",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel2Case(tt.args.name); got != tt.want {
				t.Errorf("Camel2Case() = %v, want %v", got, tt.want)
			}
		})
	}
}
