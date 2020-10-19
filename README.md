# Japanese Markdown Parser

I was facing the problem while learning japanese, that I couldn't easily switch keyboards while writing md files.
So i made a parser that would look for matches in a given md file and would convert the text in the match from romaji to hiragana. (Katakana coming later)

----

## Installing

Make sure to have go installed correctly and your $GOBIN is set.

```bash
git clone https://github.com/syentix/japanese-md-parser.git
cd japanese-md-parser/
go install
```

Now you're all set :)

## How to

1. While writing your markdown use `/j"<romaji-word>"` to mark an unparsed word
2. As soon as you finished writing your file use `j-md-parser <files>` to create the parsed versions of your file
3. Your parsed copies of the original text will be named:  `<filename>.hiragana.md`

## Available Flags

* `--romaji` This will add the romaji next to the parsed hiragana

### Example

Our file called `Example-File.md`:

```md
# /j"go" is great
```

In our console:

```bash
j-md-parser Example-File.md --romaji
```

This creates a file called `Example-File.hiragana.md`

Our file called `parsed/Example-File.hiragana.md`:

```md
# ご(go) is great
```
