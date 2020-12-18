# mdrender
A quick and dirty CLI markdown render utility using [Glamour renderer][1]

### Features

- Renders a markdown file in the terminal
- Uses the same markdown renderer as the [Github CLI][2]
- Mostly works
- No error handling
- First Golang I've ever written

### Install

```sh
brew install mattbedell/formulae/mdrender
```

### Build

```sh
go build github.com/mattBedell/md_render
```

### Usage

```sh
mdrender path/to/file.md
```

### Configuration

Change the theme with env variable:
```sh
GLAMOUR_STYLE=light
```

### See also
- [mdcat][4] - A CLI markdown renderer
- [Glow][3]
- [Github CLI][2]
- [Glamour renderer][1]

[1]: https://github.com/charmbracelet/glamour
[2]: https://github.com/cli/cli
[3]: https://github.com/charmbracelet/glow
[4]: https://github.com/lunaryorn/mdcat
