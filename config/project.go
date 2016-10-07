package config

import (
  "errors"
  "os"
  "path"
)

func ProjectNameFromConfigDir(dir string) string {
  parentPath := path.Clean(path.Join(dir, ".."))
  return path.Base(parentPath)
}

func ValidateConfigDir(dir string) error {
  if _, err := os.Stat(path.Join(dir, "docker-compose.yml")); err != nil {
    return errors.New("docker-compose.yml does not exist")
  }

  return nil
}
