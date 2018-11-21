glog
====

Leveled execution logs for Go.

This is an efficient pure Go implementation of leveled logs in the
manner of the open source C++ package
	https://github.com/google/glog
original from
    https://github.com/golang/glog

Original repo is not under development and pr is ginored. Original 
source is taged as v1.0. See README from original repo for more 
detail.

I will try to leave original files unchanged, add new feature 
in new files.

	Basic examples:
	
		glog.Info("Prepare to repel boarders")
		glog.Fatalf("Initialization failed: %s", err)
	
	See the documentation for the V function for an explanation
	of these examples:
	
		if glog.V(2) {
			glog.Info("Starting transaction...")
		}
	
		glog.V(2).Infoln("Processed", nItems, "elements")
