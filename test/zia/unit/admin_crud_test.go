package unit

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

// import (
// 	"testing"

// 	"github.com/jarcoal/httpmock"
// )

// var expected_raw_page string = "{\"object\":\"page\"," +
// 	"\"id\":\"xxxxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"," +
// 	"\"created_time\":\"2022-01-30T23:51:00.000Z\"," +
// 	"\"last_edited_time\":\"2022-02-03T23:42:00.000Z\",\"cover\":null,\"icon\":null," +
// 	"\"parent\":{\"type\":\"database_id\",\"database_id\":\"xxxxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"}," +
// 	"\"archived\":false,\"properties\":{\"Status\":" +
// 	"{\"id\":\"bxnI\",\"type\":\"select\",\"select\":null}," +
// 	"\"Category\":{\"id\":\"mAOk\",\"type\":\"select\",\"select\":null}," +
// 	"\"Name\":{\"id\":\"title\",\"type\":\"title\"," +
// 	"\"title\":[{\"type\":\"text\",\"text\":{\"content\":\"Test Page\"," +
// 	"\"link\":null},\"annotations\":{\"bold\":false,\"italic\":false,\"strikethrough\":false," +
// 	"\"underline\":false,\"code\":false,\"color\":\"default\"}," +
// 	"\"plain_text\":\"Test Page\",\"href\":null}]}},\"url\":\"https://www.notion.so/xxx\"}"

// func TestFetchAdminUsers(t *testing.T) {
// 	// Mock Notion API response
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
// 	httpmock.RegisterResponder(
// 		"GET",
// 		utils.BaseUrl+"pages/111",
// 		httpmock.NewBytesResponder(
// 			200,
// 			[]byte(expected_raw_page)),
// 	)

// 	page := notion_api.FetchPageByPageId("", "111")

// 	// To test a key in object page
// 	if len(page.Id) == 0 {
// 		t.Errorf("Can't parse a property `id`.")
// 	}
// 	if len(page.CreatedTime) == 0 {
// 		t.Errorf("Can't parse a property `created_time`.")
// 	}
// 	if len(page.CreatedTime) == 0 {
// 		t.Errorf("Can't parse a property `created_time`.")
// 	}
// 	if len(page.LastEditedTime) == 0 {
// 		t.Errorf("Can't parse a property `last_edited_time`.")
// 	}
// }
