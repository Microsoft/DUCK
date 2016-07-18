package ducklib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/Microsoft/DUCK/backend/ducklib/structs"
)

func NewConfiguration(confpath string) structs.Configuration {

	c := structs.Configuration{}

	//setting defaults
	c.JwtKey = "secret"
	c.WebDir = "/src/github.com/Microsoft/DUCK/frontend/dist"
	c.RulebaseDir = "/src/github.com/Microsoft/DUCK/RuleBases"
	c.Gopathrelative = true
	c.Loadtestdata = false

	if err := getFileConfig(&c, confpath); err != nil {
		log.Printf("Could not load configuration file: %s", err)

	}

	getEnv(&c)

	c.DBConfig.Port = 5984
	c.Gopathrelative = true
	log.Printf("Config: %+v", c)
	log.Printf("Datab: %+v", c.DBConfig)
	return c
}

func getFileConfig(config *structs.Configuration, confpath string) error {
	dat, err := ioutil.ReadFile(confpath)
	if err != nil {
		return err

	}
	err = json.Unmarshal(dat, &config)
	return err

}

func getEnv(c *structs.Configuration) {
	//Get Environment Variables

	env := os.Getenv("DUCK_JWTKEY")
	if env != "" {
		c.JwtKey = env
	}

	log.Printf("ENV %s", os.Getenv("DUCK_TEST"))
	env = os.Getenv("DUCK_WEBDIR")
	if env != "" {
		c.WebDir = env
	}

	env = os.Getenv("DUCK_RULEBASEDIR")
	if env != "" {
		c.RulebaseDir = env
	}
	env = os.Getenv("DUCK_GOPATHRELATIVE")
	if env != "" {
		if gpr, err := strconv.ParseBool(env); err == nil {
			c.Gopathrelative = gpr
		} else {
			log.Printf("Could not read value for GOPATHRELATIVE: %s", err)
		}
	}
	env = os.Getenv("DUCK_LOADTESTDATA")
	if env != "" {
		if ldt, err := strconv.ParseBool(env); err == nil {
			c.Loadtestdata = ldt
		} else {
			log.Printf("Could not read value for LOADTESTDATA: %s", err)
		}

	}
	env = os.Getenv("DUCK_DATABASE.LOCATION")
	if env != "" {
		c.DBConfig.Location = env
	}
	env = os.Getenv("DUCK_DATABASE.PORT")
	if env != "" {
		if p, err := strconv.Atoi(env); err == nil {
			c.DBConfig.Port = p
		} else {
			log.Printf("Could not read value for PORT: %s", err)
		}

	}
	env = os.Getenv("DUCK_DATABASE.USERNAME")
	if env != "" {
		c.DBConfig.Username = env
	}
	env = os.Getenv("DUCK_DATABASE.PASSWORD")
	if env != "" {
		c.DBConfig.Password = env
	}
	env = os.Getenv("DUCK_DATABASE.NAME")
	if env != "" {
		c.DBConfig.Name = env
	}
}
