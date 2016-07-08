# Envy

A basic environment variable file loader library for Go applications.

Simply import it in your program and run it's load command:

```
import "github.com/kyoung/envy"

...

func main() {
  err := envy.Load()
  ...
}

```


### Use
Envy loads files in the format of:

```
FOO=bar
BAZ=qux
```

#### Load
`envy.Load` will look for a .env file from the current context and set them into
the environment of your process.

#### LoadFiles
`envy.LoadFiles` takes a slice of filepaths, and will load each in turn.
