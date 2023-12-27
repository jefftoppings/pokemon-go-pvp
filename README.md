# pokemon-go-pvp
A server built in Go that provides information about Pokemon Go PVP. 

## Updating pokemon source data
Data is sourced from https://github.com/Sigafoos/iv. To regenerate the data, do the following:
- Clone the repo
- Download the current [gamemaster.json file from pvpoke](https://github.com/pvpoke/pvpoke/blob/master/src/data/gamemaster.json) and save it in the `generator/` directory
- From the `generator/` directory, run `go build`. That should create an executable called `generator`
- Run the `generator` executable by running `./generator`. That should re-populate the files in the `data/` directory.
- Copy the entire `data/` directory into this project and replace `internal/assets/data`
