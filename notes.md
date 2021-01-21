(1) Más adelante estaría bueno implementar esto:
"A production-ready lexer might also attach the line number, column number and filename to
a token. Why? For example, to later output more useful error messages in the parsing stage.
Instead of "error: expected semicolon token" it can output:
"error: expected semicolon token. line 42, column 23, program.monkey"
We’re not going to bother with that. Not because it’s too complex, but because it would take
away from the essential simpleness of the tokens and the lexer, making it harder to understand."
(Ball, 2017)

(2)
También estaría bueno implementar esto:
"in a production
environment it makes sense to attach filenames and line numbers to tokens, to better track
down lexing and parsing errors. So it would be better to initialize the lexer with an io.Reader
and the filename. But since that would add more complexity we’re not here to handle, we’ll
start small and just use a string and ignore filenames and line numbers."
(ibid)

(3)
"we simplified things a lot in readNumber. We only read in
integers. What about floats? Or numbers in hex notation? Octal notation? We ignore them
and just say that Monkey doesn’t support this. Of course, the reason for this is again the
educational aim and limited scope of this book."
Averiguar cómo implementar, por lo menos, floats.
