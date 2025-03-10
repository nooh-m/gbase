package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"strconv"
	"syscall"
)

func main(){
	// Ensure the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s [owner][:group] [file]\n", os.Args[0])
		os.Exit(1)
	}
	
	// Split the owner and group, and retrieve their IDs
	owner, group := split(os.Args[1])
	uid := userID(owner)
	gid := groupID(group)

	// Change the ownership of the specified file
	chown(uid, gid, os.Args[2])

	os.Exit(0)
}

func chown(uid int, gid int, file string) {
	// Change the file ownership using syscall.Chown
	if err := syscall.Chown(file, uid, gid); err != nil {
		fmt.Fprintf(os.Stderr, "chown %s: %v\n", file, err)
		os.Exit(1)
	}
}

func split(s string) (string, string) {
	// Split owner and group from the input string
	parts := strings.SplitN(s, ":", 2)
	owner := parts[0]
	var group string
	if len(parts) == 2 {
		group = parts[1]
	}
	return owner, group
}

func userID(name string) int {
	// Get the user ID for the given username
	if name == "" {
		return -1 // Default for missing user
	}
	u, err := user.Lookup(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no such user: %s\n", name)
		os.Exit(1)
	}
	id, _ := strconv.Atoi(u.Uid)
	return id
}

func groupID(name string) int {
	// Get the group ID for the given group name
	if name == "" {
		return -1 // Default for missing group
	}
	g, err := user.LookupGroup(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no such group: %s\n", name)
		os.Exit(1)
	}
	id, _ := strconv.Atoi(g.Gid)
	return id
}
