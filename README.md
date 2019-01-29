# teego

golang tee

You can choose ignore ANSI Control Sequences or not.

## example
```
$ cat main.c | grep --color=always func | teego tmp.txt
$ ls --color=always | teego tmp.txt
$ ccat -C=always main.c | teego tmp.txt

# stdout & stderr -> color
# file -> no color
```

## TODO
* flagで通常のteeのような挙動に
* solve `ErrTooLong = errors.New("bufio.Scanner: token too long")`

## FYI
[How can I remove the ANSI escape sequences from a string in python - Stack Overflow]( http://stackoverflow.com/questions/14693701/how-can-i-remove-the-ansi-escape-sequences-from-a-string-in-python )

# tee man
```
$ tee man
TEE(1)                    BSD General Commands Manual                   TEE(1)

NAME
     tee -- pipe fitting

SYNOPSIS
     tee [-ai] [file ...]

DESCRIPTION
     The tee utility copies standard input to standard output, making a copy in zero or more files.  The output is unbuffered.

     The following options are available:

     -a      Append the output to the files rather than overwriting them.

     -i      Ignore the SIGINT signal.

     The following operands are available:

     file  A pathname of an output file.

     The tee utility takes the default action for all signals, except in the event of the -i option.

     The tee utility exits 0 on success, and >0 if an error occurs.

STANDARDS
     The tee function is expected to be POSIX IEEE Std 1003.2 (``POSIX.2'') compatible.

BSD                              June 6, 1993                              BSD
```
