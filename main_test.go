package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_decodeImageData(t *testing.T) {
	type args struct {
		js []byte
	}
	sampleResponse, err := ioutil.ReadFile("sampleResponse.json")
	if err != nil {
		t.Errorf("Error reading sampleResponse.json: %v", err)
		return
	}
	tests := []struct {
		name      string
		args      args
		wantImage image
		wantErr   bool
	}{
		{
			name: "Load a minimal JSON structure",
			args: args{
				js: []byte(`
					{
						"images": [{
							"url": "/a/b/c/def.jpg"
						}]
					}
				`),
			},
			wantImage: image{
				URL: "/a/b/c/def.jpg",
			},
			wantErr: false,
		},
		{
			name: "Load sample JSON file",
			args: args{
				js: sampleResponse,
			},
			wantImage: image{
				URL: "/az/hprichbg/rb/TorontoSkyline_ROW10610765954_1920x1080.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImage, err := decodeImageData(tt.args.js)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeImageData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotImage, tt.wantImage) {
				t.Errorf("decodeImageData() = %v, want %v", gotImage, tt.wantImage)
			}
		})
	}
}
