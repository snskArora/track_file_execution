package main

import (
	// "fmt"
	// "os"
	"log"
	// "time"
	"path/filepath"
	"regexp"
)

type (
	Regexp = regexp.Regexp
	DirEntry = filepath.DirEntry
)

func list_all_files_recursive(base_dir string, match *Regexp, match_on_abs_path bool, fls *[]string) error {
	err := filepath.WalkDir(base_dir, func(path string, info DirEntry, err error) error {
		if err != nil {
			log.Println("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			err2 := list_all_files_recursive(path, match, match_on_abs_path, fls)
			if err2 != nil {
				return err2
			}
		} else {
			var specimen string
			if match_on_abs_path {
				specimen = path
			} else {
				specimen = info.Name()
			}
			if match.MatchString(specimen) {
				*fls = append(*fls, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Got an error while processing path %s", base_dir)
		log.Fatal(err)
	}
	return nil
}

func main() {
	pattern := "*"
	base_dir := "/media/sunny/Windows/Users/saror/github"
	use_match_on_abs_path := false

	all_files := []string

	if err := list_all_files_recursive(base_dir, regexp.MustCompile(pattern), use_match_on_abs_path, &all_files); err != nil {
		log.Fatal(err)
	}

	log.Println("checkpoint")
	log.Println(all_files)

	// content, err := os.ReadFile("/media/sunny/Windows/Users/saror/github/theSTASH/tfvar_renderer/render.py")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	//    sdata := string(content)

	// chck64 := show(hashString(64, sdata))
	// chck128 := show(hashString(128, sdata))
	// var i uint64
	//
	// log.Println("starting...")
	// start := time.Now()
	// for i = 0; i < 18446744073709551615; i++ {
	// 	h := show(hashString(64, sdata))
	// 	if h != chck64 {
	// 		log.Println("hash mismatch: %s", h)
	// 	}
	//    }
	//    duration := time.Since(start)
	// start = time.Now()
	//    log.Printf("myFunction 64 took %v\n", duration)
	// for i = 0; i < 18446744073709551615; i++ {
	//     h := show(hashString(128, sdata))
	//     if h != chck128 {
	// 		log.Println("hash mismatch: %s", h)
	// 	}
	// }
	//    duration = time.Since(start)
	//    log.Printf("myFunction 128 took %v\n", duration)
}
