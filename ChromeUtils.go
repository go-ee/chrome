package chrome

import (
	"github.com/Sirupsen/logrus"
	"os/user"
	"path/filepath"
)

func DeleteExtension(pattern string) (err error) {
	var usr *user.User
	if usr, err = user.Current(); err != nil {
		return
	}

	chromeUserDataDefault := usr.HomeDir + "/AppData/Local/Google/Chrome/User Data/Default - Copy"
	extensionPattern := chromeUserDataDefault + "/Extensions/*/*/"

	var matches []string
	searchPattern := extensionPattern + pattern
	if matches, err = filepath.Glob(searchPattern); err != nil {
		if len(matches) == 1 {

		} else {
			logrus.Infof("More that one extension foundFound: %v", matches)
		}
		return
	}

	if len(matches) != 0 {
		logrus.Infof("Found: %v", matches)
	} else {
		logrus.Infof("No matches found for: %v", searchPattern)
	}
	return
}
