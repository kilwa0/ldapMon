package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/viper"
	"gopkg.in/ldap.v3"
)

func ldapDial(host string, port int, user string, pass string, basedn string) []*ldap.Entry {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// First bind with a read only user
	err = l.Bind(user, pass)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		basedn, // The base dn to search
		ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",      // The filter to apply
		[]string{"contextCSN"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	return sr.Entries
}

func contextCSNCompare(context1, context2 []*ldap.Entry) {
	if reflect.DeepEqual(context1, context2) == true {
		log.Println("Synchronized")
	} else {
		log.Fatal("Desynchronized")
	}
}

func main() {

	viper.SetConfigFile("cred.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	username := viper.GetString("username")
	password := viper.GetString("pass")
	host1 := viper.GetString("host1")
	port1 := viper.GetInt("port1")
	host2 := viper.GetString("host2")
	port2 := viper.GetInt("port2")
	basedn := viper.GetString("basedn")

	fmt.Printf("Configuration:\nusername: %v\npassword: %v\nhost1: %v\nhost2: %v\nbasedn: %v\n", username, password, host1, host2, basedn)

	if host1 != "" && host2 != "" {
		context1 := ldapDial(host1, port1, username, password, basedn)
		context2 := ldapDial(host2, port2, username, password, basedn)
		contextCSNCompare(context1, context2)
	} else if host1 != "" {
		ldapDial(host1, port1, username, password, basedn)
	} else {
		fmt.Println("-host1 mandatory: At least one host must be given")
	}
}
