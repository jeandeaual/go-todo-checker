package main

import "log"

// infof logs a message with the INFO level
func infof(msg string, v ...interface{}) {
	log.Printf("[INFO]  "+msg, v...)
}

// infof logs a message with the ERROR level
func errorf(msg string, v ...interface{}) {
	log.Printf("[ERROR] "+msg, v...)
}
