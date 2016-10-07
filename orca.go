package main

import (
  "fmt"
  "os"
  "os/exec"
  "path"

  "github.com/Shark/orca/config"
  "github.com/Shark/orca/prepare"
  "github.com/fatih/color"
)

const (
  MODE_LOCAL = iota
  MODE_GLOBAL = iota
)

func printHelp() {
  fmt.Println("Usage: orca [--global] CMD")
  fmt.Println("")
  fmt.Println("Specify --global to use the configuration from ~/.orca")
  fmt.Println("CMD is either \"prepare\", or any docker-compose command")
  fmt.Println("e.g. orca prepare --force-rebuild, orca up -d, orca run app rails c")
}

func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    printHelp()
    os.Exit(1)
  }

  mode := MODE_LOCAL
  cmd := args
  if args[0] == "--global" {
    mode = MODE_GLOBAL
    cmd = args[1:]
    if len(cmd) == 0 {
      printHelp()
      os.Exit(1)
    }
  }

  var configDir string
  if mode == MODE_LOCAL {
    color.Blue(">> local configuration")
    configDir = config.FindLocalConfigDir()
  } else if mode == MODE_GLOBAL {
    color.Blue(">> global configuration")
    configDir = config.FindGlobalConfigDir()
  }

  if configDir != "" {
    color.Blue(">> configuration location %s", configDir)
  } else {
    color.Red(">> config could not be found, exiting")
    os.Exit(1)
  }

  if err := config.ValidateConfigDir(configDir); err != nil {
    color.Red(">> error in config dir: %v", err)
    os.Exit(1)
  }

  projectName := config.ProjectNameFromConfigDir(configDir)
  if cmd[0] == "prepare" {
    forceRebuild := cmd[len(cmd)-1] == "--force-rebuild"
    prepare.Prepare(path.Clean(path.Join(configDir, "..")), projectName, forceRebuild)
  } else {
    composeArgs := []string{"--file", path.Join(configDir, "docker-compose.yml"), "--project-name", projectName}
    composeArgs = append(composeArgs, cmd...)
    composeCmd := exec.Command("docker-compose", composeArgs...)
    composeCmd.Stdin = os.Stdin
    composeCmd.Stdout = os.Stdout
    composeCmd.Stderr = os.Stderr
    composeCmd.Run()
    composeCmd.Wait()
  }
}
