/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package args

import (
	"fmt"

	"github.com/spf13/pflag"

	"k8s.io/gengo/v2"
)

type Args struct {
	OutputFile    string
	ExtraPeerDirs []string // Always consider these as last-ditch possibilities for conversions.
	GoHeaderFile  string

	// GeneratedBuildTag is the tag used to identify code generated by execution
	// of this type. Each generator should use a different tag, and different
	// groups of generators (external API that depends on Kube generations) should
	// keep tags distinct as well.
	GeneratedBuildTag string
}

// New returns default arguments for the generator.
func New() *Args {
	return &Args{
		GeneratedBuildTag: gengo.StdBuildTag,
	}
}

// AddFlags add the generator flags to the flag set.
func (args *Args) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&args.OutputFile, "output-file", "generated.defaults.go",
		"the name of the file to be generated")
	fs.StringSliceVar(&args.ExtraPeerDirs, "extra-peer-dirs", args.ExtraPeerDirs,
		"Comma-separated list of import paths which are considered, after tag-specified peers, for conversions.")
	fs.StringVar(&args.GoHeaderFile, "go-header-file", "",
		"the path to a file containing boilerplate header text; the string \"YEAR\" will be replaced with the current 4-digit year")
	fs.StringVar(&args.GeneratedBuildTag, "build-tag", args.GeneratedBuildTag, "A Go build tag to use to identify files generated by this command. Should be unique.")
}

// Validate checks the given arguments.
func (args *Args) Validate() error {
	if len(args.OutputFile) == 0 {
		return fmt.Errorf("--output-file must be specified")
	}

	return nil
}
