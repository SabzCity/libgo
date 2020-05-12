/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	generator "../Achaemenid-generator"
	"../assets"
	"../syllab"
)

const (
	// ChaparKhaneVersion must update in each release!
	ChaparKhaneVersion = "v0.5.7"
)

var (
	// ServiceRootLocation is location of repository root folder.
	ServiceRootLocation string
)

func init() {
	// Indicate ServiceRootLocation
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		buildLog("No caller information, So we can't specify service root location")
		return
	}
	ServiceRootLocation = path.Dir(path.Dir(path.Dir(filename)))
}

func main() {
	defer saveLog()

	// Add some generic data to first of file.
	buildLog("Generator version:", ChaparKhaneVersion)

	// print contact information to code-generate.log
	buildLog("You may help us and create issue:")
	buildLog("https://github.com/SabzCity/libgo/issues/new")
	buildLog("For more information, see:")
	buildLog("https://github.com/SabzCity/libgo/Achaemenid-cli")

	start := time.Now()
	buildLog("Code generate start at", start)

	buildLog("Service root location is at", ServiceRootLocation)

	// Parse repository
	var repo = assets.NewFolder("")
	var err error
	err = repo.ReadRepositoryFromFileSystem(ServiceRootLocation)
	if err != nil {
		buildLog("Read repository face this error:", err)
	}

	buildLog("Enter desire chaparkhane CLI service ID. You can select:")
	buildLog("0  : Nothing DO!!! Prevent mistakes!!!")
	buildLog("1  : Quit without save changes")
	buildLog("2  : Save changes and quit")
	buildLog("3  : Save changes without quit")
	buildLog("4  : Add project template to repository")
	buildLog("5  : Add new service to apis/services folder")
	buildLog("6  : Update Syllab encoder||decoder methods in given file name")
	buildLog("----------------------------------------")
Choose:
	var requestedService int
	fmt.Scanln(&requestedService)
	buildLog("You choose ", requestedService)

	switch requestedService {
	case 0:
		buildLog("Nothing DO in this ID to prevent often mistakes enter multi times!!!")
	case 1:
		buildLog("See you soon!")
		goto Exit
	case 2:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			buildLog("Unable to write to repo:", err)
		}
		buildLog("All changes write to disk as you desire!")
		buildLog("See you soon!")
		goto Exit
	case 3:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			buildLog("Unable to write to repo:", err)
		}
		buildLog("All changes write to disk as you desire!")
	case 4:
		_, err := MakeNewProject(&MakeNewProjectReq{Repo: repo})
		if err != nil {
			buildLog("Add project template to repository face this error:", err)
		}
		buildLog("Add project template had been succeed!!")
	case 5:
		buildLog("Enter desire service name in ```kebab-case``` like ```register-new-person```")
		var serviceName string
		fmt.Scanln(&serviceName)
		buildLog("Desire name: ", serviceName)

		var file = assets.File{
			Name: serviceName,
		}
		err = generator.MakeNewServiceFile(&file)
		if err != nil {
			buildLog("Add new service template face this error:", err)
			break
		}
		repo.Dependencies[FolderNameAPIs].Dependencies[FolderNameAPIsServices].SetFile(&file)
		buildLog("Add new service had been succeed!!")
	case 6:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		err = syllab.CompleteEncoderMethodSafe(file)
		if err != nil {
			buildLog("Update Syllab encoder||decoder face this error:", err)
		}
		buildLog("Update Syllab encoder||decoder had been succeed!!")
	case 7:
		// res, err := generator.UpdateProjectTemplate(&ReqUpdateProjectTemplate002{readRepo})
		// if err != nil {
		// 	buildLog("Update project template error:", err)
		// }
		// repo = res.Repo
	case 8:
	}

	buildLog("----------------------------------------")
	buildLog("Enter new desire chaparkhane service ID:")
	goto Choose

Exit:
	defer buildLog("CLI app run duration:", time.Since(start))
}