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

func TestMyAllocator_UpdateVendor(t *testing.T) {
	type fields struct {
		AuthVendorID       string
		AuthVendorPassword string
		Debug              bool
	}
	type args struct {
		id       string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MyAllocator
	}{
		{
			name: "update id and password",
			fields: fields{
				AuthVendorID:       "test",
				AuthVendorPassword: "1234",
			},
			args: args{
				id:       "appleboy",
				password: "5678",
			},
			want: &MyAllocator{
				AuthVendorID:       "appleboy",
				AuthVendorPassword: "5678",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyAllocator{
				AuthVendorID:       tt.fields.AuthVendorID,
				AuthVendorPassword: tt.fields.AuthVendorPassword,
				Debug:              tt.fields.Debug,
			}
			if got := m.UpdateVendor(tt.args.id, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyAllocator.UpdateVendor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyAllocator_SetDebug(t *testing.T) {
	type fields struct {
		AuthVendorID       string
		AuthVendorPassword string
		Debug              bool
	}
	type args struct {
		enable bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MyAllocator
	}{
		{
			name: "enable debug",
			fields: fields{
				Debug: false,
			},
			args: args{
				enable: true,
			},
			want: &MyAllocator{
				Debug: true,
			},
		},
		{
			name: "disable debug",
			fields: fields{
				Debug: true,
			},
			args: args{
				enable: false,
			},
			want: &MyAllocator{
				Debug: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyAllocator{
				AuthVendorID:       tt.fields.AuthVendorID,
				AuthVendorPassword: tt.fields.AuthVendorPassword,
				Debug:              tt.fields.Debug,
			}
			if got := m.SetDebug(tt.args.enable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyAllocator.SetDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}
