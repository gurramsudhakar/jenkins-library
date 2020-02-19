package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/SAP/jenkins-library/pkg/command"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
)

const mavenExecutable = "mvn"

func mavenExecute(config mavenExecuteOptions, telemetryData *telemetry.CustomData) {
	c := command.Command{}
	// stdOud := log.Entry().Writer()

	// if config.ReturnStdout {
	// 	stdOutBuf = new(bytes.Buffer)
	// 	stdOut = io.MultiWriter(log.Entry().Writer(), &stdOutBuf)
	// }
	stdOutBuf := new(bytes.Buffer)
	stdOut := io.MultiWriter(os.Stdout, stdOutBuf)

	c.Stdout(stdOut)
	c.Stderr(log.Entry().Writer())

	// error situations should stop execution through log.Entry().Fatal() call which leads to an os.Exit(1) in the end
	err := runMavenExecute(config, &c) //runMavenExecute(&config, telemetryData, &c)
	if err != nil {
		log.Entry().WithError(err).Fatal("step execution failed")
	}
}

func runMavenExecute(config mavenExecuteOptions, command execRunner) error {
	return command.RunExecutable("mvn", "--version")
}

// func runMavenExecute(config *mavenExecuteOptions, telemetryData *telemetry.CustomData, command execRunner) error, string {

// 	parameters := []string{}

// 	if config.GlobalSettingsFile != "" {
// 		globalSettingsFileParameter := "--global-settings " + config.GlobalSettingsFile
// 		parameters = append(parameters, globalSettingsFileParameter)
// 	}

// 	if config.ProjectSettingsFile != "" {
// 		projectSettingsFileParameter := "--settings " + config.ProjectSettingsFile
// 		parameters = append(parameters, projectSettingsFileParameter)
// 	}

// 	if config.M2Path != "" {
// 		m2PathParameter := "-Dmaven.repo.local=" + config.M2Path
// 		parameters = append(parameters, m2PathParameter)
// 	}

// 	if config.PomPath != "" {
// 		pomPathParameter := "--file " + config.PomPath
// 		parameters = append(parameters, pomPathParameter)
// 	}

// 	if config.Flags != nil {
// 		parameters = append(parameters, config.Flags...)
// 	}

// 	parameters = append(parameters, "--batch-mode")

// 	if config.LogSuccessfulMavenTransfers {
// 		parameters = append(parameters, "-Dorg.slf4j.simpleLogger.log.org.apache.maven.cli.transfer.Slf4jMavenTransferListener=warn")
// 	}

// 	parameters = append(parameters, config.Goals...)

// 	err := command.RunExecutable(mavenExecutable, parameters...)
// 	if err != nil {
// 		log.Entry().
// 			WithError(err).
// 			WithField("command", config.RunCommand).
// 			Fatal("failed to execute run command")
// 	}

// 	return nil
// }
