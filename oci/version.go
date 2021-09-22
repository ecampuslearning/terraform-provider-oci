// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"log"
)

const Version = "4.45.0"
const ReleaseDate = "2021-09-22"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
