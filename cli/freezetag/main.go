package main

import (
  "os"
  "github.com/codegangsta/cli"
)

var (
  Version string
  BuildTime string
)

func dir_exists(path) (bool, error) {
  _, err := os.Stat(path)
  if err == nil { return true, nil }
  if os.IsNotExist(err) { return false, nil }
  return true, err
}

func is_dir(path) (bool, error) {
  file_info, err := os.Stat(path)
  if err != nil {
    if os.IsNotExist(err) { return false, nil }
    return false, err
  }
  return file_info.IsDir(), nil
}

func validate_target_dir(target_dir) (bool, error) {
  // ensure dir exists
  exists, err := dir_exists(target_dir)
  if err != nil { return false, err }
  if !exists { return false, nil }

  // ensure dir is a dir
  is_a_dir, err := is_dir(target_dir)
  if err != nil { return false, err }
  if !is_a_dir { return false, nil }

  // ensure user has write access to dir
}

func main() {
  app := cli.NewApp()
  app.Name = "freezetag"
  app.Usage = "Discovers tags on an EC2 instance and saves them to disk in JSON format"
  app.Version = Version
  app.HideVersion = true
  app.HideHelp = true
  app.Copyright = "2016"
  app.Authors = []cli.Author {
    cli.Author{
      Name: "Scott Brown",
    },
  }

  verbose := false
  showHelp := false
  showVersion := false
  target_dir := ""

  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "verbose, v",
      Usage: "Adds more information to the output",
      Destination: &verbose,
    },
    cli.BoolFlag{
      Name: "help, h",
      Usage: "Displays this message",
      Destination: &showHelp,
    },
    cli.BoolFlag{
      Name: "version",
      Usage: "Displays the version",
      Destination: &showVersion,
    },
    cli.StringFlag{
      Name: "directory, d",
      Usage: "The directory to save the tags",
      Destination: &target_dir,
    },
  }

  app.Action = func(c *cli.Context) {
    if (showHelp) {
      println("Help here")
      return
    }

    if (showVersion) {
      println(app.Version)
      if (verbose) {
        println("Build Time:", BuildTime)
      }
      return
    }

    if (err := validate_target_dir() != nil) {

    }

    // save the aws tags
  }

  app.Run(os.Args)
}
