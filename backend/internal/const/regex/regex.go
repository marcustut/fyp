package regex

import "regexp"

// RegType indicates the type of regular expression.
type RegType string

const (
	// Email is a type of regex.
	Email RegType = "email"
	// PhoneNumber is a type of regex.
	PhoneNumber = "phone_number"
	// URL is a type of regex.
	URL = "url"
	// URLGroup is a type of capturing regex.
	URLGroup = "url_group"
)

// Regexes stores regular expression used throughout the application.
var Regexes = map[RegType]*regexp.Regexp{
	Email:       regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`),
	PhoneNumber: regexp.MustCompile(`^(\+?6?01)[02-46-9]-*[0-9]{7}$|^(\+?6?01)[1]-*[0-9]{8}$`),
	URL:         regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&\/=]*)`),
	URLGroup:    regexp.MustCompile(`(http[s]?:\/\/)?([^\/\s]+\/)(.*)`),
}

// ExtractProtocolFromURL ...
func ExtractProtocolFromURL(url string) string {
	matches := Regexes[URLGroup].FindStringSubmatch(url)
	return matches[1]
}

// ExtractHostFromURL ...
func ExtractHostFromURL(url string) string {
	matches := Regexes[URLGroup].FindStringSubmatch(url)
	return matches[2]
}
