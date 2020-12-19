# FSbox

FSbox is a [`packd.Box`](https://github.com/gobuffalo/packd) implementation that uses go 1.16's `io/fs` package.

### Important
- ⚠️ This package only works with Go 1.16 or higher version of it.
- ⚠️ This package still a WIP, use with caution.
## Usage

In your root folder create `files.go` with the following content:

```go
package myapp

var (
    //go:embed templates public
    fsys embed.FS

    // The boxes your app may need.
    AssetsBox  = fsbox.New(fsys, "public")
    TemplatesBox = fsbox.New(fsys, "templates")
)
```

And then use it in your `actions/render` use these when creating your Render engine. (Or anywhere you want it.)

```go
//p.e render.go
var Engine = render.New(render.Options{
	HTMLLayout:   "application.html",
	TemplatesBox: app.TemplatesBox,
	AssetsBox:    app.AssetsBox,
	Helpers: map[string]interface{}{
        ...
```

The reason to do it at the top level is because go:embed uses relative paths and [seems not to support .. expressions](https://go.googlesource.com/proposal/+/master/design/draft-embed.md#go_embed-directives).

### Can I use this in Buffalo?

Not yet. This still an experimental package and will need to solve some issues before becoming stable:

- [ ] Partials

Syntax for partials requires it to start with underscode, however go:embed ignores files that start with underscore.
