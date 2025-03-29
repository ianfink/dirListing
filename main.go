/*
 * Copyright (C) 2025 Ian M. Fink.  All rights reserved.
 *
 * This program is free software:  you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
 * more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program.  If not, please see: https://www.gnu.org/licenses.
 *
 * Tabstop:	4
 */

package main

/*
 * Imports
 */

import (
	"fmt"
	"os"
	"path"
	"log"
)

/**********************************************************************/

func main() {
	var (
		theDirEntries		[]os.DirEntry
		theDirEntry			os.DirEntry
		targetDirectory		string
		err					error
	)

	if len(os.Args) < 2 {
		fmt.Printf("Usage:  %s <target directory>\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	targetDirectory = os.Args[1]
	theDirEntries, err = os.ReadDir(targetDirectory)

	if err != nil {
		log.Fatal(err)
	}

	for _, theDirEntry = range theDirEntries {
		fmt.Println(theDirEntry.Name())
	}

} /* main */

/**********************************************************************/

/*
 * End of file:	main.go
 */

