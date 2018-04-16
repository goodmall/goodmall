package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-config"
)

func main() {
	// create a Config object
	c := config.New()

	// load configuration from a JSON string
	c.LoadJSON([]byte(`{
        "Version": "2.0",
        "Author": {
            "Name": "Foo",
            "Email": "bar@example.com"
        }
    }`))

	// get the "Version" value, return "1.0" if it doesn't exist in the config
	version := c.GetString("Version", "1.0")

	var author struct {
		Name, Email string
	}
	// populate the author object from the "Author" configuration
	c.Configure(&author, "Author")

	fmt.Println(version)
	fmt.Println(author.Name)
	fmt.Println(author.Email)
	// Output:
	// 2.0
	// Foo
	// bar@example.com

	loadFromTOML(c, "app.toml")
}

func loadFromTOML(c *config.Config, conf string) {
	// NOTE 如果是多文件加载 后者会覆盖前者内容  这个跟 金柱的configor库 刚好相反  有点jquery.extend 的感觉
	c.Load(conf /* "app.toml" , "app.dev.json"*/)

	// Retrieves "Author.Email". The default value "bar@example.com"
	// should be returned if "Author.Email" is not found in the configuration.
	email := c.GetString("Author.Email", "bar@example.com")
	email2 := c.GetString("adminEmail", "somedefaultname@example.com")

	fmt.Printf(" email is : %s and amdin-email is : %s ", email, email2)
}
