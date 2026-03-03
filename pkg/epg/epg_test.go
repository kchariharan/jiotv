package epg

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "Initialize EPG with mock server",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO
		})
	}
}

func TestNewProgramme(t *testing.T) {
	type args struct {
		channelID int
		start     string
		stop      string
		title     string
		desc      string
		category  string
		iconSrc   string
	}
	tests := []struct {
		name string
		args args
		want Programme
	}{
		{
			name: "Create new programme",
			args: args{
				channelID: 123,
				start:     "20231225120000 +0530",
				stop:      "20231225130000 +0530",
				title:     "Test Show",
				desc:      "Test Description",
				category:  "Entertainment",
				iconSrc:   "test_icon.jpg",
			},
			want: Programme{
				Channel: "123",
				Start:   "20231225120000 +0530",
				Stop:    "20231225130000 +0530",
				Title: Title{
					Value: "Test Show",
					Lang:  "en",
				},
				Desc: Desc{
					Value: "Test Description",
					Lang:  "en",
				},
				Category: Category{
					Value: "Entertainment",
					Lang:  "en",
				},
				Icon: Icon{
					Src: "https://jiotv.catchup.cdn.jio.com/dare_images/shows/test_icon.jpg",
				},
			},
		},
		{
			name: "Create programme with empty values",
			args: args{
				channelID: 0,
				start:     "",
				stop:      "",
				title:     "",
				desc:      "",
				category:  "",
				iconSrc:   "",
			},
			want: Programme{
				Channel: "0",
				Start:   "",
				Stop:    "",
				Title: Title{
					Value: "",
					Lang:  "en",
				},
				Desc: Desc{
					Value: "",
					Lang:  "en",
				},
				Category: Category{
					Value: "",
					Lang:  "en",
				},
				Icon: Icon{
					Src: "https://jiotv.catchup.cdn.jio.com/dare_images/shows/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProgramme(tt.args.channelID, tt.args.start, tt.args.stop, tt.args.title, tt.args.desc, tt.args.category, tt.args.iconSrc)
			if got.Channel != tt.want.Channel {
				t.Errorf("NewProgramme().Channel = %v, want %v", got.Channel, tt.want.Channel)
			}
			if got.Start != tt.want.Start {
				t.Errorf("NewProgramme().Start = %v, want %v", got.Start, tt.want.Start)
			}
			if got.Title.Value != tt.want.Title.Value {
				t.Errorf("NewProgramme().Title.Value = %v, want %v", got.Title.Value, tt.want.Title.Value)
			}
			if got.Icon.Src != tt.want.Icon.Src {
				t.Errorf("NewProgramme().Icon.Src = %v, want %v", got.Icon.Src, tt.want.Icon.Src)
			}
		})
	}
}

func TestGenXML(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Generate XML with mock server",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO
		})
	}
}

func TestFormatTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Format specific time",
			args: args{t: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)},
			want: "20231225153045 +0000",
		},
		{
			name: "Format with different timezone",
			args: args{t: time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("EST", -5*3600))},
			want: "20230101000000 -0500",
		},
		{
			name: "Format with positive timezone",
			args: args{t: time.Date(2023, 6, 15, 12, 0, 0, 0, time.FixedZone("IST", 5*3600+30*60))},
			want: "20230615120000 +0530",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTime(tt.args.t); got != tt.want {
				t.Errorf("formatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenXMLGz(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Generate gzipped XML with mock server",
			filename: "/tmp/test_epg.xml.gz",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO
		})
	}
}

func TestJSONInt64_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JSONInt64
		wantErr bool
	}{
		{
			name:    "Unmarshal from integer",
			args:    args{data: []byte("1609459200123")},
			want:    JSONInt64(1609459200123),
			wantErr: false,
		},
		{
			name:    "Unmarshal from string",
			args:    args{data: []byte(`"1609459200123"`)},
			want:    JSONInt64(1609459200123),
			wantErr: false,
		},
		{
			name:    "Unmarshal invalid string",
			args:    args{data: []byte(`"abc"`)},
			want:    JSONInt64(0),
			wantErr: true,
		},
		{
			name:    "Unmarshal invalid JSON",
			args:    args{data: []byte("invalid json")},
			want:    JSONInt64(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i JSONInt64
			err := i.UnmarshalJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONInt64.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && i != tt.want {
				t.Errorf("JSONInt64.UnmarshalJSON() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestJSONInt_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JSONInt
		wantErr bool
	}{
		{
			name:    "Unmarshal from integer",
			args:    args{data: []byte("144")},
			want:    JSONInt(144),
			wantErr: false,
		},
		{
			name:    "Unmarshal from string",
			args:    args{data: []byte(`"144"`)},
			want:    JSONInt(144),
			wantErr: false,
		},
		{
			name:    "Unmarshal invalid string",
			args:    args{data: []byte(`"abc"`)},
			want:    JSONInt(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i JSONInt
			err := i.UnmarshalJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONInt.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && i != tt.want {
				t.Errorf("JSONInt.UnmarshalJSON() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestJSONString_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JSONString
		wantErr bool
	}{
		{
			name:    "Unmarshal from string",
			args:    args{data: []byte(`"test"`)},
			want:    JSONString("test"),
			wantErr: false,
		},
		{
			name:    "Unmarshal from array",
			args:    args{data: []byte(`["test"]`)},
			want:    JSONString("test"),
			wantErr: false,
		},
		{
			name:    "Unmarshal from empty array",
			args:    args{data: []byte(`[]`)},
			want:    JSONString(""),
			wantErr: false,
		},
		{
			name:    "Unmarshal null",
			args:    args{data: []byte(`null`)},
			want:    JSONString(""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s JSONString
			err := s.UnmarshalJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONString.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && s != tt.want {
				t.Errorf("JSONString.UnmarshalJSON() = %v, want %v", s, tt.want)
			}
		})
	}
}
