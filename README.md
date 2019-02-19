# Git CLI for lazy devs

## Description
Here you can find some common Git commands aliases.

Of course we could achive this using regular Git aliases, but the purpose of it was to learn how to create a GoLang CLI in something useful for me, and it did the work ;)

## Usage

### Adding files and commiting

```console
$ gt c Your commit message
```
This command will add all unstaged files and commit them.\
Shortcut for:

```console
$ git add .
$ git commit -m "Your commit message"
```

### Pushing

```console
$ gt ps -b=master
```
This command will push your branch to `origin`.\
Shortcut for:

```console
$ git push origin master
```
If you ommit the branch flag the default `develop` will be assumed.\
You can still use `-b=m` to push to `master`.