package synergize

func RunPipeline(targetFolder string) {
	foundFilesChannel := ScanTree(targetFolder, 10)
	hashedFilesChannel := TakeFileFingerprint(foundFilesChannel, 5)
	SkipOrSynergizeFile(hashedFilesChannel)
}
