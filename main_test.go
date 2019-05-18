package main

import (
	"testing"
)

func TestMakeNameBackup(t *testing.T) {
	got := makeName("heavensVaultSave_Backup.json", "2019-05-18-13_53_18")
	want := "heavensVaultSave_Backup.2019-05-18-13_53_18.json"
	if got != want {
		t.Errorf("got '%s', wanted '%s'", got, want)
	}
}

func TestMakeNameRegular(t *testing.T) {
	got := makeName("heavensVaultSave.json", "2019-05-18-13_53_18")
	want := "heavensVaultSave.2019-05-18-13_53_18.json"
	if got != want {
		t.Errorf("got '%s', wanted '%s'", got, want)
	}
}

func Test_isSave(t *testing.T) {
	type args struct {
		fileInfoName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isRegular", args{fileInfoName: "heavensVaultSave.json"}, true},
		{"isSomethingElse", args{fileInfoName: "output_log.txt"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSave(tt.args.fileInfoName); got != tt.want {
				t.Errorf("isSave() = %v, want %v", got, tt.want)
			}
		})
	}
}
