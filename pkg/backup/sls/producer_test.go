package sls

import (
	"context"
	"testing"

	"github.com/kubevela/workflow/api/v1alpha1"
)

func TestHandler_Store(t *testing.T) {
	type fields struct {
		LogStoreName    string
		ProjectName     string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
	}
	type args struct {
		ctx context.Context
		run *v1alpha1.WorkflowRun
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	} {
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Handler{
				LogStoreName:    tt.fields.LogStoreName,
				ProjectName:     tt.fields.ProjectName,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
			}
			if err := s.Store(tt.args.ctx, tt.args.run); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
