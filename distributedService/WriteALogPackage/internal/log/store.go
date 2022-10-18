package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	// 永続化するためにエンコーディングを定義
	enc = binary.BigEndian
)

const (
	// レコードの長さを定義
	lenWidth = 8
)

// fileの保持, fileへのバイトを追加,バイトの読み出し
type store struct {
	*os.File
	mu   sync.Mutex // guards
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	// file sizeの取得
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	size := uint64(fi.Size())

	return &store{
		File: f,
		size: size,
		buf:  bufio.NewWriter(f),
	}, nil
}

// 与えられたbyteをstoreに永続化する
func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pos = s.size
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, nil
	}

	w, err := s.buf.Write((p))

	if err != nil {
		return 0, 0, err
	}

	w += lenWidth
	s.size += uint64(w)
	return uint64(w), pos, nil
}

// 指定された位置に格納されているレコードを返す。
func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	// ファイルの全体の容量をsizeに格納
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}

	// size分byteを確保したスライスを作成
	b := make([]byte, enc.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(pos+lenWidth)); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return 0, err
	}

	return s.File.ReadAt(p, off)
}

func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.buf.Flush()

	if err != nil {
		return err
	}

	return s.File.Close()
}
