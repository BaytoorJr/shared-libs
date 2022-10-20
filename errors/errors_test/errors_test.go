package errors_test

import (
	"reflect"
	"testing"

	"github.com/BaytoorJr/shared-libs/errors"
)

// Error formatter test
func TestArgError_Error(t *testing.T) {
	type fields struct {
		System           string
		Status           int
		Message          string
		DeveloperMessage string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "formatting",
			fields: fields{
				System:           errors.ErrSystem,
				Status:           400,
				Message:          "wrong request",
				DeveloperMessage: "some field empty",
			},
			want: "400 some field empty",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := &errors.ArgError{
				System:           tc.fields.System,
				Status:           tc.fields.Status,
				Message:          tc.fields.Message,
				DeveloperMessage: tc.fields.DeveloperMessage,
			}
			if got := e.Error(); got != tc.want {
				t.Errorf("Error() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Developer message setter test
func TestArgError_SetDevMessage(t *testing.T) {
	type fields struct {
		System           string
		Status           int
		Message          string
		DeveloperMessage string
	}
	type args struct {
		developMessage string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errors.ArgError
	}{
		{
			name: "dev message",
			fields: fields{
				System:  errors.ErrSystem,
				Status:  400,
				Message: "wrong request",
			},
			args: args{
				"some field empty",
			},
			want: &errors.ArgError{
				System:           errors.ErrSystem,
				Status:           400,
				Message:          "wrong request",
				DeveloperMessage: "some field empty",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := &errors.ArgError{
				System:           tc.fields.System,
				Status:           tc.fields.Status,
				Message:          tc.fields.Message,
				DeveloperMessage: tc.fields.DeveloperMessage,
			}
			if got := e.SetDevMessage(tc.args.developMessage); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SetDevMessage() = %v, want %v", got, tc.want)
			}
		})
	}
}
