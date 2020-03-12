package synergize

func RunPipeline(targetFolder string) {
	foundFilesChannel := ScanTree(targetFolder, 10)
	hashedFilesChannel := TakeFingerprints(foundFilesChannel, 5)
	SkipOrSynergizeFiles(hashedFilesChannel, SynergizeFile)
}
