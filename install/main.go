package install

import "fmt"

func Main() {
	// ListRepos("http://godoc.org/-/subrepo")
	// ListRepos("http://godoc.org/code.google.com/p/go.tools")

	for _, repo := range SubRepoUrls("http://godoc.org/-/subrepo") {
		fmt.Println("list of", repo)
		for _, u := range DirectoryURLs(repo) {
			fmt.Printf("\t%s\n", u)
		}
		fmt.Println()
	}
}
