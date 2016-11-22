# arcanum
A Golang API for delivering D&D 5e spell data to spellfoc.us and the tome app.

## Setup
### Go Version
If you haven't already done so, you will need to install Go 1.7 or higher. Please see [the Golang site](https://golang.org) for more info and instructions.

### Build
To build the arcanum application, simply run ```go build``` or ```go install```.

### Run
To run the arcanum application without building, run ```go run *.go```; if the application was built or installed (as in the previous step), simply run ```arcanum```.

Once the server app is running on your machine, navigate to [localhost:8080](http://localhost:8080) to access arcanum.

## Routes
The application currently has a bare-bones HTML UI. It also supports an API interface to provide spell data to other projects. The current routes supported are:
* /api/list - lists all spells in the arcanum database.
* /api/spell/# - lists the spell data for a spell with the ID of '#'.

## Still to Come/TODOs
* Refining of spell data (data is not included in the project)
* More routes
* UI improvements
* And *more!*

Special thanks to [Wizards of the Coast](http://company.wizards.com/) for D&D 5e - a great game and the inspiration for this app!
