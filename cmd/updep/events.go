package main

import (
	"errors"
	"fmt"
	"time"

	"updep/pkg/components/row"
	packagemodel "updep/pkg/models/package"

	tea "github.com/charmbracelet/bubbletea"
)

type OutdatedPackagesMsg []packagemodel.Package

func getOutdatedPackages() tea.Msg {
	result, err := packagemodel.FetchOutdatedPackages()
	if err != nil {
		panic(err)
	}

	packages := []packagemodel.Package{}
	for packageName, value := range result {
		pkg, err := packagemodel.New(
			packageName,
			value.Wanted,
			value.Latest,
			value.Current,
		)
		if err != nil {
			_ = errors.New("invalid package versions")
			// TODO: handle error
			continue
		}

		packages = append(packages, *pkg)
	}

	return OutdatedPackagesMsg(packages)
}

type UpdateResultCmd struct {
	success bool
}

func updatePackages(_ []row.Row) tea.Cmd {
	return func() tea.Msg {
		timer := time.NewTimer(time.Second * 5)
		<-timer.C
		fmt.Println("🪚 timer.C:", timer.C)
		// output, err := exec.Command("npm", "update").Output()
		// if err != nil {
		// 	fmt.Println("🪚 err:", err)
		// }
		// fmt.Println("🪚 output:", output)
		return UpdateResultCmd{
			success: true,
		}
	}
}
