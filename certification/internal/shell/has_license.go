package shell

import (
	"fmt"
	"strings"

	"github.com/redhat-openshift-ecosystem/openshift-preflight/certification"
	"github.com/redhat-openshift-ecosystem/openshift-preflight/cli"
	log "github.com/sirupsen/logrus"
)

const (
	licensePath         = "/licenses"
	newLine             = "\n"
	minLicenseFileCount = 1
	//size format string for find command, c denotes bytes
	minLicenseSize = "+1c"
)

// HasLicenseCheck evaluates that the image contains a license definition available at
// /licenses.
type HasLicenseCheck struct{}

func (p *HasLicenseCheck) Validate(imgRef certification.ImageReference) (bool, error) {
	licenseFileList, err := p.getDataToValidate(imgRef.ImageURI)
	if err != nil {
		return false, err
	}
	return p.validate(licenseFileList)
}

func (p *HasLicenseCheck) getDataToValidate(image string) (string, error) {
	runOpts := cli.ImageRunOptions{
		EntryPoint:     "stat",
		EntryPointArgs: []string{licensePath},
		LogLevel:       "debug",
		Image:          image,
	}
	runReport, err := podmanEngine.Run(runOpts)
	if err != nil {
		log.Error(fmt.Sprintf("Error when checking for %s : ", licensePath), err)
		log.Errorf("Stdout: %s", runReport.Stdout)
		log.Debugf("Stderr: %s", runReport.Stderr)
		return "", err
	}
	runOpts.EntryPoint = "find"
	runOpts.EntryPointArgs = []string{licensePath, "-type", "f", "-size", minLicenseSize}
	runReport, err = podmanEngine.Run(runOpts)
	if err != nil {
		log.Error(fmt.Sprintf("Error when listing %s : ", licensePath), err)
		log.Errorf("Stdout: %s", runReport.Stdout)
		log.Debugf("Stderr: %s", runReport.Stderr)
		return "", err
	}
	return runReport.Stdout, nil
}

func (p *HasLicenseCheck) validate(licenseFileList string) (bool, error) {
	licenseFileCount := strings.Count(licenseFileList, newLine)
	log.Infof("%d Licenses found", licenseFileCount)
	return licenseFileCount >= minLicenseFileCount, nil
}

func (p *HasLicenseCheck) Name() string {
	return "HasLicense"
}

func (p *HasLicenseCheck) Metadata() certification.Metadata {
	return certification.Metadata{
		Description:      "Checking if terms and conditions applicable to the software including open source licensing information are present. The license must be at /licenses",
		Level:            "best",
		KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
		CheckURL:         "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	}
}

func (p *HasLicenseCheck) Help() certification.HelpText {
	return certification.HelpText{
		Message:    "Check HasLicense encountered an error. Please review the preflight.log file for more information.",
		Suggestion: "Create a directory named /licenses and include all relevant licensing and/or terms and conditions as text file(s) in that directory.",
	}
}
