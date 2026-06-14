package filetree

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWalkDirAndParseTreeData(t *testing.T) {
	tests := []struct {
		name         string
		setupTestDir func() string
		depth        int
		wantErr      bool
	}{
		{
			name: "depth=0 なら直下のみ取得",
			setupTestDir: func() string {
				root := t.TempDir()
				os.Mkdir(filepath.Join(root, "testDIr"), 0755)
				os.WriteFile(filepath.Join(root, "testDIr", "test.md"), []byte("test"), 0644)
				return root
			},
			depth:   0,
			wantErr: false,
		},
		{
			name: "depth=1 なら1階層下まで取得",
			setupTestDir: func() string {
				root := t.TempDir()
				os.Mkdir(filepath.Join(root, "testDIr"), 0755)
				os.WriteFile(filepath.Join(root, "testDIr", "test.md"), []byte("test"), 0644)
				return root
			},
			depth:   1,
			wantErr: false,
		},
		{
			name: "depth=-1 ならエラー",
			setupTestDir: func() string {
				root := t.TempDir()
				os.Mkdir(filepath.Join(root, "testDIr"), 0755)
				os.WriteFile(filepath.Join(root, "testDIr", "test.md"), []byte("test"), 0644)
				return root
			},
			depth:   -1,
			wantErr: true,
		},
		{
			name: "存在しないパスならエラー",
			setupTestDir: func() string {
				return filepath.Join(os.TempDir(), "invalid-path")
			},
			depth:   1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setupTestDir()
			_, err := walkDirAndParseTreeNode(path, tt.depth)

			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
		})
	}
}
