package strings

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

func AddShortRandomSuffix(filename string) string {
	rand.Seed(time.Now().UnixNano())

	// 分离文件名和扩展名
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// 生成6位随机字符串 (数字+小写字母)
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	shortRandom := make([]byte, 6)
	for i := range shortRandom {
		shortRandom[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s_%s%s", name, string(shortRandom), ext)
}
