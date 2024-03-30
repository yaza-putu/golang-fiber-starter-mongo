package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/core"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/database"
	_ "github.com/yaza-putu/golang-fiber-starter-mongo/internal/database/migrations"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/pkg/logger"
	"github.com/yaza-putu/golang-fiber-starter-mongo/pkg/unique"
)

func main() {
	if os.Args[1] != "key:generate" {
		core.Env()
		core.Mongo()
	}

	command := New()

	if len(os.Args) == 1 {
		fmt.Println("Hello i'm zoro, can i help you ?")
		fmt.Println("Available command :")

		// migration collection
		m := []string{
			"- key:generate",
			"- make:migration",
			"- migrate:up",
			"- migrate:down",
			"- make:seeder",
			"- seed:up",
		}
		for _, v := range m {
			fmt.Println(v)
		}
		os.Exit(0)
	}

	for i, v := range os.Args {
		if i != 0 {
			switch v {
			case "make:migration":
				command.newMigration()
				break
			case "migrate:up":
				command.upMigration()
				break
			case "migrate:down":
				command.downMigration()
				break
			case "key:generate":
				command.keyGenerate()
				break
			}
		}
	}
}

type (
	artisanCommand   struct{}
	artisanInterface interface {
		newMigration() bool
		upMigration() bool
		downMigration() bool
		keyGenerate() bool
	}
)

func New() artisanInterface {
	return &artisanCommand{}
}

func (z *artisanCommand) newMigration() bool {
	if len(os.Args) != 3 {
		fmt.Println("ex : make:migration name-of-file")
		return false
	}

	// file
	fName := fmt.Sprintf("./internal/database/migrations/%s_create_collection_%s.go", time.Now().Format("20060102150405"), os.Args[2])

	// from template
	from, err := os.Open("./internal/database/migration.stub")
	logger.New(err, logger.SetType(logger.FATAL))

	defer from.Close()

	// to file
	to, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0666)
	logger.New(err, logger.SetType(logger.FATAL))

	defer to.Close()
	// copy file with template
	_, err = io.Copy(to, from)
	logger.New(err, logger.SetType(logger.FATAL))

	fmt.Printf("New migration : %s\n", fName)

	return true
}

func (z *artisanCommand) upMigration() bool {
	err := database.UpMIgration()
	if err != nil {
		logger.New(err, logger.SetType(logger.FATAL))
		return false
	} else {
		fmt.Println("Migrating collections successfully")
		return true
	}
}

func (z *artisanCommand) downMigration() bool {
	err := database.DownMigration()
	if err != nil {
		logger.New(err, logger.SetType(logger.FATAL))
		return false
	} else {
		fmt.Println("Drop collections successfully")
		return true
	}
}

func (z *artisanCommand) keyGenerate() bool {
	key := unique.Key(51)

	findAndReplaceByKey("APP_KEY", fmt.Sprintf("'%s'", key))

	fmt.Println("Generate key successfully")
	return true
}

func findAndReplaceByKey(key, newValue string) error {
	filename := ".env"
	// Read the entire file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Split content by lines
	lines := strings.Split(string(content), "\n")

	// Find and replace the key if found
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = key + "=" + newValue
			found = true
			break
		}
	}

	// If key not found, return an error
	if !found {
		return fmt.Errorf("key '%s' not found in file", key)
	}

	// Write the modified content back to the file
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}
