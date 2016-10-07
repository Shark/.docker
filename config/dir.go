package config

import (
  "os"
  "os/user"
  "log"
  "path"
  "path/filepath"
)

func FindGlobalConfigDir() (string) {
  user, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  configDir := path.Join(user.HomeDir, ".orca")
  _, err = os.Stat(configDir)
  if err == nil {
    return configDir
  } else {
    return ""
  }
}

func FindLocalConfigDir() (string) {
  workingDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  configDir := findLocalConfigDirRecursive(workingDir)
  if configDir != "" {
    user, err := user.Current()
    if err != nil {
      log.Fatal(err)
    }
    if path.Clean(path.Join(configDir, "..")) != user.HomeDir {
      return configDir
    } else {
      return ""
    }
  } else {
    return ""
  }
}

func findLocalConfigDirRecursive(searchPath string) string {
  curPath := path.Join(searchPath, ".orca")
  _, err := os.Stat(curPath)
  if err == nil {
    return curPath
  } else {
    parentPath := path.Clean(path.Join(searchPath, ".."))

    if parentPath != searchPath { // if we are at /, then parentPath == path
      return findLocalConfigDirRecursive(parentPath)
    } else {
      return ""
    }
  }
}
