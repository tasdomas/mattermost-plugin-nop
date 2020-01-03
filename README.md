Mattermost server nop plugin
============================

This plugin has empty implementations of all the hooks available to
mattermost server plugins.

It is meant for testing purposes only.

Building
--------

To create a mattermost plugin archive, just run:
```
$ make package
```
This will build plugin executables for linux, macos and windows and package
them along with the manifest into the plugin.tar.gz file.

Generate empty implementation
-----------------------------

To generate a new empty implementation of the `Hooks` interface, first update
the mattermost-server dependency and then run
```
$ make generate
```
