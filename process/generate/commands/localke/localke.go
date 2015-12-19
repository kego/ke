package localke // import "kego.io/process/generate/commands/localke"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"kego.io/context/cmdctx"
	"kego.io/editor/server"
	"kego.io/process"
)

// Run executes the local "ke" command that is generated in the data directory.
func Main(recursive bool, path string) {
	update := false
	ctx, cancel, err := process.Initialise(&process.FromFlags{
		Update:    &update,
		Recursive: &recursive,
		Path:      &path,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		fmt.Println("No cmd in ctx")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()

	typesChanged, err := process.ValidateCommand(ctx)
	if err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
	if typesChanged {
		if cmd.Verbose {
			fmt.Println("Types have changed - rebuilding...")
		}
		if err := process.RunAllCommands(ctx); err != nil {
			if !cmd.Verbose {
				// in verbose mode, we have already written the output of the exec'ed ke command,
				// so we don't need to duplicate the error message.
				fmt.Println(process.FormatError(err))
			}
			os.Exit(1)
		}
		// if the types have changed, RunAllCommands will opened the server etc.
		os.Exit(0)
	}
	if cmd.Edit {
		if err = server.Start(ctx); err != nil {
			fmt.Println(process.FormatError(err))
			os.Exit(1)
		}
	}
}
