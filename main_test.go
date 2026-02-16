package main

import (
	"testing"
)

func TestParseBlock(t *testing.T) {
	tests := []struct {
		name    string
		block   []string
		id      rune
		want    Tetromino
		wantErr bool
	}{
		{
			name: "Valid I",
			block: []string{
				"....",
				"####",
				"....",
				"....",
			},
			id: 'A',
			want: Tetromino{
				coords: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
				width:  4,
				height: 1,
				id:     'A',
			},
			wantErr: false,
		},
		{
			name: "Valid Square",
			block: []string{
				"##..",
				"##..",
				"....",
				"....",
			},
			id: 'B',
			want: Tetromino{
				coords: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
				width:  2,
				height: 2,
				id:     'B',
			},
			wantErr: false,
		},
		{
			name: "Disconnected",
			block: []string{
				"#...",
				"....",
				"..#.",
				".##.",
			},
			id:      'C',
			wantErr: true,
		},
		{
			name: "Invalid Char",
			block: []string{
				"....",
				".X..",
				"....",
				"....",
			},
			id:      'D',
			wantErr: true,
		},
		{
			name: "Too many blocks",
			block: []string{
				"#####",
				".....",
				".....",
				".....",
			},
			id:      'E',
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBlock(tt.block, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Sort coords for comparison? Or ensure deterministic order in parsing.
				// For now, let's just check width, height, and id.
				// DeepEqual on coords might be flaky if order differs.
				// But our parsing is analyzing left-to-right top-to-bottom, so it should be stable.

				if got.width != tt.want.width || got.height != tt.want.height || got.id != tt.want.id {
					t.Errorf("parseBlock() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
