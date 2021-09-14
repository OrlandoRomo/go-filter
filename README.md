
# Go Filter
Go Filter is a CLI-tool to transform an image to any filter such a grey-scale, inverted colors etc

# How to start
```bash
# Clone the repository
git clone https://github.com/OrlandoRomo/go-filter.git

# Change directory to go-filter
cd go-filter

# Install dependencies
go mod tidy

# Build project
cd cmd/go-filter && go build .
```
Move the binary in your local environment
  - Linux/Mac: `sudo go-filter /usr/local/bin`

# Go Filter manual
To use `go-filter` open a new terminal in your current operative system and type `go-filter` and enter.

## Commands
The `go-filter` tool integrates the following commands and subcommands
   - `filter`: filter will apply the given filter name
   - `filter`'s flags
      - --output: Path where the final result will be created. By default is the user directory.  
   - `filter`'s subcommands:  
      - `list`: displays all available filter names
         - args: path of the image as a input 
         
      - `gray`: applies the gray scale effect
         - args: path of the image as a input 

      - `negative`: applies the negative scale effect
         - args: path of the image as a input 

      - `red`: applies the red scale effect
           - args: path of the image as a input 

      - `blue`: applies the blue scale effect
           - args: path of the image as a input
  
      - `green`: applies the green scale effect
           - args: path of the image as a input
 
      - `mirror`: applies the mirror effect
           - args: path of the image as a input 

      - `sepia`: applies the sepia effect
           - args: path of the image as a input 

      - `sketch`: applies the sketch effect
           - args: path of the image as a input

## Examples
| Command       |    Result     |
| ------------- |:-------------:|
| `go-filter filter gray $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297171-c9b00477-4a1e-49ad-8d6d-0b730ba0285f.jpg" width="150" height="150"> |
| `go-filter filter negative $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297151-c4494112-7856-4c27-aae1-b07a2bd6b384.jpg" width="150" height="150"> |
| `go-filter filter red $HOME/Desktop/future_nostalgia.jpg`      | <img src="https://user-images.githubusercontent.com/34588445/133297177-714859ee-301c-429e-851a-dce40378e25c.jpg" width="150" height="150"> |
| `go-filter filter blue $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297188-843f51fe-f9d7-473d-9c54-ae7a1faf25f2.jpg" width="150" height="150"> |
| `go-filter filter green $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297128-16c7ad56-f2f6-4a8d-8684-d1c795177e5c.jpg" width="150" height="150"> |
| `go-filter filter mirror $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297175-8d2aa032-902d-4f6a-8459-2c52ad12148b.jpg" width="150" height="150"> |
| `go-filter filter sepia $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297141-c022155d-05ef-4162-a61e-4920509cad8e.jpg" width="150" height="150"> |
| `go-filter filter sketch $HOME/Desktop/future_nostalgia.jpg`      | <img src="https://user-images.githubusercontent.com/34588445/133297150-646feaaa-4126-46d7-aecc-2c3df93e28aa.jpg" width="150" height="150"> |
| `go-filter filter sharp $HOME/Desktop/future_nostalgia.jpg`      |  <img src="https://user-images.githubusercontent.com/34588445/133297162-3abbd4b1-1d35-4997-bebc-05d61ab95cfa.jpg" width="150" height="150"> |

## TODO
1. Refactor the whole CLI using interfaces
2. Some filters are taking a little bit of time. Need to apply go's concurrency
3. Add more filters
4. Fix some bugs in `blur` effect
