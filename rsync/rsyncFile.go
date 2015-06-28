package rsync
import (
	"crypto/md5"
	"hash/fnv"
)

const (
	BLOCK_SIZE = 500
	MOD_VALUE = 65536 // 2^16
)

type RsyncFile struct {
	file []byte
	blocks []Block
}

type RsyncUpdatedFile struct {
	file []byte
	hashes map[uint16][]Block
}

type Block struct {
	start    int
	end      int
	content []byte
	md5     [md5.Size]byte
	checksum uint32
}

func createBlocks(file []byte) (blocks []Block) {
	chunks := split(file, BLOCK_SIZE)
	blocks = make([]Block, len(chunks))
	for i, chunk := range chunks {
		blocks[i] = Block{
			start: i * BLOCK_SIZE,
			end: (i * BLOCK_SIZE) + len(chunk),
			content: chunk,
			md5: md5.Sum(chunk),
			checksum: newRollingChecksum(chunk, BLOCK_SIZE).Sum(),
		}
	}

	return
}

func NewRsyncFile(file []byte) *RsyncFile {
	f := &RsyncFile{
		file: file,
		blocks: createBlocks(file),
	}
	return f
}

func computeHash(checksum uint32) uint16 {
	hash := fnv.New32()
	hash.Write([]byte{byte(block.checksum)})
	return uint16(hash.Sum32())
}

func NewRsyncUpdatedFile(file []byte, blocks []Block) *RsyncUpdatedFile {
	f := &RsyncUpdatedFile{
		file: file,
	}
	m := make(map[uint16][]Block)
	for _, block := range blocks {
		hash := computeHash(block.checksum)
		m[hash] = append(m[hash], block)
	}
	f.hashes = m
	return f
}

func (r *RsyncUpdatedFile) scanBlocks() {
	checksum := newRollingChecksum(r.file, BLOCK_SIZE)
	for range r.file {
		hash := computeHash(checksum)
		if len(r.hashes[hash]) != 0 {
			for block := range r.hashes[hash] {
				if block.checksum == checksum && md5
			}
		}
	}
}