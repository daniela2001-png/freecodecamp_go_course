package packages

// --- MODULES ---
/*

	Go programs are organized into packages.
	A package  is a directory  of Go code that is all compiled together

	A repository contains one or more modules. A module is a collection
	of Go packages that are released together.


	A GO REPOSITORY TYPICALLY CONTAINS ONLY ONE MODULE, LOCATED
	AT THE ROOT OF THE REPOSITORY.

		- A file named go.mod  at the root of a project  declares the module
		It contains:
			- The module path (is just the import path prefix for all packages within the module)
			- The version of the Go language your project requires.
			- Optionally, any external package dependencies that the project has.

	What is an import path ?
		- A module path  + package subdirectory

	Why does Go include a remote URL in module paths (module github.com/daniela2001-png/freecodecamp_go_course
	) ?
		- In this case "freecodecamp_go_course" at the end of the module path, will be the name of our package and also the name of our repository at GitHub
		- To simplify remote downloading of packages

*/
