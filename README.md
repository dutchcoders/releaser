# releaser
Golang library to check for new releases. 


```
u := releaser.New("honeytrap", "honeytrap")
if release, err := u.Available(Version); err != nil {
        panic(err)
} else if release != nil {
        fmt.Println(color.YellowString("New version available: %s. More information: %s.", release.Name, release.HtmlURL))
}
```
