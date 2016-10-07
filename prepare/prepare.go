package prepare

import (
  "os"
  "fmt"
  "path"
  "strings"

  "github.com/fatih/color"
  "github.com/codeskyblue/go-sh"
)

func Prepare(projectDir string, projectName string, forceRebuild bool) error {
  rebuild := forceRebuild
  imageID := fmt.Sprintf("sh4rk/%s-dev", projectName)

  if !rebuild {
    _, err := sh.Command("docker", "inspect", imageID).Output()
    if err != nil {
      rebuild = true
    }
  }

  if rebuild {
    color.Green(">> run docker build")
    err := sh.Command("docker", "build", "-t", imageID, "--force-rm", "-f", ".orca/Dockerfile", ".", sh.Dir(projectDir)).Run()
    if err != nil {
      return err
    }
  }

  _, err := os.Stat(path.Join(projectDir, ".orca", "prepare.sh"))
  if err == nil {
    color.Green(">> run docker prepare")
    out, err := sh.Command("docker-compose", "--file", ".orca/docker-compose.yml", "--project-name", projectName, "run", "-d", "app", ".orca/prepare.sh", sh.Dir(projectDir)).Output()
    if err != nil {
      return err
    }
    containerID := strings.TrimSpace(string(out))

    err = sh.Command("docker", "logs", "-f", containerID, sh.Dir(projectDir)).Run()
    if err != nil {
      return err
    }

    color.Green(">> commit the result")
    err = sh.Command("docker", "commit", containerID, imageID, sh.Dir(projectDir)).Run()
    if err != nil {
      return err
    }

    out, err = sh.Command("docker", "rm", "-f", "-v", containerID, sh.Dir(projectDir)).Output()
    if err != nil {
      fmt.Print(string(out))
      return err
    }
  } else {
    color.Green(">> skip docker prepare")
  }
  return nil
}
