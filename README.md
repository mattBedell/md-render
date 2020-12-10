# mdrender
A quick and dirty CLI markdown render utility using [Glamour renderer](https://github.com/charmbracelet/glamour)

### Features

- Renders a markdown file in the terminal
- Mostly works
- No error handling
- First Golang I've ever written

Change the theme with env variable:
```sh
GLAMOUR_STYLE=light
```

### Build

```sh
go build github.com/mattBedell/md_render
```

### See also
- [Glow][3]
- [Github CLI][2]
- [Glamour renderer][1]

[1]: https://github.com/charmbracelet/glamour
[2]: https://github.com/cli/cli
[3]: https://github.com/charmbracelet/glow
