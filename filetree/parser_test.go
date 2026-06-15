package filetree

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWalkDirAndParseTreeNode(t *testing.T) {
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

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		expect string
		node   treeNode
	}{
		{
			name:   "1ノードのみで表示する",
			expect: ".\n",
			node:   treeNode{name: ".", isDir: false, children: []treeNode{}},
		},
		{
			name:   "単一のchildrenを表示する",
			expect: ".\n└── child\n",
			node:   treeNode{name: ".", isDir: false, children: []treeNode{{name: "child"}}},
		},
		{
			name:   "2以上のchildrenを表示する",
			expect: ".\n├── first\n└── second\n",
			node:   treeNode{name: ".", isDir: false, children: []treeNode{{name: "first"}, {name: "second"}}},
		},
		{
			name:   "dirのネストされたchildrenを表示する",
			expect: ".\n└── first\n    └── second\n",
			node:   treeNode{name: ".", isDir: false, children: []treeNode{{name: "first", isDir: true, children: []treeNode{{name: "second"}}}}},
		},
		{
			name:   "dirのネストされたchildrenと単一ファイルを表示する",
			expect: ".\n├── first\n│   └── second\n└── third\n",
			node:   treeNode{name: ".", isDir: false, children: []treeNode{{name: "first", isDir: true, children: []treeNode{{name: "second"}}}, {name: "third"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expect != tt.node.String() {
				t.Fatalf("the result is different from the expected String: %s\nexpect:\n %s\nresult:\n %s", tt.name, tt.expect, tt.node)
			}
		})
	}
}
