# slugme

slugme is a golang package to convert any string to it's slug according to some options.

## Usage

```go
    import "github.com/colbee1/slugme"

    title := slugme.New(slugme.DefaultOptions)
    slug := title.Slug(" L'oiseau à   deux becs ")
    // l-oiseau-a-deux-becs


    ref := slugme.New(slugme.Options{
        opts.Allowed = "-+*/",
        opts.Replace = '',
        opts.KeepCase = true,
        })
    slug = ref.Slug("MF 218 F/A_LIMF 218 FA")
    // MF218F/A_LIMF218FA
```

## Options

- Allowed (string)
  List of allowed characters in slug.
  Default is "-\_+".

- Replace (string)
  Set the character to use to replace each disalowed characters.
  Default is "-".

- KeepCase (bool)
  Do not convert slug to lower case.
  Default is false.

- KeepNonAscii (bool)
  Do not try to remove diacritics.
  Default is false.

- NoShrink (bool)
  Do not replace repetition of Options.Replace by only one instance.
  Default is false.

- NoTrim (bool)
  Do not trim (begin and end) the slug.
  Default if false.
