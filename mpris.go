package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const introPlayer = `
<node> ` + introspect.IntrospectDataString + `
	<interface name="org.mpris.MediaPlayer2">
		<method name="Raise">
		</method>
		<method name="Quit">
		</method>
		<property name="CanQuit">
		</property>
		<property name="CanRaise">
		</property>
		<property name="HasTrackList">
		</property>
		<property name="Identity">
		</property>
		<property name="SupportedUriSchemes"
		</property>
		<property name="SupportedMimeTypes">
		</property>

	</interface>

	<interface name="org.mpris.MediaPlayer2.Player">
		<method name="Next">
		</method>
		<method name="Previous">
		</method>
		<method name="Pause">
		</method>
		<method name="PlayPause">
		</method>
		<method name="Stop">
		</method>
		<method name="Play">
		</method>
		<method name="Seek">
		</method>
		<method name="SetPosition">
		</method>
		<signal name="Seeked">
		</signal>
		<property name="PlaybackStatus">
		</property>
		<property name="LoopStatus">
		</property>
		<property name="Rate">
		</property>
		<property name="Shuffle">
		</property>
		<property name="Metadata">
		</property>
		<property name="Volume">
		</property>
		<property name="Position">
		</property>
		<property name="MinimumRate">
		</property>
		<property name="MaximumRate">
		</property>
		<property name="CanGoNext">
		</property>
		<property name="CanGoPrevious">
		</property>
		<property name="CanPlay">
		</property>
		<property name="CanSeek">
		</property>
		<property name="CanControl">
		</property>

	</interface> </node>`


func PlayPause() () {
	pausetoggle()
}

func mprize() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Export(introspect.Introspectable(introPlayer), "/org/mpris/MediaPlayer2",
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
