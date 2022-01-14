# wordle hack

This is a simple script to crack the [wordle puzzle](https://www.powerlanguage.co.uk/wordle/).

## Usage

download and run:

```go
go run main.go
```

follow the instruction and win:

```bash
Try [raise] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > 02000
Try [bacon] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > 02001
Try [daunt] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > 02011
Try [panty] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > 02212
Try [tangy] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > 22222
You win!
```

## How to get the word list

Mac OS has a built-in word list, we use that:

```bash
cat /usr/share/dict/words | grep "^.\{5,5\}$" | grep -v "[A-Z]" > words.txt
```
