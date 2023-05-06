package model

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHost_IsAlive(t *testing.T) {
	type fields struct {
		HostName     string
		hostip       string
		HostPort     int
		HostUser     string
		HostPassword string
		isAlive      bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test01",
			fields: fields{
				HostName:     "host01",
				hostip:       "192.168.2.230",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
		{
			name: "test02",
			fields: fields{
				HostName:     "host02",
				hostip:       "192.168.2.231",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
		{
			name: "test03",
			fields: fields{
				HostName:     "host03",
				hostip:       "192.168.2.231",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Host{
				HostName:     tt.fields.HostName,
				HostIP:       tt.fields.hostip,
				HostPort:     tt.fields.HostPort,
				HostUser:     tt.fields.HostUser,
				HostPassword: tt.fields.HostPassword,
				isAlive:      tt.fields.isAlive,
			}
			h.IsAlive()
			assert.True(t, h.isAlive)
		})
	}
}

func TestHost_CreateHost(t *testing.T) {
	type fields struct {
		Model        gorm.Model
		HostName     string
		hostip       string
		HostPort     int
		HostUser     string
		HostPassword string
		isAlive      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Host
	}{

		{
			name: "test01",
			fields: fields{
				HostName:     "K8S-node01",
				hostip:       "192.168.2.230",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
		{
			name: "test02",
			fields: fields{
				HostName:     "K8S-node02",
				hostip:       "192.168.2.231",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
		{
			name: "test03",
			fields: fields{
				HostName:     "K8S-master02",
				hostip:       "192.168.2.221",
				HostPort:     22,
				HostUser:     "root",
				HostPassword: "111000",
				isAlive:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Host{
				Model:        tt.fields.Model,
				HostName:     tt.fields.HostName,
				HostIP:       tt.fields.hostip,
				HostPort:     tt.fields.HostPort,
				HostUser:     tt.fields.HostUser,
				HostPassword: tt.fields.HostPassword,
				isAlive:      tt.fields.isAlive,
			}
			// assert.Equalf(t, tt.want, h.CreateHost(), "CreateHost()")
			assert.NotNil(t, h.CreateHost())
		})
	}
}

func TestDeleteHost(t *testing.T) {
	type args struct {
		Id int64
	}
	tests := []struct {
		name string
		args args
		want Host
	}{
		{
			name: "test delete",
			args: args{Id: 6},
			want: Host{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DeleteHost(tt.args.Id), "DeleteHost(%v)", tt.args.Id)
		})
	}
}

func TestHostAssignment(t *testing.T) {
	type args struct {
		userId  int64
		hostIds []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test assign assignment",
			args: args{userId: 100, hostIds: []int64{1, 2, 3}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HostAssignment(tt.args.userId, tt.args.hostIds)
		})
	}
}
