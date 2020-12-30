![Test](https://github.com/paganotoni/fsbox/workflows/Test/badge.svg)

# FSbox

FSbox is a [`packd.Box`](https://github.com/gobuffalo/packd) implementation that uses go 1.16's `io/fs` package. It facilitates the packing of your assets, templates and other things into the Buffalo binary using the new `embed` and `io/fs` libraries and ditching `Packr` or any other third-party tool.

On development mode (GO_ENV != `production`) FSbox falls back to open files using `os.Open` to provide the ability to do changes in those files without the need to recompile the app, facilitating an agile and iterative approach while developing.

⚠️ This package ONLY works with Go 1.16x or higher version of it.
## Usage

You need to have a variable that embeds your templates and public folder. Then use that variable to instantiate two fsbox.

```go
package app

var (
    
    //go:embed templates public
    fsys embed.FS

    // The boxes your app may need.
    AssetsBox  = fsbox.New(fsys, "public")
    TemplatesBox = fsbox.New(fsys, "templates")
)
```

These two boxes will be used in your Buffalo application for plush templates and assets serving as you can see next:

```go
...
    // Adding custom partialFeeder
    helpers["partialFeeder"] = app.TemplatesBox.FindString
    
    // Render engine initialization
    Engine = render.New(render.Options{
		HTMLLayout:   "application.plush.html",
		TemplatesBox: app.TemplatesBox,
		AssetsBox:    app.AssetsBox,
		...
	})
...
```

Defining the partialFeeder is an important step since the default partialFeeder that Buffalo uses adds an underscore prefix to partials, and the embed functionality seems not to support embedding underscore prefixed files.

⚠️ Don't forget to rename your partials without the underscore prefix, otherwise these will not be embedded in the binary.


```go
// Serving assets
bapp.ServeFiles("/", app.AssetsBox)
```
### Can I use this in Buffalo?

Yes, however this could imply some changes in your application layout and webpack configuration.
