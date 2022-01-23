package formicidae

import (
	"os"
	"testing"
)

const filename = "tests.env"

func setup() {
	_, _ = os.Create(filename)
}

func teardown() {
	_ = os.Remove(filename)
}

func TestUpdateOneVariable(t *testing.T) {
	//setup
	setup()

	updateFileWithContent(
		"DATABASE_ADDRESS=localhost\n"+
			"DATABASE_NAME=postgres\n"+
			"DATABASE_USERNAME=administrat0r\n"+
			"DATABASE_PASSWORD=passw0rd\n"+
			"DATABASE_PORT=5432", false)

	//arrange
	expected := "DATABASE_ADDRESS=localhost\n" +
		"DATABASE_NAME=postgres\n" +
		"DATABASE_USERNAME=administrat0r\n" +
		"DATABASE_PASSWORD=NewPassw0rd\n" +
		"DATABASE_PORT=5432"

	//act
	got, _ := UpdateVariable(filename, "DATABASE_PASSWORD", "NewPassw0rd")

	//assert
	if got != expected {
		t.Error("Environment variable is not updated correctly.")
	}

	//teardown
	teardown()
}

func TestUpdateTwoVariables(t *testing.T) {
	//setup
	setup()

	updateFileWithContent(
		"DATABASE_ADDRESS=localhost\n"+
			"DATABASE_NAME=postgres\n"+
			"DATABASE_USERNAME=administrat0r\n"+
			"DATABASE_PASSWORD=passw0rd\n"+
			"DATABASE_PORT=5432", false)

	//arrange
	expected := "DATABASE_ADDRESS=localhost\n" +
		"DATABASE_NAME=postgres\n" +
		"DATABASE_USERNAME=NewAdministrat0r\n" +
		"DATABASE_PASSWORD=NewPassw0rd\n" +
		"DATABASE_PORT=5432"

	//act
	_, _ = UpdateVariable(filename, "DATABASE_USERNAME", "NewAdministrat0r")
	got, _ := UpdateVariable(filename, "DATABASE_PASSWORD", "NewPassw0rd")

	//assert
	if got != expected {
		t.Error("Environment variables are not updated correctly.")
	}

	//teardown
	teardown()
}

func TestUpdateDuplicateVariables(t *testing.T) {
	//setup
	setup()

	updateFileWithContent(
		"DATABASE_ADDRESS=localhost\n"+
			"DATABASE_NAME=postgres\n"+
			"DATABASE_USERNAME=administrat0r\n"+
			"DATABASE_PASSWORD=passw0rd\n"+
			"DATABASE_PORT=5432\n"+
			"DATABASE_USERNAME=admin", false)

	//arrange
	expected := "DATABASE_ADDRESS=localhost\n" +
		"DATABASE_NAME=postgres\n" +
		"DATABASE_USERNAME=NewAdministrat0r\n" +
		"DATABASE_PASSWORD=passw0rd\n" +
		"DATABASE_PORT=5432\n" +
		"DATABASE_USERNAME=NewAdministrat0r"

	//act
	got, _ := UpdateVariable(filename, "DATABASE_USERNAME", "NewAdministrat0r")

	//assert
	if got != expected {
		t.Error("Duplicated Environment variables are not updated correctly.")
	}

	//teardown
	teardown()
}

func TestUpdateMissingFile(t *testing.T) {
	//setup
	setup()

	updateFileWithContent(
		"DATABASE_ADDRESS=localhost\n"+
			"DATABASE_NAME=postgres\n"+
			"DATABASE_USERNAME=administrat0r\n"+
			"DATABASE_PASSWORD=passw0rd\n"+
			"DATABASE_PORT=5432\n"+
			"DATABASE_USERNAME=admin", false)

	//arrange
	const missingFile = ".env"

	//act
	_, err := UpdateVariable(missingFile, "DATABASE_USERNAME", "NewAdministrat0r")

	//assert
	if err == nil {
		t.Error("Expected error for missing file")
	}

	//teardown
	teardown()
}

func updateFileWithContent(content string, readOnly bool) bool {
	var permissions = os.FileMode(0644)
	flag := os.O_RDWR
	if readOnly {
		permissions = os.FileMode(44)
		flag = os.O_RDONLY
	}

	file, _ := os.OpenFile(filename, flag, permissions)

	_, err := file.WriteString(content)
	if err != nil {
		_ = file.Close()

		return false
	}

	err = file.Close()
	return err == nil
}
