package main

import (
	"testing"
)

func Test_getVideoAspectRatio(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filePath string
		want     string
		wantErr  bool
	}{
		{
			filePath: "./samples/boots-video-horizontal.mp4",
			want:     "16:9",
			wantErr:  false,
		},
		{
			filePath: "./samples/boots-video-vertical.mp4",
			want:     "9:16",
			wantErr:  false,
		},
		{
			filePath: "",
			want:     "",
			wantErr:  true,
		},
		{
			filePath: "./samples/is-bootdev-for-you.pdf",
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := getVideoAspectRatio(tt.filePath)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getVideoAspectRatio() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getVideoAspectRatio() succeeded unexpectedly")
			}

			if tt.want != got {
				t.Errorf("getVideoAspectRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
