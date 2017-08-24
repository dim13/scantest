package spin

import (
	"bytes"
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSpinFixture(t *testing.T) {
	gunit.Run(new(SpinFixture), t)
}

type SpinFixture struct {
	*gunit.Fixture

	absorber *carriageReturnAbsorber
	output   *output
}

func (this *SpinFixture) Setup() {
	this.absorber = &carriageReturnAbsorber{Buffer: bytes.NewBufferString("")}
	this.output = &output{out: this.absorber}
}

func (this *SpinFixture) spin(spinner *Spinner) {
	spinner.out = this.output
	go spinner.Start()
	time.Sleep(time.Millisecond)
	spinner.Stop()
}

func (this *SpinFixture) TestSpinner() {
	this.spin(New(StyleLine, time.Nanosecond))

	expected := StyleLine

	this.Println("- The provided pattern should be written repeatedly to the output.")
	this.So(this.absorber.String(), should.StartWith, expected+expected)

	this.Println("- Each write should start with a carriage return.")
	this.So(this.absorber.CarriageReturns, should.Equal, this.absorber.TotalWrites)
}

func (this *SpinFixture) TestSpinner_PrefixSuffix() {
	this.spin(NewWithPadding(StyleLine, time.Nanosecond, ">> ", " <<"))

	expected := ">> | <<" +
		">> / <<" +
		">> - <<" +
		">> \\ <<"
	this.Println("- The prefix and suffix should surround each write.")
	this.So(this.absorber.String(), should.StartWith, expected+expected)
}

///////////////////////////////////////////////////////////////////////////////

// carriageReturnAbsorber is an io.Writer that stores an internal buffer
// allowing inspection of what was written. It substitutes for os.Stdout
// but throws out leading carriage returns to make the tests easier to read.
type carriageReturnAbsorber struct {
	*bytes.Buffer
	CarriageReturns int
	TotalWrites     int
}

func (self *carriageReturnAbsorber) Write(value []byte) (int, error) {
	self.TotalWrites++

	if value[0] == '\r' && len(value) > 1 {
		self.CarriageReturns++
		return self.Buffer.Write(value[1:])
	} else if value[0] == '\r' {
		self.CarriageReturns++
		return 0, nil
	} else {
		return self.Buffer.Write(value)
	}
}

////////////////////////////////////////////////////////////////////////////
