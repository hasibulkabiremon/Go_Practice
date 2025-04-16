package main

func reformat(message string, formatter func(string) string) string {
	return "TEXTIO: " + formatter(formatter(formatter(message)))
}