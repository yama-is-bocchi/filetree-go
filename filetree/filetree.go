package filetree

import "io"

type fileTree struct {
	targetPath string
	depth      int
}

func New(targetPath string, depth int) fileTree {
	return fileTree{
		targetPath: targetPath,
		depth:      depth,
	}
}

func (tree fileTree) WriteTo(writer io.Writer) (int64, error) {
	// ターゲットのパスを巡回しパース
	// データ構造を再帰的に書き込み
	// byte変換 & 書き込み
	return 0, nil
}
