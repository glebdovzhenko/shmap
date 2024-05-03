# SHMAP!
SHMAP (short for Shell Map) is a command-line utility for mapping shell 
commands over database entries. The idea was born when I had to manage 40 student
accounts on one Linux server, so I found it convenient once in a while to run 
the same command for a set of users.

Also this is my attempt at learning Go.

## The general idea:
* The app holds a database with some entries
* You want to execute some kind of a shell command for a subset of said entries
* You write the command as a Jinja template
* The app maps template substitution over the DB and runs the commands.
