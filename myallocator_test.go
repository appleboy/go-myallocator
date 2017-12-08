package ma

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		vendorID       string
		vendorPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    *MyAllocator
		wantErr bool
	}{
		{
			name: "missing password",
			args: args{
				vendorID: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing username",
			args: args{
				vendorPassword: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.vendorID, tt.args.vendorPassword)
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
