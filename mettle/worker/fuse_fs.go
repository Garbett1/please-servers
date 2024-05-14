package worker

import (
	"context"
	sdkdigest "github.com/bazelbuild/remote-apis-sdks/go/pkg/digest"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"os"
	"path"
	"path/filepath"
	"sync"
	"sync/atomic"
	"syscall"
)

type Unlinkable interface {
	Unlink()
}

// FilePopulator is responsible for populating the input files for an action.
// It can either eagerly or lazily fetch the contents of files, but the
// contract is that a file matching the digest proto should be on the filesystem
type FilePopulator interface {
	PopulateFile(digest sdkdigest.Digest)
}

type countingNode struct {
	fs.LoopbackNode

	parentFS *countingFs
}

var _ = (fs.NodeOpener)((*countingNode)(nil))

func (n *countingNode) Open(ctx context.Context, flags uint32) (fs.FileHandle, uint32, syscall.Errno) {
	fh, flags, errno := n.LoopbackNode.Open(ctx, flags)
	if errno == 0 {
		n.parentFS.incrementCount(n.Path(nil))
	}
	return fh, flags, errno
}

type countingFs struct {
	mu            sync.Mutex
	pathCount     map[string]*atomic.Int32
	digestToFiles map[sdkdigest.Digest][]fileNode
	fuseServer    *fuse.Server
}

func (f *countingFs) incrementCount(pathFromRoot string) {
	f.mu.Lock()
	a, ok := f.pathCount[pathFromRoot]
	if !ok {
		a = &atomic.Int32{}
		f.pathCount[pathFromRoot] = a
	}
	a.Add(1)
	f.mu.Unlock()
}

func (f *countingFs) ResetState(files map[sdkdigest.Digest][]fileNode) {
	f.mu.Lock()
	f.digestToFiles = files
	f.pathCount = make(map[string]*atomic.Int32)
	f.mu.Unlock()
}

// newCountingNode decorates existing
func (f *countingFs) newCountingNode(rootData *fs.LoopbackRoot, _ *fs.Inode, _ string, _ *syscall.Stat_t) fs.InodeEmbedder {
	n := &countingNode{
		LoopbackNode: fs.LoopbackNode{
			RootData: rootData,
		},
		parentFS: f,
	}

	return n
}

func (w *worker) SetupFs() {

	f := countingFs{
		pathCount: make(map[string]*atomic.Int32),
	}

	join := path.Join(os.TempDir(), w.name)
	os.MkdirAll(join, 0755)

	root := &fs.LoopbackRoot{NewNode: f.newCountingNode, Path: join}
	mount, err := fs.Mount(w.rootDir, f.newCountingNode(root, nil, "", nil), &fs.Options{
		MountOptions: fuse.MountOptions{
			Debug: false,
		},
	})
	if err != nil {
		log.Fatalf("failed to mount Fuse FS: %v", err)
	}
	f.fuseServer = mount
	w.countingFs = &f
}

func (w *worker) Cleanup() {
	log.Infof("cleaning up mounts")
	if w.countingFs == nil {
		log.Errorf("we never inited the FS?")
		return
	}
	err := w.countingFs.fuseServer.Unmount()
	if err != nil {
		log.Warningf("failed to unmount, %v", err)
	}
	log.Infof("mounts cleaned")
}

func (w *worker) reportFileUsage() {
	if !w.fuseEnabled {
		return
	}

	var wastedSpace int64
	var totalSpace int64
	for digest, files := range w.countingFs.digestToFiles {
		c := 0
		totalSpace += digest.Size
		for _, file := range files {
			rel, _ := filepath.Rel(w.rootDir, file.Name)
			a, ok := w.countingFs.pathCount[rel]
			if !ok {
				continue
			}
			c = c + int(a.Load())
		}
		if c == 0 {
			log.Errorf("digest %s never opened, files referenced are %v", digest, files)
			wastedSpace += digest.Size
		}
	}
	log.Errorf("Total wasted downloads are: %d", wastedSpace)
	log.Errorf("Total downloads are: %d", totalSpace)
	log.Errorf("Fraction of usage is: %g", float64(wastedSpace)/float64(totalSpace))
}

func (w *worker) populateAllFiles(files map[sdkdigest.Digest][]fileNode, packs map[sdkdigest.Digest][]string) {
	leavesToUnlink := make([]Unlinkable, 0, len(files)+len(packs))
	defer func() {
		for _, leaf := range leavesToUnlink {
			leaf.Unlink()
		}
	}()

}
