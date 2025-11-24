package entities

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Version will break currently if version has prefix or suffix
type Version struct {
	major int
	minor int
	patch int
}

func NewVersion(str string) (*Version, error) {
	versions := strings.Split(str, ".")

	if len(versions) != 3 {
		return nil, errors.New("invalid version")
	}

	major, err := strconv.Atoi(versions[0])
	if err != nil {
		return nil, errors.New("invalid version")
	}
	minor, err := strconv.Atoi(versions[1])
	if err != nil {
		return nil, errors.New("invalid version")
	}
	patch, err := strconv.Atoi(versions[2])
	if err != nil {
		return nil, errors.New("invalid version")
	}

	return &Version{
		major: major,
		minor: minor,
		patch: patch,
	}, nil
}

func (v Version) String() string {
	return fmt.Sprintf(
		"%d.%d.%d",
		v.major,
		v.minor,
		v.patch,
	)
}

func (v Version) Compare(b Version) int {
	if v.major > b.major {
		return 1
	} else if v.major < b.major {
		return -1
	}

	if v.minor > b.minor {
		return 1
	} else if v.minor < b.minor {
		return -1
	}

	if v.patch > b.patch {
		return 1
	} else if v.patch < b.patch {
		return -1
	}

	return 0
}
