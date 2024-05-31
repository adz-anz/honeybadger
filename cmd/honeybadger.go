package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	configKey     string
	apiHost       string
	targetDataset string
	dryRun		  bool
)

const (
	appName = "honeybadger"

	// The name of our config file, without the file extension because viper
	// supports many different config file languages.
	viperDefaultConfigFilename = "honeybadger"

	// The environment variable prefix of all environment variables bound to our
	// command line flags. For example, --number is bound to PREFIX_NUMBER.
	viperEnvPrefix = "HONEYBADGER"

	// Replace hyphenated flag names with camelCase in the config file
	viperReplaceHyphenWithCamelCase = false
)

type commandGroup struct {
	Name     string
	Commands []*cobra.Command
}

func (cg *commandGroup) commandWidth() int {
	width := 0
	for _, com := range cg.Commands {
		if com.Hidden {
			continue
		}
		newWidth := len(com.Name())
		if newWidth > width {
			width = newWidth
		}
	}
	return width
}

func displayCommands(cgs []commandGroup) {
	width := 0
	for _, cg := range cgs {
		newWidth := cg.commandWidth()
		if newWidth > width {
			width = newWidth
		}
	}

	for _, cg := range cgs {
		if cg.commandWidth() == 0 {
			continue
		}
		fmt.Printf("%s:\n", cg.Name)
		for _, com := range cg.Commands {
			if com.Hidden {
				continue
			}
			spacing := strings.Repeat(" ", width-len(com.Name()))
			fmt.Println("  " + com.Name() + spacing + strings.Repeat(" ", 8) + com.Short)
		}
		fmt.Println()
	}
}

func setCommandGroups(cmd *cobra.Command, rootCgs []commandGroup) {
	for _, cg := range rootCgs {
		for _, com := range cg.Commands {
			cmd.AddCommand(com)
		}
	}

	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		header := c.Long
		if header == "" {
			header = c.Short
		}

		if header != "" {
			fmt.Println(strings.TrimSpace(header))
			fmt.Println()
		}

		if c != cmd.Root() {
			fmt.Print(c.UsageString())
			return
		}

		fmt.Println("Usage:")
		fmt.Printf("  %v [command]\n", appName)
		fmt.Println()

		displayCommands(rootCgs)

		fmt.Println("Flags:")
		fmt.Println(cmd.Flags().FlagUsages())

		fmt.Printf("Use `%v [command] --help` for more information about a command.\n", appName)
	})
}

func NewHoneybadgerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   appName,
		Short: "Honeybadger command line",
		Long: "Honeybadger - Tearing Into Honeycomb\n" +
			"\n" +
			"TODO: Put some more stuff here",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but
			// PersistencePreRunE on the root command works well.
			return initializeConfig(cmd)
		},
	}
	cmd.PersistentFlags().StringVarP(&configKey, "configkey", "k", "",
		"Honeycomb configuration key from https://ui.honeycomb.io/<team>/environments/<environment>/api_keys")
	cmd.MarkPersistentFlagRequired("configkey")
	cmd.PersistentFlags().StringVar(&apiHost, "api_host",
		"https://api.honeycomb.io/", "The host to query, don't change it unless it's a hosted Honeycomb environment.")
	cmd.PersistentFlags().MarkHidden("api_host")
	cmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Print the request that would be sent without actually sending it.")

	setCommandGroups(cmd, []commandGroup{
		{
			Name: "Authorization Commands",
			Commands: []*cobra.Command{
				newAuthCmd(),
			},
		},
		{
			Name: "Board Commands",
			Commands: []*cobra.Command{
				newBoardsCmd(),
			},
		},
		// {
		// 	Name: "Burn Alert Commands",
		// 	Commands: []*cobra.Command{
		// 		newBurnAlertsCmd(),
		// 	},
		// },
		// {
		// 	Name: "Column Commands",
		// 	Commands: []*cobra.Command{
		// 		newColumnsCmd(),
		// 		newDerivedColumnsCmd(),
		// 	},
		// },
		{
			Name: "Dataset Commands",
			Commands: []*cobra.Command{
				newDatasetsCmd(),
				newDatasetDefinitionsCmd(),
			},
		},
		// {
		// 	Name: "Event Commands",
		// 	Commands: []*cobra.Command{
		// 		newEventCmd(),
		// 		newKinesisEventCmd(),
		// 	},
		// },
		{
			Name: "Marker Commands",
			Commands: []*cobra.Command{
				newMarkersCmd(),
				newMarkerSettingsCmd(),
			},
		},
		// {
		// 	Name: "Query Commands",
		// 	Commands: []*cobra.Command{
		// 		newQueriesCmd(),
		// 		newQueryAnnotationsCmd(),
		// 		newQueryDataCmd(),
		// 	},
		// },
		// {
		// 	Name: "Recipient Commands",
		// 	Commands: []*cobra.Command{
		// 		newRecipientsCmd(),
		// 	},
		// },
		// {
		// 	Name: "SLO Commands",
		// 	Commands: []*cobra.Command{
		// 		newSLOsCmd(),
		// 	},
		// },
		// {
		// 	Name: "Trigger Commands",
		// 	Commands: []*cobra.Command{
		// 		newTriggersCmd(),
		// 	},
		// },
	})

	return cmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	// Set the base name of the config file, without the file extension.
	v.SetConfigName(viperDefaultConfigFilename)

	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath(".")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// When we bind flags to environment variables expect that the environment
	// variables are prefixed, e.g. a flag like --number binds to an environment
	// variable PREFIX_NUMBER. This helps avoid conflicts.
	v.SetEnvPrefix(viperEnvPrefix)

	// Environment variables can't have dashes in them, so bind them to their
	// equivalent keys with underscores, e.g. --favorite-color to
	// PREFIX_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name

		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if viperReplaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
