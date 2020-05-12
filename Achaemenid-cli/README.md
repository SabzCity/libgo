# ChaparKhane Command-Line Interface
Achaemenid-cli is the command-line client for the Achaemenid-generator APIs! It provides access to all API functions related to make server application runtime!

## Usage

### Make new project - Use git as version control
- Make project version control by ```git init``` or clone exiting repo by ```git clone ${repository path}```.
- Add libgo to project as submodule by ```git submodule add -b master https://github.com/SabzCity/libgo```
- Build Achaemenid-cli by ```go build ./libgo/Achaemenid-cli```
- Run ChaparKhane in a terminal by ```./Achaemenid-cli```
- Choose desire services to make needed files or other actions!

### APIs
- Complete manifest in main package of service.
- Add other data to main package if needed.
- Add as many service you need by CLI services and add business logic to them!
- From CLI update service file to autogenerate some code for you!
- As you can see in file services logic layers are independent layer and you must just think locally! But if you need network stream data use ```st *achaemenid.Stream``` in your each function parameters. Don't remove it even don't need it!

### DB

### GUI

## Some useful git command
- Clone existing project with ```git clone ${repository path} --recursive --shallow-submodules```
- Change libgo version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.