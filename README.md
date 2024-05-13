# GOMMAND

TODO

## Build activate-window

To build Activate window run:

```sh
make build_activate-window
```

## Using activate-window

To use activate-window add a keyboard shortcut on a system level for example :https://help.ubuntu.com/stable/ubuntu-help/keyboard-shortcuts-set.html.en
as command use fullpath to executable with the name of app that you wanna rotate, for example:

```sh
path/to/the/repo/bin/activate-window code
```

this will rotate focus/ontop app window between all open visual studio code instances.

## activate-window fallback

Yu can use a fallback command if there is no running window find.
to do so just add another parameter.
It makes activate-window to run a app when he did not found open window.

```sh
path/to/the/repo/bin/activate-window code code
```

### TODO

- activate-window : add property to bring app to defined screen 
- gommand : add sync/async recognition on gmd level and if sync run it in gmd.
Needed to use like normal alias.
NOTE: It is nesesery  to emulate terminal on some how switch user to newly generated pid.
- gommand : podzielić logike na dwa etapy : command info który zwróci informację o finalnej wartości komendy i exec który przyjmie tylko komende niezależnie od zawartości. 