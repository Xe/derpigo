/*
Command db-archive downloads and serves images from Derpibooru, fetched via derpigo and stored in boltdb.

Optionally this will also have a worker that connects to irc.ponychat.net and joins #derpibooru-livefeed to automatically download and add new images.
*/
package main
