# My Exercisms for GO

see: [https://exercism.org/]

## Test script
The provided bash script testloop.sh is translated into a native Go application that monitors a directory for file changes and executes go test and go build using the os/exec package. The implementation mimics the original's sleep-and-check logic, utilizing os.ReadDir to check for modification timestamps. For more robust, event-driven file monitoring, developers may use the fsnotify library, or for simpler shell-like scripting, the bitfield/script library.