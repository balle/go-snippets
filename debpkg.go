package main

func count() error {
	rel, _, err := archive..CachedRelease("stable")
	if err != nil {
		return err
	}

	log.Printf("rel: %=v", rel)

	return nil
}

func main() {
	if err := count(); err != nil {
		log.Fatal(err)
	}
}
