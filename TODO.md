# To-do list:
## Bootstrap
- [x] Create a database
- [x] Create a TOML config
- [x] Read a table from the database
- [x] Show the read table on screen using bubbletea
- [x] Refactor the tea table generation code
- [x] Make a better default DB with multiple tables
- [x] Read multiple tables from DB
- [ ] ~~Show multiple tables via tabs in UI~~
- - [x] Show multiple tables via a list on the left
- [x] Add a text box for commands
- [ ] ~~Fix DB path according to XDG as in [taskcli](https://github.com/charmbracelet/taskcli)~~
- - [x] Put configuration directory in $XDG_CONFIG_PATH / $HOME same as [NeoVIM](https://wiki.archlinux.org/title/Neovim#:~:text=%7C%7C%20fvimAUR-,Configuration,config%2Fnvim%2Finit.)
- [x] Text edit should accept `q` button press and not quit
## Optimization
- [x] Better interface for DB representation! Less pointers.
- [x] Figure out what should be passed by value and what by pointer in my database interface
- [ ] Fix that every time GetConfig() is called, the config file is loaded
## UI
- [x] Fix all columns width to content
- [x] Add UI configuration parameters to TOML
- [ ] Find a more compact list styling
- [ ] Figure out layout sizes and widths
## Functionality
- [x] Change table view on list selection + enter
- [x] Add logging
- [ ] Add logging EVERYWHERE
- [ ] ~~Figure out how focusing works and use it instead of my own switch~~
- - [ ] Not all `bubbles` implement `Focus()` / `Blur()`. What do?
- [ ] Show in the list which table is currently selected
- [ ] On `enter` get the text from textbox
## Other
- [ ] Refactor: move all widget styling from constructors to `styles.go`
- [ ] Correct initialization with 0 tables
