# Japanese Markdown Parser

I was facing the problem while learning japanese, that I couldn't easily switch keyboards while writing md files.
So i made a parser that would look for matches in a given md file and would convert the text in that matched my regex to Hiragana. (Katakana coming later)

----

## How to

1. While writing your markdown use `/j"<romaji-word>"` to mark a to mark an unparsed word
2. As soon as you finished writing your file use `./j-md-parser <files>` to create the parsed versions of your file
3. Your parsed copies of the original text will be named:  `<filename>(parsed).md`

### Example

Our file called `Example-File.md`:

```md
# /j"go" is great
```

In our console:

```bash
./j-md-parser Example-File.md
```

This creates a file called `Example-File(parsed).md`

Our file called `Example-File(parsed).md`:

```md
# ご(go) is great
```