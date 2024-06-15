# Marker Replacer
Have you ever had to dynamically change something in a text configuration, but you wanted to keep old content
and just change it into a really specific place? No? Neither do I. But WordPress, long time ago, used to use this
technique to replace 

Actually I just wanted to do something like this today and I failed miserably to bash. Sed wasn't working even after
trying several ways, so I got angry and wrote this.

Basically it reads a file, looks for a marker like the one below (please se raw version as Markdown doesn't render it)
and then you can give any arbitrary text, and it should replace for you.

Just [install](#install) it (or download the binaries, I spend some time building multi-arch binaries for you 🥲) and run it.

## Install
Easiest way, `go install github.com/williamokano/marker-replacer/cmd/cli && mv ${GOBIN}/cli ${GOBIN}/marker-replacer`

## Run it
Example `marker-replacer -file README.md -marker foobar "awesome new text here"`

Then you can redirect it to a file with `... > README.md` or concatenate with whatever you like it.

## Example
Here's an example

### Original
```markdown
### Original
Text that I should keep
<!--foobar-->
v1.0.0
<!--/foobar-->
This part shouldn't be messed around too.
```

### Replaced
```markdown
### Original
Text that I should keep
<!--foobar-->
awesome new text here
<!--/foobar-->
This part shouldn't be messed around too.
```