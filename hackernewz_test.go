package main

import "testing"

func TestNewClient(t *testing.T) {
	want := NewClient()

	if want.Version != "v0" {
		t.Errorf("want nil but got  %v", want)
	}
}

func TestGetItem(t *testing.T) {
	client := NewClient()

	//Only Id field is require for an item
	want := Item{
		Id: 8863,
	}

	got := client.GetItem(8863)

	if want.Id != got.Id {
		t.Errorf("Error reading story with Id:%v", got.Id)
	}

}

//why is this failing when func is 't' and not T
func TestGetComment(t *testing.T) {
	client := NewClient()

	got, err := client.GetComment(2921983)

	if err != nil {
		t.Errorf("want nil error but got %v", err)
	}

	want := Comment{
		By:     "norvig",
		Id:     2921983,
		Kids:   []int{2922097, 2922429, 2924562, 2922709, 2922573, 2922140, 2922141},
		Parent: 2921506,
		Text:   "Aw shucks, guys ... you make me blush with your compliments.<p>Tell you what, Ill make a deal: I'll keep writing if you keep reading. K?",
		Time:   1314211127,
	}

	if got.Id != want.Id {
		t.Error("want got.Id to 2921983 got", got.Id)
	}

	emptyComment := Comment{}
	if emptyComment.Id == got.Id {
		t.Error("WTF! DANGER WILL ROBINSON!")
	}
}
