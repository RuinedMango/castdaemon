package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const intro = `
<node>
	<interface name="org.mpris.MediaPlayer2">
		<method name="PlayPause">
		</method>
			
	</interface>` + introspect.IntrospectDataString + `</node> `

func PlayPause() () {
	fmt.Println("Pause/Play Requested")
}

func mprize() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Export("Boo!", "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.castcli")
	conn.Export(introspect.Introspectable(intro), "/org/mpris/Mediaplayer2",
		"org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName("org.mpris.MediaPlayer2.castcli",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
	fmt.Println("Listening on com.github.guelfey.Demo / /com/github/guelfey/Demo ...")
	select {}
}
