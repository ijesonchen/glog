package glog

// AlsoToStderr .
func AlsoToStderr(also bool) {
	logging.alsoToStderr = also
}

// SetLevel .
func SetLevel(l string) error {
	return logging.stderrThreshold.Set(l)
}

// SetLogDir .
func SetLogDir(d string) {
	logDirs = append(logDirs, d)
}
