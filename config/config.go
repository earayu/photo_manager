package config

import flag "github.com/spf13/pflag"

type globalConfig struct {
	inputDir       string
	outputDir      string
	targetFileName string
	sourceType     []string
}

func (g globalConfig) InputDir() string {
	return g.inputDir
}

func (g globalConfig) OutputDir() string {
	if g.outputDir == "" {
		return g.inputDir + "/output"
	}
	return g.outputDir
}

func (g globalConfig) TargetFileName() string {
	return g.targetFileName
}

func (g globalConfig) FullNameTargetFileName(fileType string) string {
	return g.OutputDir() + "/" + g.TargetFileName() + "." + fileType
}

func (g globalConfig) SourceType() []string {
	return g.sourceType
}

var (
	GlobalConfig = &globalConfig{}
)

func init() {
	flag.StringVar(&GlobalConfig.inputDir, "input_dir", "~/Documents/GitHub/photo_manager/testdata/resizer", "The base directory of the project")
	flag.StringVar(&GlobalConfig.outputDir, "output_dir", "", "The base directory of the project")
	flag.StringVar(&GlobalConfig.targetFileName, "target_file_name", "target", "The base directory of the project")
	flag.StringArrayVar(&GlobalConfig.sourceType, "source_type", []string{"jpg", "jpeg", "png"}, "The base directory of the project")
}
