package slideshow

import "io/ioutil"

func loadFromFile(file string) (string, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(src), nil
}
