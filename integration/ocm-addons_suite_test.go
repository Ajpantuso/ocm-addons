// SPDX-FileCopyrightText: 2022 Red Hat, Inc. <sd-mt-sre@redhat.com>
//
// SPDX-License-Identifier: Apache-2.0

/* #nosec */

package integration

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestAddonsPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ocm-addons Suite")
}

var (
	_pluginPath string
	_ocmBinary  string
)

var _ = BeforeSuite(func() {
	var err error

	_ocmBinary, err = getOCMBinPath()
	Expect(err).ToNot(
		HaveOccurred(),
		"The ocm-cli is not available in the system PATH "+
			"and must be installed before running these tests.",
	)

	DeferCleanup(gexec.CleanupBuildArtifacts)

	_pluginPath, err = buildPluginBinary()
	Expect(err).ToNot(
		HaveOccurred(),
		"Unable to build plug-in binary.",
	)
})

var errSetup = errors.New("test setup failed")

func getOCMBinPath() (string, error) {
	return exec.LookPath("ocm")
}

func buildPluginBinary() (string, error) {
	root, err := projectRoot()
	if err != nil {
		return "", fmt.Errorf("determining project root: %w", err)
	}

	ldflags := "-ldflags=" + strings.Join([]string{
		"-X", "'github.com/mt-sre/ocm-addons/internal/meta.version=0.0.0'",
		"-X", "'github.com/mt-sre/ocm-addons/internal/meta.commit=abcdefg'",
		"-X", "'github.com/mt-sre/ocm-addons/internal/meta.date=0000-00-00T00:00:00'",
		"-X", "'github.com/mt-sre/ocm-addons/internal/meta.builtBy=test-suite'",
	}, " ")

	args := []string{
		ldflags,
	}

	return gexec.Build(filepath.Join(root, "cmd", "ocm-addons"), args...)
}

func projectRoot() (string, error) {
	if root, ok := os.LookupEnv("PROJECT_ROOT"); ok {
		return root, nil
	}

	var buf bytes.Buffer

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stdout = &buf
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("determining top level directory from git: %w", errSetup)
	}

	return strings.TrimSpace(buf.String()), nil
}
