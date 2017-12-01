Basic scripts in a few languages, so I don't have to waste time on setup. 

## How to Run 

I'm running these on MacOS in iTerm2. I used Homebrew to install the language packages: 

```bash
$ brew install go
$ brew install scriptcs
```

(MacOS ships with Ruby installed.)

Then run from iTerm:

```bash
$ go run hello.go # you can compile first but it's not necessary
$ scriptcs hello.csx
$ ruby hello.rb
```