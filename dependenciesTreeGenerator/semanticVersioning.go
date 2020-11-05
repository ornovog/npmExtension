package dependenciesTreeGenerator

import (
	"strconv"
	"strings"
)

type packageNameAndMajorNum struct {
	name string
	majorNum int
}


func UpdateDependenciesBySemanticVersion(rootPackage *PackageNode){
	packageQueue := NewQueue()
	packageQueue.Enqueue(rootPackage)

	updateDependenciesBfs(packageQueue, map[packageNameAndMajorNum]string{})
}

func updateDependenciesBfs(packageQueue queue, mostUpToDateVersions map[packageNameAndMajorNum]string){
	if packageQueue.IsEmpty(){
		return
	}

	package_ := packageQueue.Dequeue()
	version := package_.Package.Version
	name := package_.Package.Name

	major, minor, patch := majorMinorAndPatchNums(version)
	nameAndMajorNum := packageNameAndMajorNum{name, major}

	updateMostUpToDateVersions(mostUpToDateVersions, nameAndMajorNum, version, minor, patch)

	iterateDependenciesBfs(package_, packageQueue, mostUpToDateVersions)

	mostUpToDateVersion, _ := mostUpToDateVersions[nameAndMajorNum]
	package_.Package.Version = mostUpToDateVersion
}

func majorMinorAndPatchNums(version string)(major, minor, patch int)  {
	numsStr :=strings.Split(version, ".")

	majorStr := numsStr[0]
	if majorStr[0] =='^'{
		majorStr = majorStr[1:]
	}
	minorStr := numsStr[1]
	patchStr := numsStr[2]

	major, _ = strconv.Atoi(majorStr)
	minor, _ = strconv.Atoi(minorStr)
	patch, _ = strconv.Atoi(patchStr)

	return major, minor, patch
}

func updateMostUpToDateVersions(mostUpToDateVersions map[packageNameAndMajorNum]string,
								nameAndMajorNum packageNameAndMajorNum, version string, minor int, patch int){
	mostUpToDateVersion, ok := mostUpToDateVersions[nameAndMajorNum]
	if !ok {
		mostUpToDateVersions[nameAndMajorNum] = version
	} else {
		_, upToDateMinor, upToDatePatch := majorMinorAndPatchNums(mostUpToDateVersion)
		if isMoreUpToDateVersion(minor, patch, upToDateMinor, upToDatePatch) {
			mostUpToDateVersions[nameAndMajorNum] = version
		}
	}
}

func isMoreUpToDateVersion(minor1, patch1,  minor2, patch2 int)bool{
	if minor1 > minor2 || (minor1 == minor2 && patch1 >= patch2){
		return true
	}

	return false
}

func iterateDependenciesBfs(package_ *PackageNode, packageQueue queue, mostUpToDateVersions map[packageNameAndMajorNum]string) {
	for i, _ := range package_.Dependencies {
		dependency := &package_.Dependencies[i]
		packageQueue.Enqueue(dependency)
	}

	updateDependenciesBfs(packageQueue, mostUpToDateVersions)
}



