/*
Everything from the root of this repo downwards, is called a 'module'. A module
can contain one or more packages. A 'package' is essentially a folder of golang code
in the module. The root folder by design houses 'package main'. Any other child folder
such as the 'members' folder is referred to as packages.

child Packages must be named after the name of the folder that houses then.
Therefore in this case it is:
*/
package members

// Employee - this is employee struct
type Employee struct {
	ID        int
	Title     string
	FirstName string
	Lastname  string
	Age       int
	Smoker    bool
}