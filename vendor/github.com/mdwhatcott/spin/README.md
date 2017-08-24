# spin
--
    import "github.com/mdwhatcott/spin"

Package spin implements a simple console spinner. See the tests and examples for
examples.

## Usage

```go
var (
	// StyleLine is a simple example of the kinds of styles you could pass into the New... functions.
	StyleLine = []string{"|", "/", "-", "\\"}

	// StyleSteps is another example of the kinds of styles you could pass into the New... functions.
	StyleSteps = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}

	// StyleShutter is another example of the kinds of styles you could pass into the New... functions.
	StyleShutter = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
)
```

#### func  GoStart

```go
func GoStart()
```
GoStart forwards to a package-level *Spinner (for convenience).

#### func  Start

```go
func Start()
```
Start forwards to a package-level *Spinner (for convenience).

#### func  Stop

```go
func Stop()
```
Stop forwards to a package-level *Spinner (for convenience).

#### type Spinner

```go
type Spinner struct {
}
```

Spinner prints a repeating pattern to os.Stdout by printing a sequence of
characters interspersed with carriage returns. A Spinner is controlled by the
provided methods: Start, GoStart (like calling `go Start()`, and Stop.

#### func  New

```go
func New(style []string, delay time.Duration) *Spinner
```
New creates a spinner which you can start and stop.

#### func  NewWithPadding

```go
func NewWithPadding(style []string, delay time.Duration, prefix, suffix string) *Spinner
```
NewWithPadding creates a spinner which you can start and stop. Allows a prefix
and suffix to be printed along with the specified style.

#### func (*Spinner) GoStart

```go
func (self *Spinner) GoStart()
```
GoStart begins the spinner on a fresh goroutine.

#### func (*Spinner) Start

```go
func (self *Spinner) Start()
```
Start begins the spinner on the current goroutine (hopefully you've got another
goroutine that can call Stop...).

#### func (*Spinner) Stop

```go
func (self *Spinner) Stop()
```
Stop sends a signal to stop the spinner.
