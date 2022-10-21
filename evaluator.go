package vhs

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func Evaluate(tape string, w io.Writer, outputFile string) error {
	l := NewLexer(tape)
	p := NewParser(l)

	cmds := p.Parse()
	errs := p.Errors()
	if len(errs) != 0 {
		lines := strings.Split(tape, "\n")
		for _, err := range errs {
			fmt.Fprint(w, LineNumber(err.Token.Line))
			fmt.Fprintln(w, lines[err.Token.Line-1])
			fmt.Fprint(w, strings.Repeat(" ", err.Token.Column+5))
			fmt.Fprintln(w, Underline(len(err.Token.Literal)), err.Msg)
			fmt.Fprintln(w)
		}
		return errors.New("parse error")
	}

	// Start VHS

	v := New()

	var offset int
	for i, cmd := range cmds {
		if cmd.Type == SET || cmd.Type == OUTPUT {
			fmt.Fprintln(w, cmd.Highlight(false))
			cmd.Execute(&v)
		} else {
			offset = i
			break
		}
	}

	v.Setup()
	v.Record()
	defer v.Cleanup()

	for _, cmd := range cmds[offset:] {
		// FIXME: When changing the FontFamily, FontSize, LineHeight, Padding
		// The xterm.js canvas changes dimensions and causes FFMPEG to not work
		// correctly (specifically) with palettegen.
		// It will be possible to change settings on the fly in the future, but it is currently not
		// as it does not result in a proper render of the GIF as the frame sequence
		// will change dimensions. This is fixable.
		//
		// TODO: Remove if isSetting statement.
		isSetting := cmd.Type == SET && cmd.Options != "TypingSpeed"
		if isSetting {
			fmt.Fprintln(w, cmd.Highlight(true))
			continue
		}
		fmt.Fprintln(w, cmd.Highlight(!v.recording || cmd.Type == SHOW || cmd.Type == HIDE || isSetting))
		cmd.Execute(&v)
	}

	// If running as an SSH server, the output file is a temporary file
	// to use for the output.
	//
	// We need to do this before the GIF is created but after all of the settings
	// and commands are executed.
	//
	// Since the GIF creation is deferred, setting the output file here will
	// achieve what we want.
	if outputFile != "" {
		v.Options.Video.Output.GIF = outputFile
	}

	return nil
}
