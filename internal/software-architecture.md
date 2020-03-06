## Data Flow

```
ScanFolderTree() -{FileInfo}-> TakeFileFingerprint() -{FileOnDiskInfo}->

-> PersistFileOnDiskInfo(add)
-> SkipOrSynergizeFile() -{FileOnDiskInfo}-> PersistFileOnDiskInfo(remove)
```

## Classes

*FileInfo*
- os.FileInfo
- Path() String
- FileSystem() String
- Inode() String

*InodeInfo*
- fileSystem String
- inodeID uint64
- Address String
- Fingerprint() String
- FromFileInfo() InodeInfo

*FileOnDiskInfo*
- FileInfo
- InodeInfo
- Save() *-- for later*
- Load() *-- for later*