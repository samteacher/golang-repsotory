// Copyright 2013 gopm authors.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package doc

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Unknwon/com"
	"github.com/codegangsta/cli"
)

var launchpadPattern = regexp.MustCompile(`^launchpad\.net/(?P<repo>(?P<project>[a-z0-9A-Z_.\-]+)(?P<series>/[a-z0-9A-Z_.\-]+)?|~[a-z0-9A-Z_.\-]+/(\+junk|[a-z0-9A-Z_.\-]+)/[a-z0-9A-Z_.\-]+)(?P<dir>/[a-z0-9A-Z_.\-/]+)*$`)

// getLaunchpadDoc downloads tarball from launchpad.net.
func getLaunchpadDoc(client *http.Client, match map[string]string, installRepoPath string, nod *Node, ctx *cli.Context) ([]string, error) {

	if match["project"] != "" && match["series"] != "" {
		rc, err := com.HttpGet(client, com.Expand("https://code.launchpad.net/{project}{series}/.bzr/branch-format", match), nil)
		_, isNotFound := err.(com.NotFoundError)
		switch {
		case err == nil:
			rc.Close()
			// The structure of the import path is launchpad.net/{root}/{dir}.
		case isNotFound:
			// The structure of the import path is is launchpad.net/{project}/{dir}.
			match["repo"] = match["project"]
			match["dir"] = com.Expand("{series}{dir}", match)
		default:
			return nil, err
		}
	}

	var downloadPath string
	// Check if download with specific revision.
	if len(nod.Value) == 0 {
		downloadPath = com.Expand("https://bazaar.launchpad.net/+branch/{repo}/tarball", match)
	} else {
		downloadPath = com.Expand("https://bazaar.launchpad.net/+branch/{repo}/tarball/"+nod.Value, match)
	}

	// Scrape the repo browser to find the project revision and individual Go files.
	p, err := com.HttpGetBytes(client, downloadPath, nil)
	if err != nil {
		return nil, err
	}

	installPath := installRepoPath + "/" + nod.ImportPath

	// Remove old files.
	os.RemoveAll(installPath + "/")
	os.MkdirAll(installPath+"/", os.ModePerm)

	gzr, err := gzip.NewReader(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	var autoPath string // Auto path is the root path that generated by bitbucket.org.
	// Get source file data.
	dirs := make([]string, 0, 5)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		fn := h.Name
		// Check root path.
		if len(autoPath) == 0 {
			autoPath = fn[:strings.Index(fn, match["repo"])+len(match["repo"])]
		}
		absPath := strings.Replace(fn, autoPath, installPath, 1)

		switch {
		case h.FileInfo().IsDir(): // Directory.
			// Create diretory before create file.
			os.MkdirAll(absPath+"/", os.ModePerm)
			// Check if current directory is example.
			if !(!ctx.Bool("example") && strings.Contains(absPath, "example")) {
				dirs = append(dirs, absPath)
			}
		case !strings.HasPrefix(fn, "."):
			// Get data from archive.
			fbytes := make([]byte, h.Size)
			if _, err := io.ReadFull(tr, fbytes); err != nil {
				return nil, err
			}

			if err = com.WriteFile(absPath, fbytes); err != nil {
				return nil, err
			}
		}
	}

	var imports []string

	// Check if need to check imports.
	if nod.IsGetDeps {
		for _, d := range dirs {
			importPkgs, err := CheckImports(d+"/", match["importPath"], nod)
			if err != nil {
				return nil, err
			}
			imports = append(imports, importPkgs...)
		}
	}

	return imports, err
}
