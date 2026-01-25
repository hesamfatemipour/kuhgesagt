# Kuhgesagt
cowsay's evil german twin.

```
  -----------------------------------------------------
< WORT: die Wahrnehmung, -en                          >
< DEF:  perception                                    >
<                                                     >
< DE: Menschen haben eine ganz unterschiedliche       >
< Wahrnehmung von Schmerz.                            >
< EN: People have very different perceptions of pain. >
  -----------------------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```

Like Cowsay everytime you open a terminal, print's out a random word selected from [this](https://ankiweb.net/shared/info/1431033948) dictionary.

how to build: 
```
go build -o kuhgesagt
```
put it in the /bin directory (`sudo cp kuhgesagt /usr/local/bin`) and then add it to your terminal profile file (either `.zshrc` or
.bashprofile)
