package event

import "testing"

func TestFeedPanics(t *testing.T)  {
	{
		var f Feed
		f.Send(int(2))
		//want := feedTypeError{}
	}
}