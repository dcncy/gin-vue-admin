package dcncy

import (
	"github.com/bwmarrin/snowflake"
	"strings"
)

// 删除字符串中的空格信息
func TrimSpace(str string) string {
	// 替换所有的空格
	str = strings.ReplaceAll(str, " ", "")
	// 替换所有的换行
	return strings.ReplaceAll(str, "\n", "")
}

// 精简字符串中的“二手房成成交”
func SimpleStrErShouFang(str string) string {
	str = TrimSpace(str)
	return strings.ReplaceAll(str, "二手房成交", "")
}

// 分割字符串，得到贝壳主键ID
func GetBeikeId(sourceStr string) string {
	// 以 "/" 进行分割，取得最后一个部分
	parts := strings.Split(sourceStr, "/")
	lastPart := parts[len(parts)-1]
	// 去掉 ".html" 部分
	id := strings.TrimSuffix(lastPart, ".html")
	return id
}

// 获取ID
func GetSnowFlakeId() int {
	node, _ := snowflake.NewNode(1)
	return int(node.Generate())
}
