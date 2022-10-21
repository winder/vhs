# VHS

<p>
  <img src="https://user-images.githubusercontent.com/42545625/197049037-b38fea25-a885-4945-825e-d29842c5e44b.png#gh-dark-mode-only" width="500" />
  <img src="https://user-images.githubusercontent.com/42545625/197049039-83498ce6-d01d-4a08-8794-64770606ca8e.png#gh-light-mode-only" width="500" />
  <br>
  <a href="https://github.com/charmbracelet/vhs/releases"><img src="https://img.shields.io/github/release/charmbracelet/vhs.svg" alt="Latest Release"></a>
  <a href="https://pkg.go.dev/github.com/charmbracelet/vhs?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="Go Docs"></a>
  <a href="https://github.com/charmbracelet/vhs/actions"><img src="https://github.com/charmbracelet/vhs/workflows/build/badge.svg" alt="Build Status"></a>
</p>

Write terminal GIFs as code for integration testing and demoing your CLI tools.

<img alt="Display of welcome.gif" src="./examples/welcome.gif" width="600" />

The above example is generated from VHS ([View Tape](./examples/welcome.tape)).

## Tutorial

To get started, [install VHS](#installation) and create a new `.tape` file.

```sh
vhs new demo.tape
```

Open the `.tape` file with your favorite `$EDITOR`.

```sh
vim demo.tape
```

In the file, write [commands](#commands) to perform on the terminal.
View [Documentation](#commands) for a list of all the possible commands.

```elixir
# Render the output GIF to demo.gif
Output demo.gif

# Set up a 1200x600 terminal with 46px font size.
Set FontSize 46
Set Width 1200
Set Height 600

# Type a command in the terminal.
Type "echo 'Welcome to VHS!'"

# Pause for dramatic effect...
Sleep 500ms

# Run the command by pressing enter.
Enter

# Admire the output for a bit.
Sleep 5s
```

Once you've written the commands to perform, save and exit the file. And, run
the VHS tool on the file.

```sh
vhs < demo.tape
```

All done! You should see a new file called `demo.gif` (or whatever you named
the `Output`) in the directory.

More examples are in the [`examples/`](https://github.com/charmbracelet/vhs/tree/main/examples) folder.

<img alt="Display of demo.gif" src="./examples/demo.gif" width="600" />

## Installation

> **Note**
> VHS requires [`ttyd`](https://github.com/tsl0922/ttyd) and [`ffmpeg`](https://ffmpeg.org) to be installed.

Use a package manager:

```sh
# macOS or Linux
brew install vhs ttyd ffmpeg

# Arch Linux (btw)
yay -S vhs ttyd ffmpeg

# Nix
nix-env -iA nixpkgs.vhs nixpkgs.ttyd nixpkgs.ffmpeg
```

Or, use docker:

```sh
docker run ghcr.io/charmbracelet/vhs <cassette>.tape
```

Or, download it:

* [Packages][releases] are available in Debian and RPM formats
* [Binaries][releases] are available for Linux, macOS, and Windows

Or, just install it with `go`:

```sh
go install github.com/charmbracelet/vhs@latest
```

[releases]: https://github.com/charmbracelet/vhs/releases

## Teams

If you are a team using VHS, reach out to [vt100@charm.sh](mailto:vt100@charm.sh)
to set up a VHS rendering server with an `ssh` interface to avoid any local
setup for your team.

```sh
ssh vhs.charm.sh < demo.tape > demo.gif
```

## Commands

For documentation on the command line, run:

```sh
vhs manual
```

* [`Output <path>`](#output)
* [`Set <Setting> Value`](#settings)
* [`Type "<characters>"`](#type)
* [`Sleep <time>`](#sleep)
* [`Hide`](#hide)
* [`Show`](#show)

### Keys

Key commands take an optional `@time` and repeat `count`.
For example, the following presses the `Left` key 5 times with a 500
millisecond delay between each keystroke.

```elixir
Left@500ms 5
```

* [`Backspace`](#backspace)
* [`Ctrl`](#ctrl)
* [`Down`](#down)
* [`Enter`](#enter)
* [`Space`](#space)
* [`Tab`](#tab)
* [`Left`](#arrow-keys)
* [`Right`](#arrow-keys)
* [`Up`](#arrow-keys)
* [`Down`](#arrow-keys)

### Settings

The `Set` command allows you to change aspects of the terminal, such as the
font settings, window dimensions, and output GIF location.

Setting commands must be set at the top of the tape file. Any setting (except
`TypingSpeed`) command that is set after a non-setting or non-output command
will be ignored.

* [`Set FontSize <Number>`](#set-font-size)
* [`Set FontFamily <String>`](#set-font-family)
* [`Set Height <Number>`](#set-height)
* [`Set Width <Number>`](#set-width)
* [`Set LetterSpacing <Float>`](#set-letter-spacing)
* [`Set LineHeight <Float>`](#set-line-height)
* [`Set TypingSpeed <Time>`](#set-typing-speed)
* [`Set Theme <String>`](#set-theme)
* [`Set Padding <Number>[em|px]`](#set-padding)
* [`Set Framerate <Number>`](#set-framerate)

### Sleep

The `Sleep` command allows you to continue capturing frames without interacting
with the terminal. This is useful when you need to wait on something to
complete while including it in the recording like a spinner or loading state.
The command takes a number argument in seconds.

```elixir
Sleep 0.5   # 500ms
Sleep 2     # 2s
Sleep 100ms # 100ms
Sleep 1s    # 1s
```

### Type

The `Type` command allows you to type in the terminal and emulate key presses.
This is useful for typing commands or interacting with the terminal.
The command takes a string argument with the characters to type.

```elixir
Type "Whatever you want"
```

<img alt="Type" src="./examples/commands/type.gif" width="600" />

### Output

The `Output` command allows you to specify the location and file format
of the render. You can specify more than one output in a tape file which
will render them to the respective locations.

```elixir
Output out.gif
Output out.mp4
Output out.webm
Output frames/ # .png frames
```

### Keys

#### Backspace

Press the backspace key with the `Backspace` command.

```elixir
Backspace 18
```

<img alt="Press Backspace" src="./examples/commands/backspace.gif" width="600" />

#### Ctrl

Press a control sequence with the `Ctrl` command.

```elixir
Ctrl+R
```

<img alt="" src="./examples/commands/ctrl.gif" width="600" />

#### Enter

Press the enter key with the `Enter` command.

```elixir
Enter 2
```

<img alt="" src="./examples/commands/enter.gif" width="600" />

#### Arrow Keys

Press any of the arrow keys with the `Up`, `Down`, `Left`, `Right` commands.

```elixir
Up 2
Down 3
Left 10
Right 10
```

<img alt="" src="./examples/commands/arrow.gif" width="600" />

#### Tab

Press the tab key with the `Tab` command.

```elixir
Tab@500ms 2
```

<img alt="" src="./examples/commands/tab.gif" width="600" />

#### Space

Press the space bar with the `Space` command.

```elixir
Space 10
```

<img alt="" src="./examples/commands/space.gif" width="600" />

### Settings

#### Set Font Size

Set the font size with the `Set FontSize <Number>` command.

```elixir
Set FontSize 10
Set FontSize 20
Set FontSize 30
Set FontSize 40
```

<img alt="" src="./examples/settings/font-size-10.gif" width="600" />

<img alt="" src="./examples/settings/font-size-20.gif" width="600" />

<img alt="" src="./examples/settings/font-size-40.gif" width="600" />

#### Set Font Family

Set the font family with the `Set FontFamily "<Font>"` command

```elixir
Set FontFamily "Fira Code"
Set FontFamily "Menlo"
Set FontFamily "Monaco"
Set FontFamily "Monoflow"
Set FontFamily "DejaVu Sans Mono"
```

<img alt="" src="./examples/settings/font-family.gif" width="600" />

#### Set Height

Set the height of the terminal with the `Set Height` command.

```elixir
Set Height 600
Set Height 1000
```

#### Set Width

Set the width of the terminal with the `Set Width` command.

```elixir
Set Width 1200
Set Width 2000
```

#### Set Letter Spacing

Set the spacing between letters (tracking) with the `Set LetterSpacing`
Command.

```elixir
Set LetterSpacing 1.2
Set LetterSpacing 2.4
Set LetterSpacing 3.6
```

<img alt="" src="./examples/settings/letter-spacing.gif" width="600" />

#### Set Line Height

```elixir
Set LineHeight 1.4
Set LineHeight 1.6
Set LineHeight 1.8
```

Set the spacing between lines with the `Set LineHeight` Command.

<img alt="" src="./examples/settings/line-height.gif" width="600" />

#### Set Typing Speed

```elixir
Set TypingSpeed 500ms # 500ms
Set TypingSpeed 1s    # 1s
```

Set the typing speed of seconds per key press. For example, a typing speed of
`0.1` would result in a `0.1s` (`100ms`) delay between each character being typed.

This setting can also be overridden per command with the `@<time>` syntax.

```elixir
Set TypingSpeed 0.1
Type "100ms delay per character"
Type@500ms "500ms delay per character"
```

<img alt="" src="./examples/settings/typing-speed.gif" width="600" />

#### Set Theme

Set the theme of the terminal with the `Set Theme` command. The theme value
should be a JSON string with the base 16 colors and foreground + background.

```elixir
Set Theme { "name": "Whimsy", "black": "#535178", "red": "#ef6487", "green": "#5eca89", "yellow": "#fdd877", "blue": "#65aef7", "purple": "#aa7ff0", "cyan": "#43c1be", "white": "#ffffff", "brightBlack": "#535178", "brightRed": "#ef6487", "brightGreen": "#5eca89", "brightYellow": "#fdd877", "brightBlue": "#65aef7", "brightPurple": "#aa7ff0", "brightCyan": "#43c1be", "brightWhite": "#ffffff", "background": "#29283b", "foreground": "#b3b0d6", "selectionBackground": "#3d3c58", "cursorColor": "#b3b0d6" }
Set Theme { "name": "wilmersdorf", "black": "#34373e", "red": "#e06383", "green": "#7ebebd", "yellow": "#cccccc", "blue": "#a6c1e0", "purple": "#e1c1ee", "cyan": "#5b94ab", "white": "#ababab", "brightBlack": "#434750", "brightRed": "#fa7193", "brightGreen": "#8fd7d6", "brightYellow": "#d1dfff", "brightBlue": "#b2cff0", "brightPurple": "#efccfd", "brightCyan": "#69abc5", "brightWhite": "#d3d3d3", "background": "#282b33", "foreground": "#c6c6c6", "selectionBackground": "#1f2024", "cursorColor": "#7ebebd" }
Set Theme { "name": "Wombat", "black": "#000000", "red": "#ff615a", "green": "#b1e969", "yellow": "#ebd99c", "blue": "#5da9f6", "purple": "#e86aff", "cyan": "#82fff7", "white": "#dedacf", "brightBlack": "#313131", "brightRed": "#f58c80", "brightGreen": "#ddf88f", "brightYellow": "#eee5b2", "brightBlue": "#a5c7ff", "brightPurple": "#ddaaff", "brightCyan": "#b7fff9", "brightWhite": "#ffffff", "background": "#171717", "foreground": "#dedacf", "selectionBackground": "#453b39", "cursorColor": "#171717" }
Set Theme { "name": "Wryan", "black": "#333333", "red": "#8c4665", "green": "#287373", "yellow": "#7c7c99", "blue": "#395573", "purple": "#5e468c", "cyan": "#31658c", "white": "#899ca1", "brightBlack": "#3d3d3d", "brightRed": "#bf4d80", "brightGreen": "#53a6a6", "brightYellow": "#9e9ecb", "brightBlue": "#477ab3", "brightPurple": "#7e62b3", "brightCyan": "#6096bf", "brightWhite": "#c0c0c0", "background": "#101010", "foreground": "#999993", "selectionBackground": "#4d4d4d", "cursorColor": "#9e9ecb" }
Set Theme { "name": "Abernathy", "black": "#000000", "red": "#cd0000", "green": "#00cd00", "yellow": "#cdcd00", "blue": "#1093f5", "purple": "#cd00cd", "cyan": "#00cdcd", "white": "#faebd7", "brightBlack": "#404040", "brightRed": "#ff0000", "brightGreen": "#00ff00", "brightYellow": "#ffff00", "brightBlue": "#11b5f6", "brightPurple": "#ff00ff", "brightCyan": "#00ffff", "brightWhite": "#ffffff", "background": "#111416", "foreground": "#eeeeec", "selectionBackground": "#eeeeec", "cursorColor": "#bbbbbb" }
```

<img alt="" src="./examples/settings/theme.gif" width="600" />


#### Set Padding

Set the padding of the terminal frame with the `Set Padding` command.

```elixir
Set Padding 1em
Set Padding 2em
Set Padding 3em
Set Padding 4em
Set Padding 5em
```

<img alt="" src="./examples/settings/padding.gif" width="600" />

#### Set Framerate

Set the rate at which VHS captures frames with the `Set Framerate` command.

```elixir
Set Framerate 60
```

### Hide

The `Hide` command allows you to specify that the following commands should not
be shown in the output.

```elixir
Hide
```

### Show

The `Show` command allows you to specify that the following commands should
be shown in the output. Since this is the default case, the show command will
usually be seen with the `Hide` command.

```elixir
Hide
Type "You won't see this being typed."
Show
Type "You will see this being typed."
```

<img alt="Hide Show" src="./examples/commands/hide.gif" width="600" />


## Feedback

We’d love to hear your thoughts on this project. Feel free to drop us a note!

* [Twitter](https://twitter.com/charmcli)
* [The Fediverse](https://mastodon.technology/@charm)
* [Slack](https://charm.sh/slack)

## License

[MIT](https://github.com/charmbracelet/vhs/raw/main/LICENSE)

---

Part of [Charm](https://charm.sh).

<a href="https://charm.sh/">
  <img
    alt="The Charm logo"
    width="400"
    src="https://stuff.charm.sh/charm-badge.jpg"
  />
</a>

Charm热爱开源 • Charm loves open source
