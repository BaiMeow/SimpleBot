package data

var anonymousFlag = make(map[int64]string)

func AnonymousIDGetFlag(id int64) string {
	return anonymousFlag[id]

}

func AddAnonymous(id int64, flag string) {
	anonymousFlag[id] = flag
}
