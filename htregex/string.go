package htregex

import "strings"

func MakeEmailInvisible(email string) string {
	if email == "" {
		return ""
	}
	if !IsEmail(email) {
		return MakeUsernameInvisible(email)
	}
	split := strings.Split(email, "@")

	hiddenEmail := ""
	for i := 0; i < len(split[0]); i++ {
		if i >= 3 {
			if i <= 5 {
				hiddenEmail += "*"
			}
		} else {
			hiddenEmail += string(split[0][i])
		}
	}

	hiddenEmail += "@" + split[1]
	return hiddenEmail
}

func MakeUsernameInvisible(username string) string {
	if username == "" {
		return ""
	}

	hiddenEmail := ""
	for i := 0; i < len(username)-2; i++ {
		if i >= 2 {
			if i <= 4 {
				hiddenEmail += "*"
			}
		} else {
			hiddenEmail += string(username[i])
		}
	}
	hiddenEmail += username[len(username)-2:]
	return hiddenEmail
}
